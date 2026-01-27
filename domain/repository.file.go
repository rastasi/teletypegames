package domain

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileRepositoryInterface interface {
	GetPath(path string) string
	FileExists(path string) bool
	CreateDir(path string) error
	DeleteFile(path string) error
	MoveFile(srcPath, destPath string) error
	UnzipFile(path, destPath string) error
}

type FileRepository struct {
	fileContainerPath string
}

func NewFileRepository() *FileRepository {
	fileContainerPath, _ := os.LookupEnv("FILE_CONTAINER_PATH")
	return &FileRepository{
		fileContainerPath: fileContainerPath,
	}
}

func (fr *FileRepository) GetPath(path string) string {
	return filepath.Join(fr.fileContainerPath, path)
}

func (fr *FileRepository) FileExists(path string) bool {
	fullPath := fr.GetPath(path)
	_, err := os.Stat(fullPath)
	return err == nil
}

func (fr *FileRepository) CreateDir(path string) error {
	fullPath := fr.GetPath(path)
	return os.MkdirAll(fullPath, 0755)
}

func (fr *FileRepository) DeleteFile(path string) error {
	fullPath := fr.GetPath(path)
	return os.RemoveAll(fullPath)
}

func (fr *FileRepository) MoveFile(srcPath, destPath string) error {
	fullSrcPath := fr.GetPath(srcPath)
	fullDestPath := fr.GetPath(destPath)

	destDir := filepath.Dir(fullDestPath)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	return os.Rename(fullSrcPath, fullDestPath)
}

func (fr *FileRepository) UnzipFile(path, destPath string) error {
	fullPath := fr.GetPath(path)
	fullDestPath := fr.GetPath(destPath)

	reader, err := zip.OpenReader(fullPath)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %w", err)
	}
	defer reader.Close()

	if err := os.MkdirAll(fullDestPath, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	for _, file := range reader.File {
		if err := extractZipFile(file, fullDestPath); err != nil {
			return fmt.Errorf("failed to extract %s: %w", file.Name, err)
		}
	}

	return nil
}

func extractZipFile(file *zip.File, destPath string) error {
	filePath := filepath.Join(destPath, file.Name)

	if !strings.HasPrefix(filePath, filepath.Clean(destPath)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", file.Name)
	}

	if file.FileInfo().IsDir() {
		return os.MkdirAll(filePath, file.Mode())
	}

	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return err
	}

	srcFile, err := file.Open()
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}
