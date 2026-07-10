package errs

import "errors"

var (
	// ErrEmptyHash is returned when an empty commit hash is provided.
	ErrEmptyHash = errors.New("empty hash provided")
	// ErrGitFailed is returned when a git command execution fails.
	ErrGitFailed = errors.New("git command failed")
)
