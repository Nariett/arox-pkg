package minio

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	mn "github.com/minio/minio-go/v7"
)

func (m *minio) PutImage(ctx context.Context, bucketName, path, imageName, photoBase64 string) (string, error) {
	imageData, err := base64.StdEncoding.DecodeString(photoBase64)
	if err != nil {
		return "", err
	}

	_, err = m.client.PutObject(
		ctx,
		bucketName,
		fmt.Sprintf("%s/%s", path, imageName),
		bytes.NewReader(imageData),
		int64(len(imageData)),
		mn.PutObjectOptions{
			ContentType: "image/jpeg",
		},
	)
	if err != nil {
		return "", err
	}

	publicURL := fmt.Sprintf("http://%s:%s/%s/%s", m.cfg.MinioHost, m.cfg.MinioPort, bucketName, fmt.Sprintf("%s/%s", path, imageName))

	return publicURL, nil
}
