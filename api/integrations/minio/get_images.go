package minio

import (
	"context"
	"fmt"
	mn "github.com/minio/minio-go/v7"
	"log"
	"path/filepath"
	"strings"
)

func (m *minio) GetImages(ctx context.Context, bucketName, path string) []string {
	objectCh := m.client.ListObjects(ctx, bucketName, mn.ListObjectsOptions{
		Prefix:    path,
		Recursive: true,
	})

	var images []string

	for object := range objectCh {
		if object.Err != nil {
			log.Printf("Error to get image from %s, %v\n", bucketName, object.Err)
			continue
		}

		ext := strings.ToLower(filepath.Ext(object.Key))
		if ext == ".jpg" || ext == ".jpeg" {
			url := fmt.Sprintf("http://localhost:9000/%s/%s", bucketName, object.Key)
			images = append(images, url)
		}

	}

	return images
}
