package minio

import (
	"context"
	"fmt"
	"github.com/Nariett/arox-pkg/config"
	mn "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio interface {
	CreateBucket(ctx context.Context, bucketName string) error
	CreateFolder(ctx context.Context, bucketName, folderName string) error
	PutImage(ctx context.Context, bucketName, path, imageName, photoBase64 string) (string, error)
	GetImages(ctx context.Context, bucketName, path string) []string
	DeleteImage(ctx context.Context, bucketName, imagePath string) error
}

type minio struct {
	cfg    *config.Config
	client *mn.Client
}

func NewMinio(cfg *config.Config) (Minio, error) {
	client, err := mn.New(fmt.Sprintf("%s:%s", cfg.MinioHost, cfg.MinioPort), &mn.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioUser, cfg.MinioPass, ""),
		Secure: false,
	})

	if err != nil {
		return nil, err
	}

	return &minio{
		cfg:    cfg,
		client: client,
	}, nil
}
