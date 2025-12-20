package storage

import (
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

type UploadResult struct {
	Filename string
	Size     int64
	Path     string
}

type Storage interface {
	Save(file *multipart.FileHeader) (*UploadResult, error)
}

type S3CompatibleStorage struct {
	Client   *minio.Client
	Bucket   string
	BasePath string
}

