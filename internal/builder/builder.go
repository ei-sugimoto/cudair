package builder

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func Build(cmd string, tmpDirPath string) error {
	if err := os.MkdirAll(tmpDirPath, os.ModePerm); err != nil {
		return err
	}
	parts := strings.Fields(cmd)
	buildCmd := exec.Command(parts[0], parts[1:]...)
	buildCmd.Stderr = os.Stderr
	buildCmd.Stdout = os.Stdout
	log.Println("building...")
	startTime := time.Now()
	if err := buildCmd.Run(); err != nil {
		return err
	}
	log.Printf("build time: %v\n", time.Since(startTime))

	return nil
}
