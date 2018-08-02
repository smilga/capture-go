package shell

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

// Constant definitions
var (
	ErrProcessKilled = errors.New("Shell: Error process killed")
	ErrProcessExited = errors.New("Shell: Error process exited")
	ErrExecutingCmd  = errors.New("Shell: Error executing command")
)

// Command stores command to be executed
type Command struct {
	Timeout time.Duration
	Env     []string
	Cmd     string
}

// Exec executes command and returns stdOut or stdErr and error
// Checks for stdout, stderr because slimer doesn`t return to stderr
func Exec(c *Command) (string, error) {
	var stdout, stderr bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "bash", "-c", c.Cmd)
	if len(c.Env) > 0 {
		cmd.Env = append(os.Environ(), c.Env...)
	}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		return "", fmt.Errorf("%s, %s", ErrExecutingCmd, err)
	}

	_ = cmd.Wait()
	status := cmd.ProcessState.Sys().(syscall.WaitStatus)

	switch status.ExitStatus() {
	case -1:
		return "", fmt.Errorf("%s: %s", ErrProcessKilled, errOut(stdout, stderr))
	case 1:
		return "", fmt.Errorf("%s: %s", ErrProcessExited, errOut(stdout, stderr))
	}

	return stdout.String(), nil
}

func errOut(stdout bytes.Buffer, stderr bytes.Buffer) string {
	if len(stderr.String()) < 1 {
		return stdout.String()
	}
	return stderr.String()
}
