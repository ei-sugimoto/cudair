package executor

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Execute(bin string) error {
	parts := strings.Fields(bin)
	runCmd := exec.Command(parts[0], parts[1:]...)
	log.Println("Running command:", runCmd.String())
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr
	if err := runCmd.Run(); err != nil {
		return fmt.Errorf("command execution failed: %w", err)
	}

	log.Println("Command finished successfully")
	return nil
}
