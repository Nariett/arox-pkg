package minio

import (
	"context"
	"fmt"
	mn "github.com/minio/minio-go/v7"
	"log"
)

func (m *minio) CreateBucket(ctx context.Context, bucketName string) error {
	err := m.client.MakeBucket(ctx, bucketName, mn.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := m.client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("bucket %s is alredy exists\n", bucketName)
		} else {
			return err
		}
	} else {
		log.Printf("bucket %s created\n", bucketName)
	}

	policy := fmt.Sprintf(`{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Principal": {"AWS": ["*"]},
                "Action": ["s3:GetObject"],
                "Resource": ["arn:aws:s3:::%s/*"]
            }
        ]
    }`, bucketName)

	err = m.client.SetBucketPolicy(ctx, bucketName, policy)
	if err != nil {
		log.Printf(err.Error())
	}

	return nil
}
