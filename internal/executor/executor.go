package executor

import (
	"os"
	"os/exec"
)

func Execute(bin string) error {
	runCmd := exec.Command(bin)
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr

	return runCmd.Run()
}
