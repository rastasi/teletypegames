package domain

import (
	"encoding/json"
	"fmt"
	"os"
)

type SoftwareUpdaterEbitengineMetadata struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Desc    string `json:"desc"`
	Site    string `json:"site"`
	License string `json:"license"`
}

type SoftwareUpdaterEbitengineServiceInterface interface {
	Update(name, version string) error
}

type SoftwareUpdaterEbitengineService struct {
	softwareRepository SoftwareRepositoryInterface
	releaseRepository  ReleaseRepositoryInterface
	fileRepository     FileRepositoryInterface
}

func NewSoftwareUpdaterEbitengineService(
	software_repository SoftwareRepositoryInterface,
	release_repository ReleaseRepositoryInterface,
	file_repository FileRepositoryInterface,
) *SoftwareUpdaterEbitengineService {
	return &SoftwareUpdaterEbitengineService{
		softwareRepository: software_repository,
		releaseRepository:  release_repository,
		fileRepository:     file_repository,
	}
}

func (s *SoftwareUpdaterEbitengineService) Update(name, version string) error {
	fmt.Printf("Ebitengine Updater: Starting update for name: %s, version: %s\n", name, version)

	versioned_name := name + "-" + version

	zip_filename := versioned_name + ".html.zip"
	metadata_filename := versioned_name + ".meta.json"
	html_dirname := versioned_name

	if s.fileRepository.FileExists(html_dirname) {
		s.fileRepository.DeleteDir(html_dirname)
	}

	s.fileRepository.CreateDir(html_dirname)
	s.fileRepository.UnzipFile(zip_filename, html_dirname)

	html_dir_path := s.fileRepository.GetPath(html_dirname)
	metadata_path := s.fileRepository.GetPath(metadata_filename)

	var err error

	meta_data, err := s.GetMetadata(metadata_path)
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
		CartridgePath:  "",
		SourcePath:     "",
		HTMLFolderPath: html_dir_path,
	}

	s.releaseRepository.CreateIfNotExist(&release)

	fmt.Printf("Ebitengine Updater: Successfully processed Lua cartridge: %s\n", name)
	return nil
}

func (s *SoftwareUpdaterEbitengineService) BuildSoftware(metadata SoftwareUpdaterEbitengineMetadata) *Software {
	software := &Software{
		Name:     metadata.Name,
		Title:    metadata.Title,
		Author:   metadata.Author,
		Desc:     metadata.Desc,
		Site:     metadata.Site,
		License:  metadata.License,
		Platform: "ebitengine",
	}

	return software
}

func (s *SoftwareUpdaterEbitengineService) GetMetadata(metadata_path string) (SoftwareUpdaterEbitengineMetadata, error) {
	var metadata SoftwareUpdaterEbitengineMetadata

	data, err := os.ReadFile(metadata_path)
	if err != nil {
		return metadata, err
	}

	if err := json.Unmarshal(data, &metadata); err != nil {
		return metadata, err
	}

	return metadata, nil
}
