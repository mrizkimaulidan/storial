package file

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/mrizkimaulidan/storial/pkg/time"
)

type fileService struct {
	Path     string
	Filename string
}

func (fs *fileService) CreateDir() error {
	err := os.MkdirAll(fs.Path, os.ModePerm)

	return err
}

func (fs *fileService) SetFilename(filename string) {
	fs.Filename = filename
}

func (fs *fileService) Get(fullPath string) ([]byte, error) {
	fileBytes, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

func (fs *fileService) GetFullPath() string {
	return fs.Path + "/" + fs.Filename
}

func (fs *fileService) Upload(f multipart.File, fileheader *multipart.FileHeader) (string, error) {
	filename := fmt.Sprintf("%d%s", time.CurrentTimeToUnixTimestamp(), filepath.Ext(fileheader.Filename))

	fs.SetFilename(filename)

	dst := fs.GetFullPath()

	err := fs.CreateDir()
	if err != nil {
		return "", err
	}

	file, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, f)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (fs *fileService) RemoveFile(filename string) {
	os.Remove(fs.GetPath() + "/" + filename)
}

func (fs *fileService) GetPath() string {
	return fs.Path
}

func NewService(path string) FileService {
	return &fileService{
		Path: path,
	}
}
