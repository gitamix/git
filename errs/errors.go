package errs

import (
	"context"
	"errors"
)

var (
	// ErrEmptyHash is returned when an empty commit hash is provided.
	ErrEmptyHash = errors.New("empty hash provided")
	// ErrGitFailed is returned when a git command execution fails.
	ErrGitFailed = errors.New("git command failed")
)

// IsContextError checks if the provided error is a type of any context error.
func IsContextError(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, context.Canceled) ||
		errors.Is(err, context.DeadlineExceeded)
}
