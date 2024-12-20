package builder_test

import (
	"testing"

	"github.com/ei-sugimoto/cudair/internal/builder"
)

func TestBuild_ValidCommand(t *testing.T) {
	t.Parallel()
	err := builder.Build("echo")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestBuild_InvalidCommand(t *testing.T) {
	t.Parallel()
	err := builder.Build("invalidcommand")
	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}
}

func TestBuild_CommandWithError(t *testing.T) {
	t.Parallel()
	err := builder.Build("false") // 'false' is a command that always returns an error
	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}
}
