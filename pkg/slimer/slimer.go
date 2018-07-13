package slimer

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"

	capture "github.com/smilga/capture-go"
)

// for local testing
const (
	slimerPath  = "/home/maxtraffic/Projects/scratch/slimerjs/slimerjs-1.0.0/slimerjs"
	scriptPath  = "./slimer-script/index.js"
	firefoxPath = "/home/maxtraffic/Projects/scratch/slimerjs/firefox/firefox"
)

const timoutSec = 20

// Error definitions
var (
	ErrCmd     = errors.New("Error executing slimerjs binary")
	ErrKilled  = errors.New("Error process killed by caller, timeout")
	ErrProcess = errors.New("Error process exited with error")
)

// CaptureURL runs slimer process and takes screenshot of given URL
// returns base64 encoded image
func CaptureURL(url capture.URL) (capture.Base64Image, error) {
	return slimerShoot(string(url))
}

func slimerShoot(url string) (capture.Base64Image, error) {
	var stdout bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), timoutSec*time.Second)
	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		slimerPath,
		scriptPath,
		fmt.Sprintf("url=%s", url),
	)
	cmd.Env = append(os.Environ(),
		"SLIMERJSLAUNCHER="+firefoxPath,
	)
	cmd.Stdout = &stdout
	err := cmd.Start()
	if err != nil {
		return "", fmt.Errorf("%s:  %s", ErrCmd, err)
	}

	_ = cmd.Wait()
	status := cmd.ProcessState.Sys().(syscall.WaitStatus)

	switch status.ExitStatus() {
	case -1:
		return "", fmt.Errorf("Slimerjs: %s", ErrKilled)
	case 1:
		return "", fmt.Errorf("Slimerjs: %s: %s", ErrProcess, stdout.String())
	}

	return capture.Base64Image(stdout.String()), nil
}
