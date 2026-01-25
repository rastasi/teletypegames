package domain

import (
	"fmt"
	"os"
)

type SoftwareUpdaterTIC80ServiceInterface interface {
	Update(name, version string) error
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

func (s *SoftwareUpdaterTIC80Service) Update(name, version string) error {
	fmt.Printf("TIC80 Updater: Starting update for name: %s, version: %s\n", name, version)
	contentsPath, _ := os.LookupEnv("GAMES_DIR")

	if err := s.handleHTMLContent(name, version, contentsPath); err == nil {
		fmt.Printf("TIC80 Updater: Successfully processed HTML content: %s\n", name)
	}

	if err := s.handleLuaCartridge(name, version, contentsPath); err != nil {
		return err
	}

	fmt.Printf("TIC80 Updater: Successfully processed Lua cartridge: %s\n", name)
	return nil
}

func (s *SoftwareUpdaterTIC80Service) handleHTMLContent(name, version, contentsPath string) error {
	zipFileName := fmt.Sprintf("%s.html.zip", name)
	// The zip file is now in the versioned folder
	zipFilePathInVersionDir := s.fileRepository.GetFileInSoftwareVersionDir(name, version, zipFileName, contentsPath)

	if err := s.fileRepository.UnzipHTMLContent(zipFilePathInVersionDir, name, version, contentsPath); err != nil {
		return err
	}

	return s.fileRepository.DeleteFile(zipFilePathInVersionDir, contentsPath)
}

func (s *SoftwareUpdaterTIC80Service) handleLuaCartridge(name, version, contentsPath string) error {
	// The lua and tic files are now in the versioned folder.
	luaFileName := s.fileRepository.GetFileInSoftwareVersionDir(name, version, fmt.Sprintf("%s.lua", name), contentsPath)
	cartridgeFileName := s.fileRepository.GetFileInSoftwareVersionDir(name, version, fmt.Sprintf("%s.tic", name), contentsPath)


	if !s.fileRepository.FileExists(luaFileName, "") { // basePath is already included in luaFileName
		return fmt.Errorf("no recognizable content file found for '%s' in '%s'", name, contentsPath)
	}

	if !s.fileRepository.FileExists(cartridgeFileName, "") { // basePath is already included in cartridgeFileName
		return fmt.Errorf("missing cartridge file '%s' for '%s'", cartridgeFileName, luaFileName)
	}

	software, parsedVersion, err := s.parseMeta(luaFileName, name, contentsPath)
	if err != nil {
		return err
	}

	// Use the version from the webhook for consistency
	if version != parsedVersion {
		fmt.Printf("TIC80 Updater: Warning - parsed version '%s' from Lua file differs from provided version '%s'\n", parsedVersion, version)
	}
	
	if version == "" {
		return fmt.Errorf("missing version info (webhook or parsed) for '%s'", luaFileName)
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
		WebPlayable:   true,
	}

	return s.releaseRepository.Create(release)
}

func (s *SoftwareUpdaterTIC80Service) moveCartridgeFiles(name string, software Software, version, luaSrcPath, cartridgeSrcPath, contentsPath string) error {
	softwareVersionDir := s.fileRepository.GetSoftwareVersionDir(software.Name, version, contentsPath)

	fmt.Printf("TIC80 Updater: Creating software version directory: %s\n", softwareVersionDir)
	if err := s.fileRepository.CreateDir(softwareVersionDir); err != nil {
		return err
	}

	newCartridgePath := s.fileRepository.GetCartridgePath(software.Name, version, contentsPath)
	newSourcePath := s.fileRepository.GetSourcePath(software.Name, version, contentsPath)

	fmt.Printf("TIC80 Updater: Moving cartridge from %s to %s\n", cartridgeSrcPath, newCartridgePath)
	if err := s.fileRepository.MoveFile(cartridgeSrcPath, newCartridgePath, ""); err != nil { // srcPath includes basePath
		return err
	}

	fmt.Printf("TIC80 Updater: Moving Lua source from %s to %s\n", luaSrcPath, newSourcePath)
	if err := s.fileRepository.MoveFile(luaSrcPath, newSourcePath, ""); err != nil { // srcPath includes basePath
		return err
	}

	return nil
}

func (s *SoftwareUpdaterTIC80Service) parseMeta(luaFileName, name, contentsPath string) (Software, string, error) {
	var software Software
	var version string

	metaData, err := s.fileRepository.ReadMetaFromFile(luaFileName, "") // luaFileName already contains basePath
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
