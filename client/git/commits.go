package git

import (
	"context"
	"errors"
	"fmt"

	"github.com/gitamix/types/commit"
	"github.com/sitnikovik/osxec/command"
	"github.com/sitnikovik/osxec/process"

	"github.com/gitamix/git/errs"
)

// Commits retrieves a list of commits
// from the specified hash to the current HEAD.
//
// Returns error if the hash is empty
// or if the git command execution fails
// or if the context is canceled.
func (c *Client) Commits(
	ctx context.Context,
	hash commit.Hash,
) ([]commit.Commit, error) {
	if hash.Empty() {
		return nil, errs.ErrEmptyHash
	}
	res := process.
		NewProcess(
			c.shell,
			command.NewCommand(
				"git",
				"rev-list",
				"--abbrev-commit",
				hash.String()+"..HEAD",
			),
		).
		Execution(ctx)
	if err := res.Err(); err != nil {
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return nil, err
		}
		return nil, errors.Join(err, errs.ErrGitFailed)
	}
	out := res.Output()
	commits := make([]commit.Commit, 0, out.Len())
	for _, ln := range out.Lines() {
		if ln == "" {
			continue
		}
		h := commit.NewHash(ln)
		msg, err := c.CommitMessage(ctx, h)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to get commit message for %s: %w",
				h.String(),
				err,
			)
		}
		commits = append(commits, commit.NewCommit(h, msg))
	}
	return commits, nil
}
