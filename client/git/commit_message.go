package git

import (
	"context"
	"errors"

	"github.com/gitamix/types/commit"
	"github.com/sitnikovik/osxec/command"
	"github.com/sitnikovik/osxec/process"

	"github.com/gitamix/git/errs"
)

// CommitMessage retrieves the commit message for a given commit hash.
//
// Returns error if the hash is empty
// or if the git command execution fails
// or if the context is canceled.
func (c *Client) CommitMessage(
	ctx context.Context,
	hash commit.Hash,
) (commit.Message, error) {
	if hash.Empty() {
		return commit.Message{}, errs.ErrEmptyHash
	}
	res := process.
		NewProcess(
			c.shell,
			command.NewCommand(
				"git",
				"log",
				"-1",
				"--pretty=%B",
				hash.String(),
			),
		).
		Execution(ctx)
	if err := res.Err(); err != nil {
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return commit.Message{}, err
		}
		return commit.Message{}, errors.Join(err, errs.ErrGitFailed)
	}
	return commit.ParseMessage(res.Output().Bytes()), nil
}
