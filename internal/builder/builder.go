package builder

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func Build(cmd string, tmpDirPath string, errLogFileName string) error {
	errLogFilePath := filepath.Join(tmpDirPath, errLogFileName)
	if err := os.MkdirAll(tmpDirPath, os.ModePerm); err != nil {
		return err
	}
	parts := strings.Fields(cmd)
	buildCmd := exec.Command(parts[0], parts[1:]...)
	buildCmd.Stdout = os.Stdout
	captureErr := captureCmdErr(buildCmd)

	log.Println("building...")
	startTime := time.Now()

	if err := buildCmd.Run(); err != nil {
		log.Println("build failed")
		originalErr := err
		f, err := os.Create(errLogFilePath)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.WriteString(captureErr.String())
		if err != nil {
			return err
		}

		return originalErr
	}

	_, err := os.Stat(errLogFilePath)
	if err == nil {
		if err := os.Remove(errLogFilePath); err != nil {
			log.Println("failed to remove error log file")
			return nil
		}
	}

	log.Printf("build time: %v\n", time.Since(startTime))
	return nil
}

func captureCmdErr(cmd *exec.Cmd) *strings.Builder {
	var stderr strings.Builder
	cmd.Stderr = &stderr
	return &stderr
}
