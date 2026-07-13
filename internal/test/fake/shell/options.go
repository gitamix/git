package shell

import (
	"time"

	"github.com/sitnikovik/osxec/command"
	proc "github.com/sitnikovik/osxec/process/execution"
)

// Option defines a function type for configuring the Shell instance.
type Option func(*Shell)

// WithResponse sets a predefined response for a specific command.
//
// Parameters:
//   - cmd: The command for which the response is to be set.
//   - res: The execution result to be returned when the command is executed.
func WithResponse(
	cmd command.Command,
	res proc.Execution,
) Option {
	return func(s *Shell) {
		s.resp[cmd.String()] = res
	}
}

// WithDuration sets the simulated execution duration for commands.
func WithDuration(d time.Duration) Option {
	return func(s *Shell) {
		s.duration = d
	}
}
