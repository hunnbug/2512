package tools

import (
	"bytes"
	"context"
	"main/logging"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func S3Load(s3Client *s3.Client, bucket, key string, data []byte) error {

	_, err := s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
		ACL:    types.ObjectCannedACLPrivate,
	})
	if err != nil {
		logging.WriteLog(logging.ERROR, err)
		return err
	}
	logging.WriteLog(logging.DEBUG, "загрузка в s3 прошла успешно")
	return nil
}
