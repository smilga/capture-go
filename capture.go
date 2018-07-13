package capture

import (
	"errors"
	"fmt"
	"net/url"
)

// Error definitions
var (
	ErrURLInvalid = errors.New("Error invalid url")
)

// Screenshot should have url and then it goes through app
// being screenshoted and compressed if compression present
//type Screenshot struct {
//Base64image string `json:"base64image"`
//URL         string `json:"url"`
//Compression *Compression
//}

// URL represents valid url
type URL string

// Base64Image contains base 64 encoded image string
type Base64Image string

// NewURL checks if string is valid url and creates URL
func NewURL(s string) (URL, error) {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		return "", fmt.Errorf("%s: %s", ErrURLInvalid, err)
	}
	return URL(s), nil
}

//func (s *Screenshot) ReadImage() string {
//return s.Base64image
//}

//func (s *Screenshot) WriteImage(image string) {
//s.Base64image = image
//}

// Image should have encoded image and it gets compressed
// depending on compression settings
//type Image struct {
//Base64image string `json:"base64image"`
//Compression *Compression
//}

//func (i *Image) ReadImage() string {
//return i.Base64image
//}

//func (i *Image) WriteImage(image string) {
//i.Base64image = image
//}

// Imager defines methods needed to work with images
//type Imager interface {
//ReadImage() string
//WriteImage(string)
//}

// Compression keeps compression settings
// if object has compression object image gets compressed
// based on settngs
type Compression struct {
}
