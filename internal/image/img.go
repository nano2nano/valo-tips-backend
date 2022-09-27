package image

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"time"

	cloud "valo-tips/internal/cloud/azure"

	"github.com/olahol/go-imageupload"
)

func SaveImage(img *imageupload.Image) (string, error) {
	host := os.Getenv("AZURE_STORAGE_URL")
	containerName := os.Getenv("AZURE_STORAGE_CONTAINER_NAME")
	if host == "" || containerName == "" {
		return "", errors.New("invalid environment variables")
	}

	if img.ContentType != "image/jpeg" {
		return "", errors.New("invalid image type")
	}
	thumb, err := imageupload.ThumbnailPNG(img, 768, 432)
	if err != nil {
		return "", err
	}
	f_name := fmt.Sprintf("%s.jpeg", time.Now().Format("20060102150405"))
	if err := cloud.Upload(f_name, bytes.NewReader(thumb.Data)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s/%s", host, containerName, f_name), nil
}
