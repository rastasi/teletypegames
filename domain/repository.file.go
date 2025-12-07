package domain

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileRepositoryInterface interface {
	FileExists(fileName, basePath string) bool
	CreateDir(dirPath string) error
	DeleteFile(fileName, basePath string) error
	MoveFile(srcFileName, destPath, basePath string) error
	UnzipHTMLContent(zipFileName, baseName, basePath string) error
	GetSoftwareDir(softwareName, basePath string) string
	GetCartridgePath(softwareName, version, basePath string) string
	GetSourcePath(softwareName, version, basePath string) string
	ReadMetaFromFile(fileName, basePath string) (map[string]string, error)
}

type FileRepository struct{}

func NewFileRepository() *FileRepository {
	return &FileRepository{}
}

func (r *FileRepository) FileExists(fileName, basePath string) bool {
	filePath := filepath.Join(basePath, fileName)
	_, err := os.Stat(filePath)
	return err == nil
}

func (r *FileRepository) CreateDir(dirPath string) error {
	return os.MkdirAll(dirPath, os.ModePerm)
}

func (r *FileRepository) DeleteFile(fileName, basePath string) error {
	filePath := filepath.Join(basePath, fileName)
	return os.Remove(filePath)
}

func (r *FileRepository) MoveFile(srcFileName, destPath, basePath string) error {
	srcPath := filepath.Join(basePath, srcFileName)
	return os.Rename(srcPath, destPath)
}

func (r *FileRepository) UnzipHTMLContent(zipFileName, baseName, basePath string) error {
	zipFilePath := filepath.Join(basePath, zipFileName)
	destDir := filepath.Join(basePath, "html", baseName)

	fmt.Printf("FileRepository: Unzipping %s to %s\n", zipFilePath, destDir)

	reader, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer reader.Close()

	destDir, err = filepath.Abs(destDir)
	if err != nil {
		return err
	}

	for _, f := range reader.File {
		fpath := filepath.Join(destDir, f.Name)

		if !strings.HasPrefix(fpath, destDir) {
			return fmt.Errorf("%s: illegal file path", fpath)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *FileRepository) GetSoftwareDir(softwareName, basePath string) string {
	return filepath.Join(basePath, softwareName)
}

func (r *FileRepository) GetCartridgePath(softwareName, version, basePath string) string {
	softwareDir := r.GetSoftwareDir(softwareName, basePath)
	return filepath.Join(softwareDir, fmt.Sprintf("%s-%s.tic", softwareName, version))
}

func (r *FileRepository) GetSourcePath(softwareName, version, basePath string) string {
	softwareDir := r.GetSoftwareDir(softwareName, basePath)
	return filepath.Join(softwareDir, fmt.Sprintf("%s-%s.lua", softwareName, version))
}

func (r *FileRepository) ReadMetaFromFile(fileName, basePath string) (map[string]string, error) {
	filePath := filepath.Join(basePath, fileName)
	file, err := os.Open(filePath)
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
