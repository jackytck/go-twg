package git

import (
	"os/exec"
	"strings"
)

var execCommand = exec.Command

func Version() string {
	cmd := execCommand("git", "version")
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	n := len("git version ")
	version := string(stdout[n:])
	return strings.TrimSpace(version)
}

type Checker struct {
	execCommand func(name string, args ...string) *exec.Cmd
}

func (c *Checker) command(name string, args ...string) *exec.Cmd {
	if c.execCommand == nil {
		return exec.Command(name, args...)
	}
	return c.execCommand(name, args...)
}

func (c *Checker) Version() string {
	cmd := c.command("git", "version")
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	n := len("git version ")
	version := string(stdout[n:])
	return strings.TrimSpace(version)
}
