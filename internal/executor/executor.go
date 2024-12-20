package executor

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func Execute(bin string) error {
	parts := strings.Fields(bin)
	runCmd := exec.Command(parts[0], parts[1:]...)

	log.Println("execting...")
	file, err := os.Create("./tmp/result.log")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()

	runCmd.Stdout = file
	runCmd.Stderr = file

	if err := runCmd.Run(); err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return nil
}
