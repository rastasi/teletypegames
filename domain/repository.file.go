package domain

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type FileRepositoryInterface interface {
	GetPath(path string) string
	FileExists(path string) bool
	CreateDir(path string) error
	DeleteFile(path string) error
	MoveFile(src_path, dest_path string) error
	UnzipFile(path, dest_path string) error
}

type FileRepository struct {
	fileContainerPath string
}

func NewFileRepository() *FileRepository {
	file_container_path, _ := os.LookupEnv("FILE_CONTAINER_PATH")
	return &FileRepository{
		fileContainerPath: file_container_path,
	}
}

func (fr *FileRepository) GetPath(path string) string {
	return filepath.Join(fr.fileContainerPath, path)
}

func (fr *FileRepository) FileExists(path string) bool {
	full_path := fr.GetPath(path)
	_, err := os.Stat(full_path)
	return err == nil
}

func (fr *FileRepository) CreateDir(path string) error {
	full_path := fr.GetPath(path)
	return os.MkdirAll(full_path, 0755)
}

func (fr *FileRepository) DeleteFile(path string) error {
	full_path := fr.GetPath(path)
	return os.RemoveAll(full_path)
}

func (fr *FileRepository) MoveFile(src_path, dest_path string) error {
	full_src_path := fr.GetPath(src_path)
	full_dest_path := fr.GetPath(dest_path)

	dest_dir := filepath.Dir(full_dest_path)
	if err := os.MkdirAll(dest_dir, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	return os.Rename(full_src_path, full_dest_path)
}

func (fr *FileRepository) UnzipFile(path, destPath string) error {
	fullPath := fr.GetPath(path)
	fullDestPath := fr.GetPath(destPath)

	if err := os.MkdirAll(fullDestPath, 0755); err != nil {
		return err
	}

	cmd := exec.Command("unzip", "-q", fullPath, "-d", fullDestPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return os.Remove(fullPath)
}
