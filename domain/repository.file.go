package domain

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

	r, err := zip.OpenReader(fullPath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		path := filepath.Join(fullDestPath, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, 0755)
			continue
		}

		os.MkdirAll(filepath.Dir(path), 0755)

		in, err := f.Open()
		if err != nil {
			return err
		}
		defer in.Close()

		out, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, in)
		if err != nil {
			return err
		}
	}
	return nil
}
