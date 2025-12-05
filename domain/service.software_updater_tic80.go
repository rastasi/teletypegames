package domain

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const TIC80_GAME_EXTENSION = ".tic" // Assuming .tic is the extension for TIC-80 games

type SoftwareUpdaterTIC80Service interface {
	UpdateTIC80Software(name string) error
}

type softwareUpdaterTIC80Service struct {
	softwareRepository SoftwareRepository
	releaseRepository  ReleaseRepository
}

func (s *softwareUpdaterTIC80Service) UpdateTIC80Software(name string) error {
	fmt.Printf("TIC80 Updater: Starting update for name: %s\n", name)
	gamePath, _ := os.LookupEnv("GAMES_DIR")
	// Let's assume 'name' could be "mygame" which means "mygame.html.zip" or "mygame.lua" / "mygame.tic"

	// Case 1: HTML game (name.html.zip)
	zipFileName := fmt.Sprintf("%s.html.zip", name)
	zipFilePath := filepath.Join(gamePath, zipFileName)

	if _, err := os.Stat(zipFilePath); err == nil { // If the zip file exists
		baseName := name
		destDir := filepath.Join(gamePath, "html", baseName) // Destination is GAMES_DIR/html/name
		fmt.Printf("TIC80 Updater: Unzipping %s to %s\n", zipFilePath, destDir)

		if err := s.unzipSource(zipFilePath, destDir); err != nil {
			return fmt.Errorf("error unzipping %s: %w", zipFileName, err)
		}

		if err := os.Remove(zipFilePath); err != nil {
			fmt.Printf("TIC80 Updater: Deleting zip file: %s\n", zipFilePath)
			return fmt.Errorf("error deleting zip file %s: %w", zipFileName, err)
		}
		fmt.Printf("TIC80 Updater: Successfully processed HTML game: %s\n", name)
	}

	// Case 2: TIC-80 Lua game (name.lua and name.tic)
	luaFileName := fmt.Sprintf("%s.lua", name)
	cartridgeFileName := fmt.Sprintf("%s.tic", name)
	luaFilePath := filepath.Join(gamePath, luaFileName)
	cartridgeFilePath := filepath.Join(gamePath, cartridgeFileName)

	if _, err := os.Stat(luaFilePath); err == nil { // If the lua file exists
		if _, err := os.Stat(cartridgeFilePath); os.IsNotExist(err) {
			return fmt.Errorf("missing cartridge file '%s' for '%s'", cartridgeFileName, luaFileName)
		}

		software, version := s.parseMeta(luaFilePath, name)
		if version == "" {
			return fmt.Errorf("missing version info in '%s'", luaFileName)
		}

		if err := s.softwareRepository.UpdateOrCreate(&software); err != nil {
			return err
		}

		// Create software directory
		softwareDir := filepath.Join(gamePath, software.Name)
		fmt.Printf("TIC80 Updater: Creating software directory: %s\n", softwareDir)
		os.MkdirAll(softwareDir, os.ModePerm)

		// New file paths
		newCartridgePath := filepath.Join(softwareDir, fmt.Sprintf("%s-%s.tic", software.Name, version))
		newSourcePath := filepath.Join(softwareDir, fmt.Sprintf("%s-%s.lua", software.Name, version))

		// Move files
		fmt.Printf("TIC80 Updater: Moving cartridge from %s to %s\n", cartridgeFilePath, newCartridgePath)
		if err := os.Rename(cartridgeFilePath, newCartridgePath); err != nil {
			return err
		}
		fmt.Printf("TIC80 Updater: Moving Lua source from %s to %s\n", luaFilePath, newSourcePath)
		if err := os.Rename(luaFilePath, newSourcePath); err != nil {
			return err
		}

		// Create Release
		release := &Release{
			SoftwareID:    software.ID,
			Version:       version,
			CartridgePath: newCartridgePath,
			SourcePath:    newSourcePath,
		}
		if err := s.releaseRepository.Create(release); err != nil {
			return err
		}
		fmt.Printf("TIC80 Updater: Successfully processed Lua game: %s, version: %s\n", name, version)
		return nil // Lua game processed
	}

	return fmt.Errorf("no recognizable game file found for '%s' in '%s'", name, gamePath)
}

// unzipSource extracts a zip archive to a destination directory.
func (s *softwareUpdaterTIC80Service) unzipSource(source, destination string) error {
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	destination, err = filepath.Abs(destination)
	if err != nil {
		return err
	}

	for _, f := range reader.File {
		fpath := filepath.Join(destination, f.Name)

		if !strings.HasPrefix(fpath, destination) {
			return fmt.Errorf("%s: illegal file path", fpath)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

func (s *softwareUpdaterTIC80Service) parseMeta(path string, name string) (Software, string) {
	file, _ := os.Open(path)
	defer file.Close()

	var g Software
	var version string
	g.Name = name
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "--") {
			parts := strings.SplitN(strings.TrimPrefix(line, "--"), ":", 2)
			if len(parts) != 2 {
				continue
			}
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])

			switch strings.ToLower(key) {
			case "title":
				g.Title = val
			case "author":
				g.Author = val
			case "desc":
				g.Desc = val
			case "site":
				g.Site = val
			case "license":
				g.License = val
			case "version":
				version = val
			}
		} else {
			break
		}
	}
	return g, version
}
