package capture

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"

	"github.com/smilga/capture-go/pkg/logger"
)

// Error definitions
var (
	ErrURLInvalid      = errors.New("Error invalid url")
	ErrCompressInvalid = errors.New("Error invalid compress parameter")
)

// Constant definitions
const (
	JPEG MimeType = "image/jpeg"
	JPG  MimeType = "image/jpg"
	PNG  MimeType = "image/png"
)

var supported = map[MimeType]string{
	JPEG: ".jpeg",
	JPG:  ".jpg",
	PNG:  ".png",
}

// URL represents valid url
type URL string

// Image contains base64 encoded image and string compression
type Image struct {
	Encoded     Base64Image  `json:"image"`
	Mime        MimeType     `json:"mimeType"`
	Compression *Compression `json:"-"`
}

// Base64Image stores encoded image string
type Base64Image string

// Decode returns decoded base64 string
func (i Base64Image) Decode() (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(string(i))
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// MimeType stores image mime type
type MimeType string

// Ext retursn extension from mime type
func (t MimeType) Ext() string {
	ext, ok := supported[t]
	if !ok {
		logger.Error("Error returning extension from mime type")
	}
	return ext
}

func (t MimeType) String() string {
	return string(t)
}

// NewURL checks if string is valid url and creates URL
func NewURL(s string) (URL, error) {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		return "", fmt.Errorf("%s: %s", ErrURLInvalid, err)
	}
	return URL(s), nil
}

// Compression stores compression settings
type Compression struct{}
