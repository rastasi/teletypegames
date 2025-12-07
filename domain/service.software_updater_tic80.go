package domain

import (
	"fmt"
	"os"
)

type SoftwareUpdaterTIC80ServiceInterface interface {
	Update(name string) error
}

type SoftwareUpdaterTIC80Service struct {
	softwareRepository SoftwareRepositoryInterface
	releaseRepository  ReleaseRepositoryInterface
	fileRepository     FileRepositoryInterface
}

func NewSoftwareUpdaterTIC80Service(
	softwareRepository SoftwareRepositoryInterface,
	releaseRepository ReleaseRepositoryInterface,
	fileRepository FileRepositoryInterface,
) *SoftwareUpdaterTIC80Service {
	return &SoftwareUpdaterTIC80Service{
		softwareRepository: softwareRepository,
		releaseRepository:  releaseRepository,
		fileRepository:     fileRepository,
	}
}

func (s *SoftwareUpdaterTIC80Service) Update(name string) error {
	fmt.Printf("TIC80 Updater: Starting update for name: %s\n", name)
	contentsPath, _ := os.LookupEnv("GAMES_DIR")

	if err := s.handleHTMLContent(name, contentsPath); err == nil {
		fmt.Printf("TIC80 Updater: Successfully processed HTML content: %s\n", name)
	}

	if err := s.handleLuaCartridge(name, contentsPath); err != nil {
		return err
	}

	fmt.Printf("TIC80 Updater: Successfully processed Lua cartridge: %s\n", name)
	return nil
}

func (s *SoftwareUpdaterTIC80Service) handleHTMLContent(name, contentsPath string) error {
	zipFileName := fmt.Sprintf("%s.html.zip", name)
	baseName := name

	if err := s.fileRepository.UnzipHTMLContent(zipFileName, baseName, contentsPath); err != nil {
		return err
	}

	return s.fileRepository.DeleteFile(zipFileName, contentsPath)
}

func (s *SoftwareUpdaterTIC80Service) handleLuaCartridge(name, contentsPath string) error {
	luaFileName := fmt.Sprintf("%s.lua", name)
	cartridgeFileName := fmt.Sprintf("%s.tic", name)

	if !s.fileRepository.FileExists(luaFileName, contentsPath) {
		return fmt.Errorf("no recognizable content file found for '%s' in '%s'", name, contentsPath)
	}

	if !s.fileRepository.FileExists(cartridgeFileName, contentsPath) {
		return fmt.Errorf("missing cartridge file '%s' for '%s'", cartridgeFileName, luaFileName)
	}

	software, version, err := s.parseMeta(luaFileName, name, contentsPath)
	if err != nil {
		return err
	}

	if version == "" {
		return fmt.Errorf("missing version info in '%s'", luaFileName)
	}

	if err := s.softwareRepository.UpdateOrCreate(&software); err != nil {
		return err
	}

	if err := s.moveCartridgeFiles(name, software, version, luaFileName, cartridgeFileName, contentsPath); err != nil {
		return err
	}

	release := &Release{
		SoftwareID:    software.ID,
		Version:       version,
		CartridgePath: s.fileRepository.GetCartridgePath(software.Name, version, contentsPath),
		SourcePath:    s.fileRepository.GetSourcePath(software.Name, version, contentsPath),
	}

	return s.releaseRepository.Create(release)
}

func (s *SoftwareUpdaterTIC80Service) moveCartridgeFiles(name string, software Software, version, luaFileName, cartridgeFileName, contentsPath string) error {
	softwareDir := s.fileRepository.GetSoftwareDir(software.Name, contentsPath)

	fmt.Printf("TIC80 Updater: Creating software directory: %s\n", softwareDir)
	if err := s.fileRepository.CreateDir(softwareDir); err != nil {
		return err
	}

	newCartridgePath := s.fileRepository.GetCartridgePath(software.Name, version, contentsPath)
	newSourcePath := s.fileRepository.GetSourcePath(software.Name, version, contentsPath)

	fmt.Printf("TIC80 Updater: Moving cartridge from %s to %s\n", cartridgeFileName, newCartridgePath)
	if err := s.fileRepository.MoveFile(cartridgeFileName, newCartridgePath, contentsPath); err != nil {
		return err
	}

	fmt.Printf("TIC80 Updater: Moving Lua source from %s to %s\n", luaFileName, newSourcePath)
	if err := s.fileRepository.MoveFile(luaFileName, newSourcePath, contentsPath); err != nil {
		return err
	}

	return nil
}

func (s *SoftwareUpdaterTIC80Service) parseMeta(luaFileName, name, contentsPath string) (Software, string, error) {
	var software Software
	var version string

	metaData, err := s.fileRepository.ReadMetaFromFile(luaFileName, contentsPath)
	if err != nil {
		return software, "", err
	}

	software.Name = name
	software.Platform = "tic80"

	for key, val := range metaData {
		switch key {
		case "title":
			software.Title = val
		case "author":
			software.Author = val
		case "desc":
			software.Desc = val
		case "site":
			software.Site = val
		case "license":
			software.License = val
		case "version":
			version = val
		}
	}

	return software, version, nil
}
