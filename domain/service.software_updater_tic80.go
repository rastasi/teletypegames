package domain

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	software_repository SoftwareRepositoryInterface,
	release_repository ReleaseRepositoryInterface,
	file_repository FileRepositoryInterface,
) *SoftwareUpdaterTIC80Service {
	return &SoftwareUpdaterTIC80Service{
		softwareRepository: software_repository,
		releaseRepository:  release_repository,
		fileRepository:     file_repository,
	}
}

func (s *SoftwareUpdaterTIC80Service) Update(name, version string) error {
	fmt.Printf("TIC80 Updater: Starting update for name: %s, version: %s\n", name, version)

	versioned_name := name + "-" + version

	zip_path := s.fileRepository.GetPath(versioned_name + ".html.zip")
	s.fileRepository.CreateDir(versioned_name)
	s.fileRepository.UnzipFile(zip_path, versioned_name)

	cartridge_path := s.fileRepository.GetPath(versioned_name + ".tic")
	source_path := s.fileRepository.GetPath(versioned_name + ".lua")
	html_folder_path := s.fileRepository.GetPath(versioned_name)

	var err error

	meta_data, err := s.GetMetadata(source_path)
	if err != nil {
		return err
	}

	software := s.BuildSoftware(meta_data)

	err = s.softwareRepository.UpdateOrCreate(software)
	if err != nil {
		return err
	}

	release := Release{
		SoftwareID:     software.ID,
		Version:        version,
		CartridgePath:  cartridge_path,
		SourcePath:     source_path,
		HTMLFolderPath: html_folder_path,
	}

	s.releaseRepository.CreateIfNotExist(&release)

	fmt.Printf("TIC80 Updater: Successfully processed Lua cartridge: %s\n", name)
	return nil
}

func (s *SoftwareUpdaterTIC80Service) BuildSoftware(metadata map[string]string) *Software {
	software := &Software{
		Name:     metadata["name"],
		Title:    metadata["title"],
		Author:   metadata["author"],
		Desc:     metadata["desc"],
		Site:     metadata["site"],
		License:  metadata["license"],
		Platform: "tic80",
	}

	return software
}

func (s *SoftwareUpdaterTIC80Service) GetMetadata(source_path string) (map[string]string, error) {

	file, err := os.Open(source_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	meta_data := make(map[string]string)
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
			meta_data[strings.ToLower(key)] = val
		} else {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return meta_data, nil
}
