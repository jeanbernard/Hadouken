package googledrive

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type GoogleDrive struct {
	service *drive.Service
}

func NewDriveService(ctx context.Context, credPath string) (*GoogleDrive, error) {
	file, err := os.ReadFile(credPath)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := getConfig(file)
	if err != nil {
		log.Fatalf("Unable to get config: %v", err)
	}

	client, err := getClient(ctx, config)
	if err != nil {
		log.Fatalf("Unable to get client: %v", err)
	}

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client), option.WithScopes("drive.DriveScope"))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	return &GoogleDrive{service: srv}, nil
}

func (g GoogleDrive) Upload(ctx context.Context, filename string) error {
	fmt.Println("Uploading to Google...")

	// Open the video file
	video, err := os.Open(fmt.Sprintf("videos/H.265/%v", filename))
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer video.Close()

	now := time.Now().UTC()
	createdAt := now.Format("2006-01-02T15:04:05Z")

	f := &drive.File{Name: "SF6Yay", CreatedTime: createdAt}

	resp, err := g.service.Files.Create(f).Media(video).ProgressUpdater(func(now, size int64) {
		fmt.Printf("%d, %d\r", now, size)
	}).Do()

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("new file id: %s\n", resp.Id)

	return nil
}

func (g GoogleDrive) Download() error {
	fmt.Println("Downloading from Google...")

	query := fmt.Sprintf("name='%s' and trashed=false", "SF6Yay")
	files, err := g.service.Files.List().Fields("files(createdTime, name, id)").Q(query).Do()
	if err != nil {
		return err
	}

	for _, file := range files.Files {
		fmt.Println(file.Id, file.CreatedTime, file.ModifiedTime, file.Name)
	}

	return nil
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(ctx context.Context, config *oauth2.Config) (*http.Client, error) {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		if err := saveToken(tokFile, tok); err != nil {
			log.Fatalf("Unable to save token %v", err)
			return nil, err
		}
	}
	return config.Client(ctx, tok), nil
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	if err := json.NewEncoder(f).Encode(token); err != nil {
		return err
	}
	return nil
}

func getConfig(file []byte) (*oauth2.Config, error) {
	config, err := google.ConfigFromJSON(file, drive.DriveFileScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	return config, nil
}
