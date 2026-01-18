package minio

import (
	"context"
	mn "github.com/minio/minio-go/v7"
)

func (m *minio) DeleteImage(ctx context.Context, bucketName, imagePath string) error {
	err := m.client.RemoveObject(ctx, bucketName, imagePath, mn.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
