package browser

import (
	"fmt"
	"strings"
	"time"

	"github.com/smilga/capture-go"
	"github.com/smilga/capture-go/pkg/shell"
)

type Device string

func (d Device) AsArg() string {
	return strings.Replace(string(d), " ", "_", -1)
}

type Dimensions struct {
	Width  string
	Height string
}

var script = "pkg/browser/pupetteer/index.js"

func Screenshot(url capture.URL, d *Dimensions, device Device) (*capture.Image, error) {
	out, err := shell.Exec(&shell.Command{
		Timeout: time.Duration(time.Second * 20),
		Cmd:     fmt.Sprintf("node %s %s %s %s %s", script, url, d.Width, d.Height, device.AsArg()),
	})
	if err != nil {
		return nil, fmt.Errorf("%s, output: %s", err, out)
	}

	return &capture.Image{
		Encoded: capture.Base64Image(out),
		Mime:    capture.PNG,
	}, nil
}
