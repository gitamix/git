package shell

import (
	"time"

	proc "github.com/sitnikovik/osxec/process/execution"
)

// Shell is a fake implementation of the Shell interface for testing purposes.
type Shell struct {
	// resp holds predefined responses for specific commands.
	resp map[string]proc.Execution
	// duration simulates the execution duration for commands.
	duration time.Duration
}

// NewShell creates and returns a new Shell instance to execute system commands.
func NewShell(opts ...Option) *Shell {
	sh := &Shell{
		resp: make(map[string]proc.Execution),
	}
	for _, opt := range opts {
		opt(sh)
	}
	return sh
}
