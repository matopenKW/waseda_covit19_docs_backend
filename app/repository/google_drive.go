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

// Client is Google Drive Client
type Client struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

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

	cliJSON := os.Getenv("GOOGLE_DRIVE_API_CLIENT")
	if cliJSON == "" {
		return nil, fmt.Errorf("Not set google drive api client")
	}
	buf, err = base64.StdEncoding.DecodeString(cliJSON)
	if err != nil {
		return nil, err
	}

	cli := &Client{}
	err = json.NewDecoder(bytes.NewReader(buf)).Decode(cli)
	if err != nil {
		return nil, err
	}

	authConfig := &oauth2.Config{
		ClientID:     cli.ClientID,
		ClientSecret: cli.ClientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://www.googleapis.com/oauth2/v4/token",
		},
	}
	tks := authConfig.TokenSource(context.Background(), tok)
	newToken, err := tks.Token()
	if err != nil {
		return nil, err
	}

	return &googleDriveClient{
		Client: authConfig.Client(context.Background(), newToken),
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
