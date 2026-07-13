package git

import (
	"context"
	"errors"

	"github.com/gitamix/types/branch"
	"github.com/sitnikovik/osxec/command"
	"github.com/sitnikovik/osxec/process"

	"github.com/gitamix/git/errs"
)

// CurrentBranch retrieves the current branch
// or returns an error if the git command execution fails.
func (c *Client) CurrentBranch(
	ctx context.Context,
) (branch.Branch, error) {
	res := process.
		NewProcess(
			c.shell,
			command.NewCommand(
				"git",
				"rev-parse",
				"--abbrev-ref",
				"HEAD",
			),
		).
		Execution(ctx)
	if err := res.Err(); err != nil {
		if errs.IsContextError(err) {
			return branch.Branch{}, err
		}
		return branch.Branch{}, errors.Join(err, errs.ErrGitFailed)
	}
	return branch.NewBranch(
		branch.NewName(
			res.Output().
				Lines().
				First(),
		),
	), nil
}
