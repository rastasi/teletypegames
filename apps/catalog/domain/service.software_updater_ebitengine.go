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
	metadata_filename := versioned_name + ".metadata.json"
	html_dirname := versioned_name

	fmt.Printf("Ebitengine Updater: Processing zip file: %s\n", zip_filename)
	fmt.Printf("Ebitengine Updater: Processing metadata file: %s\n", metadata_filename)
	fmt.Printf("Ebitengine Updater: Processing HTML directory: %s\n", html_dirname)

	if s.fileRepository.FileExists(html_dirname) {
		fmt.Printf("Ebitengine Updater: Removing existing HTML directory: %s\n", html_dirname)
		s.fileRepository.DeleteDir(html_dirname)
	}

	fmt.Printf("Ebitengine Updater: Creating HTML directory: %s\n", html_dirname)
	s.fileRepository.CreateDir(html_dirname)
	fmt.Printf("Ebitengine Updater: Unzipping file: %s to directory: %s\n", zip_filename, html_dirname)
	s.fileRepository.UnzipFile(zip_filename, html_dirname)

	html_dir_path := s.fileRepository.GetPath(html_dirname)
	metadata_path := s.fileRepository.GetPath(metadata_filename)

	var err error

	fmt.Printf("Ebitengine Updater: Extracting metadata from metadata file: %s\n", metadata_path)
	meta_data, err := s.GetMetadata(metadata_path)
	if err != nil {
		fmt.Printf("Ebitengine Updater: Error extracting metadata: %s\n", err.Error())
		return err
	}
	fmt.Printf("Ebitengine Updater: Extracted metadata: %+v\n", meta_data)
	software := s.BuildSoftware(meta_data)

	err = s.softwareRepository.UpdateOrCreate(software)
	if err != nil {
		fmt.Printf("Ebitengine Updater: Error updating or creating software: %s\n", err.Error())
		return err
	}

	release := Release{
		SoftwareID:     software.ID,
		Version:        version,
		CartridgePath:  "",
		SourcePath:     "",
		HTMLFolderPath: html_dir_path,
	}

	fmt.Printf("Ebitengine Updater: Creating release for software ID: %d, version: %s\n", software.ID, version)
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
