package shell

import (
	"context"
	"time"

	"github.com/sitnikovik/osxec/command"
	proc "github.com/sitnikovik/osxec/process/execution"
)

// Execution executes the given command and returns its execution result.
func (s *Shell) Execution(
	ctx context.Context,
	cmd command.Command,
) proc.Execution {
	if res, ok := s.resp[cmd.String()]; ok {
		if s.duration > 0 {
			select {
			case <-ctx.Done():
				return proc.NewExecution(nil, ctx.Err())
			case <-time.After(s.duration):
			}
		}
		return res
	}
	panic("response not found for command '" + cmd.String() + "'")
}
