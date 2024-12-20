package builder

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func Build(cmd string, tmpDirPath string) error {
	log.Println("creating tmp dir")
	if err := os.MkdirAll(tmpDirPath, os.ModePerm); err != nil {
		return err
	}
	parts := strings.Fields(cmd)
	buildCmd := exec.Command(parts[0], parts[1:]...)
	buildCmd.Stderr = os.Stderr
	buildCmd.Stdout = os.Stdout
	log.Println("building...")
	return buildCmd.Run()
}
