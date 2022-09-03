package file

import "mime/multipart"

type FileService interface {
	Get(fullPath string) ([]byte, error)
	GetPath() string
	CreateDir() error
	RemoveFile(filename string)
	Upload(f multipart.File, fileheader *multipart.FileHeader) (string, error)
	SetFilename(filename string)
	GetFullPath() string
}
