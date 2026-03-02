package domain

import (
	"encoding/json"
	"fmt"
	"os"
)

type SoftwareUpdaterLoveMetadata struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Desc    string `json:"desc"`
	Site    string `json:"site"`
	License string `json:"license"`
}

type SoftwareUpdaterLoveServiceInterface interface {
	Update(name, version string) error
}

type SoftwareUpdaterLoveService struct {
	softwareRepository SoftwareRepositoryInterface
	releaseRepository  ReleaseRepositoryInterface
	fileRepository     FileRepositoryInterface
}

func NewSoftwareUpdaterLoveService(
	software_repository SoftwareRepositoryInterface,
	release_repository ReleaseRepositoryInterface,
	file_repository FileRepositoryInterface,
) *SoftwareUpdaterLoveService {
	return &SoftwareUpdaterLoveService{
		softwareRepository: software_repository,
		releaseRepository:  release_repository,
		fileRepository:     file_repository,
	}
}

func (s *SoftwareUpdaterLoveService) Update(name, version string) error {
	fmt.Printf("Love Updater: Starting update for name: %s, version: %s\n", name, version)

	versioned_name := name + "-" + version

	zip_filename := versioned_name + ".html.zip"
	metadata_filename := versioned_name + ".metadata.json"
	html_dirname := versioned_name

	fmt.Printf("Love Updater: Processing zip file: %s\n", zip_filename)
	fmt.Printf("Love Updater: Processing metadata file: %s\n", metadata_filename)
	fmt.Printf("Love Updater: Processing HTML directory: %s\n", html_dirname)

	if s.fileRepository.FileExists(html_dirname) {
		fmt.Printf("Love Updater: Removing existing HTML directory: %s\n", html_dirname)
		s.fileRepository.DeleteDir(html_dirname)
	}

	fmt.Printf("Love Updater: Creating HTML directory: %s\n", html_dirname)
	s.fileRepository.CreateDir(html_dirname)
	fmt.Printf("Love Updater: Unzipping file: %s to directory: %s\n", zip_filename, html_dirname)
	s.fileRepository.UnzipFile(zip_filename, html_dirname)

	html_dir_path := s.fileRepository.GetPath(html_dirname)
	metadata_path := s.fileRepository.GetPath(metadata_filename)

	var err error

	fmt.Printf("Love Updater: Extracting metadata from metadata file: %s\n", metadata_path)
	meta_data, err := s.GetMetadata(metadata_path)
	if err != nil {
		fmt.Printf("Love Updater: Error extracting metadata: %s\n", err.Error())
		return err
	}
	fmt.Printf("Love Updater: Extracted metadata: %+v\n", meta_data)
	software := s.BuildSoftware(meta_data)

	err = s.softwareRepository.UpdateOrCreate(software)
	if err != nil {
		fmt.Printf("Love Updater: Error updating or creating software: %s\n", err.Error())
		return err
	}

	release := Release{
		SoftwareID:     software.ID,
		Version:        version,
		CartridgePath:  "",
		SourcePath:     "",
		HTMLFolderPath: html_dir_path,
	}

	fmt.Printf("Love Updater: Creating release for software ID: %d, version: %s\n", software.ID, version)
	s.releaseRepository.CreateIfNotExist(&release)

	fmt.Printf("Love Updater: Successfully processed Lua cartridge: %s\n", name)
	return nil
}

func (s *SoftwareUpdaterLoveService) BuildSoftware(metadata SoftwareUpdaterLoveMetadata) *Software {
	software := &Software{
		Name:     metadata.Name,
		Title:    metadata.Title,
		Author:   metadata.Author,
		Desc:     metadata.Desc,
		Site:     metadata.Site,
		License:  metadata.License,
		Platform: "love",
	}

	return software
}

func (s *SoftwareUpdaterLoveService) GetMetadata(metadata_path string) (SoftwareUpdaterLoveMetadata, error) {
	var metadata SoftwareUpdaterLoveMetadata

	data, err := os.ReadFile(metadata_path)
	if err != nil {
		return metadata, err
	}

	if err := json.Unmarshal(data, &metadata); err != nil {
		return metadata, err
	}

	return metadata, nil
}
