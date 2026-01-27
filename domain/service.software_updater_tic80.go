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

	cartridgePath := s.fileRepository.GetPath(name + "-" + version + ".tic")
	sourcePath := s.fileRepository.GetPath(name + "-" + version + ".lua")
	zipPath := s.fileRepository.GetPath(name + "-" + version + ".html.zip")
	htmlFolderPath := s.fileRepository.GetPath(name + "-" + version)

	s.fileRepository.CreateDir(htmlFolderPath)
	s.fileRepository.UnzipFile(zipPath, htmlFolderPath)

	var err error

	metaData, err := s.GetMetadata(sourcePath)
	if err != nil {
		return err
	}

	software := s.BuildSoftware(metaData)

	err = s.softwareRepository.UpdateOrCreate(software)
	if err != nil {
		return err
	}

	release := Release{
		SoftwareID:     software.ID,
		Version:        version,
		CartridgePath:  cartridgePath,
		SourcePath:     sourcePath,
		HTMLFolderPath: htmlFolderPath,
	}

	s.releaseRepository.Create(&release)

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

func (s *SoftwareUpdaterTIC80Service) GetMetadata(sourcePath string) (map[string]string, error) {

	file, err := os.Open(sourcePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	metaData := make(map[string]string)
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
			metaData[strings.ToLower(key)] = val
		} else {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return metaData, nil
}
