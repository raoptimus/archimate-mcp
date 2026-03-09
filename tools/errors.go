package tools

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNotFound     = errors.New("resource not found")
	ErrInvalidInput = errors.New("invalid input")
	ErrUnavailable  = errors.New("archi server unavailable")
)

// WrapError wraps an error with operation context.
func WrapError(operation string, err error) error {
	if err == nil {
		return nil
	}

	errMsg := err.Error()

	switch {
	case strings.Contains(errMsg, "not found"):
		return fmt.Errorf("%s: %w", operation, ErrNotFound)
	case strings.Contains(errMsg, "connection refused") || strings.Contains(errMsg, "dial tcp"):
		return fmt.Errorf("%s: %w: is jArchi HTTP server running?", operation, ErrUnavailable)
	case strings.Contains(errMsg, "required") || strings.Contains(errMsg, "invalid"):
		return fmt.Errorf("%s: %w: %s", operation, ErrInvalidInput, errMsg)
	default:
		return fmt.Errorf("%s: %w", operation, err)
	}
}

// FormatToolError formats error for MCP tool response.
func FormatToolError(err error) string {
	if err == nil {
		return ""
	}

	switch {
	case errors.Is(err, ErrNotFound):
		return fmt.Sprintf("Resource not found. %v", err)
	case errors.Is(err, ErrUnavailable):
		return fmt.Sprintf("Archi server unavailable. Make sure ArchiMCPServer.ajs is running in Archi. %v", err)
	case errors.Is(err, ErrInvalidInput):
		return fmt.Sprintf("Invalid input: %v", err)
	default:
		return fmt.Sprintf("Operation failed: %v", err)
	}
}
