package slimer

import (
	"errors"
	"fmt"
	"time"

	capture "github.com/smilga/capture-go"
	"github.com/smilga/capture-go/pkg/shell"
)

type Device string

// Constant definitions
const (
	Desktop Device = "desktop"
	Mobile  Device = "mobile"
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
func CaptureURL(url capture.URL, dev Device) (*capture.Image, error) {
	base64, err := slimerShoot(url, dev)
	if err != nil {
		return nil, fmt.Errorf("slimer/CaptureURL: Error creating screenshot. %s", err)
	}

	return &capture.Image{
		Encoded: base64,
		Mime:    capture.PNG,
	}, nil
}

func slimerShoot(url capture.URL, dev Device) (capture.Base64Image, error) {
	out, err := shell.Exec(&shell.Command{
		Timeout: time.Duration(time.Second * 20),
		Cmd:     fmt.Sprintf("%s %s url=%s %s=true", slimerBin, slimerScript, url, dev),
	})

	if err != nil {
		return capture.Base64Image(""), fmt.Errorf("%s, output: %s", err, out)
	}

	return capture.Base64Image(out), nil
}
