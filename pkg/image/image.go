package image

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	capture "github.com/smilga/capture-go"
	"github.com/smilga/capture-go/pkg/logger"
	"github.com/smilga/capture-go/pkg/shell"
)

// Error definitions
var (
	ErrNoCompression = errors.New("capture/Compress: Error no compression settings provided")
	ErrNoMimeType    = errors.New("capture/Compress: Error no mime type provided")
	ErrWritingTem    = errors.New("capture/Compress: Error writing temp file")
)

var temp = "temp.txt"

func jpegCmd(fn string) string {
	return fmt.Sprintf("cat %s | base64 -d | cjpeg -quality 80 | base64", fn)
}
func pngCmd(fn string) string {
	return fmt.Sprintf("cat %s | base64 -d | pngquant 256 | base64", fn)
}

// Compress gets image and appllies compression on if have Compression settings
func Compress(image *capture.Image) error {
	logger.Info(fmt.Sprintf("Compressing image! %s", image.Mime))
	if image.Compression == nil {
		return ErrNoCompression
	}
	if image.Mime == "" {
		return ErrNoMimeType
	}

	err := ioutil.WriteFile(temp, []byte(image.Encoded), 0666)
	if err != nil {
		return ErrWritingTem
	}
	defer os.Remove(temp)

	var cmdString string
	switch image.Mime {
	case capture.JPEG, capture.JPG:
		cmdString = jpegCmd(temp)
	case capture.PNG:
		cmdString = pngCmd(temp)
	}

	out, err := shell.Exec(&shell.Command{
		Timeout: time.Duration(40 * time.Second),
		Cmd:     cmdString,
	})
	if err != nil {
		fmt.Println("error", err)
	}

	if len(out) > 0 {
		image.Encoded = capture.Base64Image(out)
	}

	return nil
}
