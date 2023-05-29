package slack

import (
	"context"
	"io"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func S3BucketObjectReader(s3ObjectUri string) ([]byte, error) {
	// TODO: Check if first character of s3ObjectUri is [a-zA-Z]

	s3ObjectUriPart := strings.Split(s3ObjectUri, "/")
	bucketName := s3ObjectUriPart[0]
	objectKey := strings.Join(s3ObjectUriPart[1:], "/")

	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Println("Couldn't load default configuration. Have you set up your AWS account?")
		log.Println(err)
	}
	s3Client := s3.NewFromConfig(sdkConfig)

	result, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Printf("Couldn't get object %v:%v. Here's why: %v\n", bucketName, objectKey, err)
		return nil, err
	}
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func LocalFileReader(filePath string) ([]byte, error) {
	body, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return body, nil
}
