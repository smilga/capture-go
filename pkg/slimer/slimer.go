package slimer

import (
	"errors"
	"fmt"
	"time"

	capture "github.com/smilga/capture-go"
	"github.com/smilga/capture-go/pkg/shell"
)

var cmd = "/home/maxtraffic/Projects/scratch/slimerjs/slimerjs-1.0.0/slimerjs ./slimer-script/index.js"

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
	out, err := shell.Exec(&shell.Command{
		Timeout: time.Duration(time.Second * 40),
		Env: []string{
			"SLIMERJSLAUNCHER=/home/maxtraffic/Projects/scratch/slimerjs/firefox/firefox",
		},
		Cmd: fmt.Sprintf("%s url=%s", cmd, url),
	})

	if err != nil {
		return capture.Base64Image(""), fmt.Errorf("%s, output: %s", err, out)
	}

	return capture.Base64Image(out), nil
}
