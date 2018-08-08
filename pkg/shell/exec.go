package shell

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/smilga/capture-go/pkg/logger"
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
	return ExecPipe([]*Command{c})
}

// ExecPipe executes multiple commands in pipe
func ExecPipe(commands []*Command) (string, error) {
	var stdout, stderr bytes.Buffer
	cmds := make([]*exec.Cmd, len(commands))
	info := make([]string, len(commands))

	for i, c := range commands {
		info[i] = fmt.Sprintf("%s | ", c.Cmd)
		name, arrgs := splitCmd(c.Cmd)
		cmd := exec.Command(name, arrgs...)
		if len(c.Env) > 0 {
			cmd.Env = append(os.Environ(), c.Env...)
		}
		cmds[i] = cmd
	}

	logger.Info(fmt.Sprintf("Executing commands: %s", info))

	// to kill process and all childs
	cmds[0].SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	time.AfterFunc(commands[0].Timeout, func() {
		syscall.Kill(-cmds[0].Process.Pid, syscall.SIGKILL)
	})

	// pipe aoutputs to inputs
	cmds[0].Stderr = &stdout
	for i, cmd := range cmds {
		if i < len(cmds)-1 {
			cmds[i+1].Stdin, _ = cmd.StdoutPipe()
			cmds[i+1].Stderr = &stderr
		} else {
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
		}
	}

	// start processes in descending order
	for i := len(cmds) - 1; i > 0; i-- {
		if err := cmds[i].Start(); err != nil {
			return stderr.String(), err
		}
	}
	// run the first process
	if err := cmds[0].Run(); err != nil {
		return stderr.String(), err
	}
	// wait on processes in ascending order
	for i := 1; i < len(cmds); i++ {
		if err := cmds[i].Wait(); err != nil {
			return stderr.String(), err
		}
	}

	return stdout.String(), nil
}

func splitCmd(s string) (string, []string) {
	split := strings.Split(s, " ")
	if len(split) > 1 {
		return split[0], split[1:]
	}
	return split[0], []string{}
}
