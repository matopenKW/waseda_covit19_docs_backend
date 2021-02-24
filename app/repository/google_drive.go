package repository

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
)

type googleDriveRepository struct{}

type googleDriveClient struct {
	Client *http.Client
}

type googleDriveService struct {
	Service *drive.Service
}

// NewGoogleDriveRepository is new google drive repository
func NewGoogleDriveRepository() GoogleDriveRepository {
	return &googleDriveRepository{}
}

func (r *googleDriveRepository) GetClient() (GoogleDriveClient, error) {
	sdk := os.Getenv("GOOGLE_DRIVE_API_TOKEN")
	if sdk == "" {
		return nil, fmt.Errorf("Not set google drive api token")
	}
	buf, err := base64.StdEncoding.DecodeString(sdk)
	if err != nil {
		return nil, err
	}

	tok := &oauth2.Token{}
	err = json.NewDecoder(bytes.NewReader(buf)).Decode(tok)
	if err != nil {
		return nil, err
	}

	config := &oauth2.Config{}
	return &googleDriveClient{
		Client: config.Client(context.Background(), tok),
	}, nil
}

func (c *googleDriveClient) GetService() (GoogleDriveService, error) {
	srv, err := drive.New(c.Client)
	if err != nil {
		return nil, err
	}

	return &googleDriveService{
		Service: srv,
	}, nil
}

func (s *googleDriveService) Create(in io.Reader, conf *drive.File) (*drive.File, error) {
	ret, err := s.Service.Files.Create(conf).Media(in).Do()
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *googleDriveService) Delete(driveID string) error {
	return s.Service.Files.Delete(driveID).Do()
}
