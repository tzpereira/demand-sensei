package storage

import (
	"context"
	"log"
	"mime/multipart"
	"net/http"
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
		log.Println("Failed to open uploaded file:", err)
		return nil, err
	}
	defer src.Close()

	log.Println("Opened file:", file.Filename, "size:", file.Size)

	exists, err := s.Client.BucketExists(context.Background(), s.Bucket)
	if err != nil {
		log.Println("Failed to check bucket existence:", err)
		return nil, err
	}
	if !exists {
		log.Println("Bucket does not exist, creating:", s.Bucket)
		err = s.Client.MakeBucket(context.Background(), s.Bucket, minio.MakeBucketOptions{})
		if err != nil {
			log.Println("Failed to create bucket:", err)
			return nil, err
		}
	}

	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + ext

	objectPath := filename
	if s.BasePath != "" {
		objectPath = filepath.Join(s.BasePath, filename)
	}

	contentType := file.Header.Get("Content-Type")
	if contentType == "" {
		buf := make([]byte, 512)
		_, _ = src.Read(buf)
		contentType = http.DetectContentType(buf)
		_, _ = src.Seek(0, 0)
	}

	log.Println("Uploading file to MinIO:", objectPath, "contentType:", contentType)
	info, err := s.Client.PutObject(
		context.Background(),
		s.Bucket,
		objectPath,
		src,
		file.Size,
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)
	if err != nil {
		log.Println("MinIO upload failed:", err)
		return nil, err
	}

	log.Println("Upload successful:", objectPath, "size:", info.Size)
	return &UploadResult{
		Filename: filename,
		Size:     info.Size,
		Path:     objectPath,
	}, nil
}
