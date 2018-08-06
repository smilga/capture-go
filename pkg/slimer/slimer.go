package slimer

import (
	"errors"
	"fmt"
	"time"

	capture "github.com/smilga/capture-go"
	"github.com/smilga/capture-go/pkg/logger"
	"github.com/smilga/capture-go/pkg/shell"
)

var slimerBin = "slimerjs"
var slimerScript = "slimer-script/index.js"

// Error definitions
var (
	ErrCmd     = errors.New("Error executing slimerjs binary")
	ErrKilled  = errors.New("Error process killed by caller, timeout")
	ErrProcess = errors.New("Error process exited with error")
)

// CaptureURL runs slimer process and takes screenshot of given URL
func CaptureURL(url capture.URL) (*capture.Image, error) {
	base64, err := slimerShoot(url)
	if err != nil {
		return nil, fmt.Errorf("slimer/CaptureURL: Error creating screenshot. %s", err)
	}

	return &capture.Image{
		Encoded: base64,
		Mime:    capture.PNG,
	}, nil
}

func slimerShoot(url capture.URL) (capture.Base64Image, error) {
	logger.Info(fmt.Sprintf("Executing command %s %s url=%s", slimerBin, slimerScript, url))

	out, err := shell.Exec(&shell.Command{
		Timeout: time.Duration(time.Second * 40),
		Cmd:     fmt.Sprintf("%s %s url=%s", slimerBin, slimerScript, url),
	})

	if err != nil {
		return capture.Base64Image(""), fmt.Errorf("%s, output: %s", err, out)
	}

	return capture.Base64Image(out), nil
}
