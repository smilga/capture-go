package image

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	capture "github.com/smilga/capture-go"
	"github.com/smilga/capture-go/pkg/shell"
)

// Error definitions
var (
	ErrNoCompression = errors.New("capture/Compress: Error no compression settings provided")
	ErrNoMimeType    = errors.New("capture/Compress: Error no mime type provided")
	ErrWritingTem    = errors.New("capture/Compress: Error writing temp file")
)

func jpegCmd(fn string) string {
	return fmt.Sprintf("cat %s | base64 -d | /opt/mozjpeg/bin/cjpeg -quality 80 | base64", fn)
}
func pngCmd(fn string) string {
	return fmt.Sprintf("cat %s | base64 -d | pngquant --quality=60-90 256 | base64", fn)
}

const (
	cat          = "cat"
	decode       = "base64 -d"
	pngCompress  = "pngquant 256"
	jpegCompress = "/opt/mozjpeg/bin/cjpeg -quality 80"
	encode       = "base64"
)

func chain(file string, m capture.MimeType) []*shell.Command {
	var compressCmd string
	if m == capture.JPEG || m == capture.JPG {
		compressCmd = jpegCompress
	} else {
		compressCmd = pngCompress
	}

	return []*shell.Command{
		&shell.Command{time.Second * 20, []string{}, fmt.Sprintf("%s %s", cat, file)},
		&shell.Command{time.Second * 20, []string{}, decode},
		&shell.Command{time.Second * 20, []string{}, compressCmd},
		&shell.Command{time.Second * 20, []string{}, encode},
	}
}

// Compress gets image and appllies compression on if have Compression settings
func Compress(image *capture.Image) error {
	if image.Compression == nil {
		return ErrNoCompression
	}
	if _, ok := capture.Supported[image.Mime]; !ok {
		return ErrNoMimeType
	}

	temp := rndFn()
	err := ioutil.WriteFile(temp, []byte(image.Encoded), 0666)
	if err != nil {
		return ErrWritingTem
	}
	defer os.Remove(temp)

	cmdPipe := chain(temp, image.Mime)

	out, err := shell.ExecPipe(cmdPipe)
	if err != nil {
		return err
	}

	if len(out) > 0 {
		image.Encoded = capture.Base64Image(out)
	}

	return nil
}

func rndFn() string {
	rand.Seed(time.Now().UnixNano())
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return hex.EncodeToString(randBytes)
}
