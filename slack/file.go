package slack

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/slack-go/slack"
)

type File struct {
	Title          string
	InitialComment string
	Filetype       string
	Filename       string
	Channels       []string
	File           string
}

func (f *File) Upload(token, channelID string) (string, error) {
	// TODO: Check if file exists
	client := slack.New(token, slack.OptionDebug(false))

	fileUploadParams := slack.FileUploadParameters{
		Channels:       []string{channelID},
		File:           f.File,
		Filename:       f.Filename,
		Filetype:       f.Filetype,
		InitialComment: f.InitialComment,
		Title:          f.Title,
	}

	_, err := client.UploadFile(fileUploadParams)
	if err != nil {
		return "error", err
	}

	return "success", nil
}

func (f *File) UploadFromS3(token, channelID string) (string, error) {

	// Deal with S3 staff
	bucketName := "delme.yovko-inm.sand.mytaverse.io"
	objectKey := "report-users-20230526.xlsx"

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
		return "", err
	}
	defer result.Body.Close()

	fileUploadParams := slack.FileUploadParameters{
		Channels:       []string{channelID},
		Reader:         result.Body,
		Filename:       f.Filename,
		InitialComment: f.InitialComment,
		Title:          f.Filename,
	}

	// Uploading
	client := slack.New(token, slack.OptionDebug(false))
	_, err = client.UploadFile(fileUploadParams)
	if err != nil {
		return "error", err
	}

	return "good", nil
}
