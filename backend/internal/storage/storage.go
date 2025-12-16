package storage

import "mime/multipart"

type UploadResult struct {
	Filename string
	Size     int64
	Path     string
}

type Storage interface {
	Save(file *multipart.FileHeader) (*UploadResult, error)
}
