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
	UnzipHTMLContent(zipFilePath, softwareName, version, basePath string) error
	GetSoftwareDir(softwareName, basePath string) string
	GetSoftwareVersionDir(softwareName, version, basePath string) string
	GetFileInSoftwareVersionDir(softwareName, version, fileName, basePath string) string
	GetHTMLContentDir(softwareName, version, basePath string) string
	GetCartridgePath(softwareName, version, basePath string) string
	GetSourcePath(softwareName, version, basePath string) string
	ReadMetaFromFile(filePath string, basePath string) (map[string]string, error)
}

type FileRepository struct{}

func NewFileRepository() *FileRepository {
	return &FileRepository{}
}

func (r *FileRepository) FileExists(fileName, basePath string) bool {
	var filePath string
	if filepath.IsAbs(fileName) {
		filePath = fileName
	} else if basePath != "" {
		filePath = filepath.Join(basePath, fileName)
	} else {
		filePath = fileName // Assume it's relative to current working dir or absolute
	}

	_, err := os.Stat(filePath)
	return err == nil
}

func (r *FileRepository) CreateDir(dirPath string) error {
	return os.MkdirAll(dirPath, os.ModePerm)
}

func (r *FileRepository) DeleteFile(fileName, basePath string) error {
	var filePath string
	if filepath.IsAbs(fileName) {
		filePath = fileName
	} else if basePath != "" {
		filePath = filepath.Join(basePath, fileName)
	} else {
		filePath = fileName
	}
	return os.Remove(filePath)
}

func (r *FileRepository) MoveFile(srcFileName, destPath, basePath string) error {
	var srcPath string
	if filepath.IsAbs(srcFileName) {
		srcPath = srcFileName
	} else if basePath != "" {
		srcPath = filepath.Join(basePath, srcFileName)
	} else {
		srcPath = srcFileName
	}
	return os.Rename(srcPath, destPath)
}

func (r *FileRepository) UnzipHTMLContent(zipFilePath, softwareName, version, basePath string) error {
	destDir := r.GetHTMLContentDir(softwareName, version, basePath)

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
			outFile.Close()
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

func (r *FileRepository) GetSoftwareVersionDir(softwareName, version, basePath string) string {
	return filepath.Join(r.GetSoftwareDir(softwareName, basePath), version)
}

func (r *FileRepository) GetFileInSoftwareVersionDir(softwareName, version, fileName, basePath string) string {
	return filepath.Join(r.GetSoftwareVersionDir(softwareName, version, basePath), fileName)
}

func (r *FileRepository) GetHTMLContentDir(softwareName, version, basePath string) string {
	return filepath.Join(r.GetSoftwareVersionDir(softwareName, version, basePath), "html")
}

func (r *FileRepository) GetCartridgePath(softwareName, version, basePath string) string {
	return filepath.Join(r.GetSoftwareVersionDir(softwareName, version, basePath), fmt.Sprintf("%s.tic", softwareName))
}

func (r *FileRepository) GetSourcePath(softwareName, version, basePath string) string {
	return filepath.Join(r.GetSoftwareVersionDir(softwareName, version, basePath), fmt.Sprintf("%s.lua", softwareName))
}

func (r *FileRepository) ReadMetaFromFile(filePath string, basePath string) (map[string]string, error) {
	var fullPath string
	if filepath.IsAbs(filePath) {
		fullPath = filePath
	} else if basePath != "" {
		fullPath = filepath.Join(basePath, filePath)
	} else {
		fullPath = filePath
	}

	file, err := os.Open(fullPath)
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
