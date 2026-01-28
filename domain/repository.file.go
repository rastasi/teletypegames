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

func (fr *FileRepository) UnzipFile(path, dest_path string) error {
	full_path := fr.GetPath(path)
	full_dest_path := fr.GetPath(dest_path)

	r, err := zip.OpenReader(full_path)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		// Prevent path traversal vulnerability
		fpath := filepath.Join(full_dest_path, f.Name)
		if !strings.HasPrefix(fpath, full_dest_path) {
			return fmt.Errorf("%s: illegal file path", fpath)
		}

		fmt.Printf("ZIP: Extracting %s\n", fpath)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, 0755)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close() // Close outFile if f.Open fails
			return err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file readers/writers immediately after use
		rc.Close()
		outFile.Close()

		if err != nil {
			return err
		}
	}
	return nil
}
