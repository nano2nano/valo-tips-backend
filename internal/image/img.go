package image

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"valo-tips/internal/cloud"

	"github.com/olahol/go-imageupload"
)

func SaveImage(img *imageupload.Image) (string, error) {
	if img.ContentType != "image/jpeg" {
		return "", errors.New("only 'png' image")
	}
	thumb, err := imageupload.ThumbnailPNG(img, 768, 432)
	if err != nil {
		return "", err
	}
	f_name := fmt.Sprintf("%s.jpeg", time.Now().Format("20060102150405"))
	if err := cloud.Upload(f_name, bytes.NewReader(thumb.Data)); err != nil {
		return "", err
	}
	return f_name, nil
}
