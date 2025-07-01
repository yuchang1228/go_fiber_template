package library

import (
	"app/config"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var (
	bucketName      = config.Config("GCS_BUCKET_NAME")
	credentialsFile = config.Config("GCS_CREDENTIALS_FILE")
)

type ServiceAccount struct {
	ClientEmail string `json:"client_email"`
	PrivateKey  string `json:"private_key"`
}

func UploadToGCS(file io.Reader, objectName string, isPublic bool) error {
	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return fmt.Errorf("GCS client error: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	obj := client.Bucket(bucketName).Object(objectName)
	wc := obj.NewWriter(ctx)
	mimeType := mime.TypeByExtension(filepath.Ext(objectName))
	wc.ContentType = mimeType
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("write to GCS failed: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("closing GCS writer failed: %v", err)
	}

	if isPublic {
		if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
			return fmt.Errorf("failed to set public ACL: %v", err)
		}
	}

	return nil
}

func GenerateGcsURL(objectName string) string {
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, objectName)
}

func GenerateGcsSignedURL(objectName string, duration time.Duration) (string, error) {
	data, err := os.ReadFile(credentialsFile)
	if err != nil {
		return "", fmt.Errorf("failed to read credentials: %v", err)
	}

	var sa ServiceAccount
	if err := json.Unmarshal(data, &sa); err != nil {
		return "", fmt.Errorf("failed to parse credentials: %v", err)
	}

	opts := &storage.SignedURLOptions{
		GoogleAccessID: sa.ClientEmail,
		PrivateKey:     []byte(sa.PrivateKey),
		Method:         "GET",
		Expires:        time.Now().Add(duration),
	}

	url, err := storage.SignedURL(bucketName, objectName, opts)
	if err != nil {
		return "", err
	}
	return url, nil
}
