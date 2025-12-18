package storage

import (
	"context"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3CompatibleStorage struct {
	Client   *minio.Client
	Bucket   string
	BasePath string
}

func NewS3CompatibleStorage(
	endpoint string,
	accessKey string,
	secretKey string,
	bucket string,
	basePath string,
	useSSL bool,
) (*S3CompatibleStorage, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return &S3CompatibleStorage{
		Client:   client,
		Bucket:   bucket,
		BasePath: basePath,
	}, nil
}

func (s *S3CompatibleStorage) Save(file *multipart.FileHeader) (*UploadResult, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	filename := uuid.New().String() + "_" + file.Filename

	objectPath := filename
	if s.BasePath != "" {
		objectPath = filepath.Join(s.BasePath, filename)
	}

	info, err := s.Client.PutObject(
		context.Background(),
		s.Bucket,
		objectPath,
		src,
		file.Size,
		minio.PutObjectOptions{
			ContentType: file.Header.Get("Content-Type"),
		},
	)
	if err != nil {
		return nil, err
	}

	return &UploadResult{
		Filename: filename,
		Size:     info.Size,
		Path:     objectPath,
	}, nil
}
