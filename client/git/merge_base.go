package git

import (
	"context"
	"errors"

	"github.com/gitamix/types/commit"
	"github.com/sitnikovik/osxec/command"
	"github.com/sitnikovik/osxec/process"

	"github.com/gitamix/git/errs"
)

// MergeBase retrieves the merge base commit hash between two specified commits or branches.
//
// For example, it can be used to find the common ancestor
// of two branches and/or commits in a Git repository.
//
// Returns error if the git command execution fails
// or if the context is canceled.
func (c *Client) MergeBase(
	ctx context.Context,
	target string,
	curr string,
) (commit.Hash, error) {
	res := process.
		NewProcess(
			c.shell,
			command.NewCommand(
				"git",
				"merge-base",
				target,
				curr,
			),
		).
		Execution(ctx)
	if err := res.Err(); err != nil {
		if errs.IsContextError(err) {
			return "", err
		}
		return "", errors.Join(err, errs.ErrGitFailed)
	}
	return commit.NewHash(
		res.Output().
			Lines().
			First(),
	), nil
}
