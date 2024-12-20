package builder

import (
	"os"
	"os/exec"
)

func Build(cmd string) error {
	buildCmd := exec.Command(cmd)
	buildCmd.Stderr = os.Stderr
	buildCmd.Stdout = os.Stdout
	return buildCmd.Run()
}
