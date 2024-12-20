package executor_test

import (
	"testing"

	"github.com/ei-sugimoto/cudair/internal/executor"
)

func TestExecute_ValidCommand(t *testing.T) {
	t.Parallel()
	err := executor.Execute("echo")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestExecute_InvalidCommand(t *testing.T) {
	t.Parallel()
	err := executor.Execute("invalidcommand")
	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}
}

func TestExecute_CommandWithError(t *testing.T) {
	t.Parallel()
	err := executor.Execute("false") // 'false' is a command that always returns an error
	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}
}
