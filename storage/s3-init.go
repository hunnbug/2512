package storage

import (
	"main/environment"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func CreateS3Client() *s3.Client {
	cfg := aws.Config{
		Credentials: credentials.NewStaticCredentialsProvider(environment.S3.PublicKey, environment.S3.PrivateKey, ""),
		EndpointResolverWithOptions: aws.EndpointResolverWithOptionsFunc(func(service, region string, _ ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           "https://hb.vkcs.cloud",
				SigningRegion: "ru-msk",
			}, nil
		}),
		Region: "ru-msk",
	}

	return s3.NewFromConfig(cfg)
}
