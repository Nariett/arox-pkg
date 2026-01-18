package minio

import (
	"bytes"
	"context"
	mn "github.com/minio/minio-go/v7"
	"strings"
)

func (m *minio) CreateFolder(ctx context.Context, bucketName, folderName string) error {
	if !strings.HasSuffix(folderName, "/") {
		folderName += "/"
	}

	_, err := m.client.PutObject(
		ctx,
		bucketName,
		folderName,
		bytes.NewReader([]byte{}),
		0,
		mn.PutObjectOptions{
			ContentType: "application/x-directory",
		},
	)
	if err != nil {
		return err
	}

	return nil
}
