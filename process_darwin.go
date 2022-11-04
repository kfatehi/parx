package main

import (
	"os/exec"
	"syscall"
)

type DarwinProcess struct {
	cmd *exec.Cmd
}

func NewProcessKillable(cmd *exec.Cmd) ProcessKillable {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	return &DarwinProcess{cmd: cmd}
}

// Kill kills the process and its children.
//
// In Linux we can get the process group id of the process
// and kill it that way, to make sure every child is killed as well.
func (l *DarwinProcess) Kill() error {
	pgid, err := syscall.Getpgid(l.cmd.Process.Pid)
	if err != nil {
		return err
	}
	return syscall.Kill(-pgid, syscall.SIGKILL)
}
