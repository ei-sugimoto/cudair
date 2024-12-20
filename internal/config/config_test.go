package config_test

import (
	"os"
	"testing"

	"github.com/ei-sugimoto/cudair/internal/config"
)

func TestNewCudairConfig_ValidFile(t *testing.T) {
	tomlContent := `
root = "/some/root"
tmp_dir = "/some/tmp"
[build]
bin = "/some/bin"
cmd = "some command"
log = "/some/log"
`
	file, err := os.CreateTemp("", "config-*.toml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	if _, err := file.Write([]byte(tomlContent)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	file.Close()

	config, err := config.NewCudairConfig(file.Name())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if config.Root != "/some/root" {
		t.Errorf("Expected root to be '/some/root', got %v", config.Root)
	}
	if config.TmpDir != "/some/tmp" {
		t.Errorf("Expected tmp_dir to be '/some/tmp', got %v", config.TmpDir)
	}
	if config.Build.Bin != "/some/bin" {
		t.Errorf("Expected build.bin to be '/some/bin', got %v", config.Build.Bin)
	}
	if config.Build.Cmd != "some command" {
		t.Errorf("Expected build.cmd to be 'some command', got %v", config.Build.Cmd)
	}
	if config.Build.Log != "/some/log" {
		t.Errorf("Expected build.log to be '/some/log', got %v", config.Build.Log)
	}
}

func TestNewCudairConfig_InvalidFile(t *testing.T) {
	tomlContent := `
root = "/some/root"
tmp_dir = "/some/tmp"
[build]
bin = "/some/bin"
cmd = "some command"
log = "/some/log
`
	file, err := os.CreateTemp("", "config-*.toml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	if _, err := file.Write([]byte(tomlContent)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	file.Close()

	_, err = config.NewCudairConfig(file.Name())
	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}
}

func TestNewCudairConfig_NonExistentFile(t *testing.T) {
	_, err := config.NewCudairConfig("non-existent-file.toml")
	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}
}
