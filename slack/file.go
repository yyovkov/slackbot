package slack

import (
	"bytes"
	"errors"
	"io"
	"path/filepath"
	"strings"

	"github.com/slack-go/slack"
)

type File struct {
	Channels       []string
	InitialComment string
	File           string
	Filename       string
	Filetype       string
	Reader         io.Reader
	Title          string
}

func (f *File) fileLoader() error {

	var err error
	var fileData []byte

	switch {
	case strings.HasPrefix(f.File, "file://"):
		f.File = strings.TrimPrefix(f.File, "file://")
		fileData, err = LocalFileReader(f.File)
	case strings.HasPrefix(f.File, "s3://"):
		f.File = strings.TrimPrefix(f.File, "s3://")
		fileData, err = S3BucketObjectReader(f.File)
	default:
		return errors.New("unknown source file protocol")
	}

	if err != nil {
		return err
	}

	// Set file Reader
	f.Reader = bytes.NewReader(fileData)

	// Set file title
	if f.Filename == "" {
		f.Filename = filepath.Base(f.Filename)
	}

	// Set file title
	if f.Title == "" {
		f.Title = f.Filename
	}

	return nil
}

func (f *File) Upload(token string) (string, error) {

	// Load data to be uploaded
	err := f.fileLoader()
	if err != nil {
		return "", err
	}

	// TODO: Check if file exists
	client := slack.New(token, slack.OptionDebug(false))

	fileUploadParams := slack.FileUploadParameters{
		Channels:       f.Channels,
		InitialComment: f.InitialComment,
		Filename:       f.Filename,
		Filetype:       f.Filetype,
		Reader:         f.Reader,
		Title:          f.Title,
	}

	_, err = client.UploadFile(fileUploadParams)
	if err != nil {
		return "error", err
	}

	return "success", nil
}
