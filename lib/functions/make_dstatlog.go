package functions

import "os/exec"

func MakeDstatLog() *exec.Cmd {
	cmd := exec.Command("dstat", "-nr", "--output", "/home/ansible/dstatlog.csv")
	cmd.Start()
	return cmd
}
