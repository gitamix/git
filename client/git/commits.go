package git

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gitamix/types/commit"
	"github.com/sitnikovik/osxec/command"
	"github.com/sitnikovik/osxec/process"

	"github.com/gitamix/git/errs"
	"github.com/gitamix/git/internal/commit/message"
)

// Commits retrieves a list of commits reachable from the current HEAD
// that are not reachable from the specified hash (i.e. commits after `hash`, exclusive).
//
// Each provided commit carries its message
// and the kind parsed from the message and parents:
// a merge commit, a revert commit, or a default commit.
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
				"--parents",
				hash.String()+"..HEAD",
			),
		).
		Execution(ctx)
	if err := res.Err(); err != nil {
		if errs.IsContextError(err) {
			return nil, err
		}
		return nil, errors.Join(err, errs.ErrGitFailed)
	}
	out := res.Output()
	commits := make([]commit.Commit, 0, out.Len())
	for _, ln := range out.Lines() {
		parents := strings.Fields(ln)
		if len(parents) == 0 {
			continue
		}
		h := commit.NewHash(parents[0])
		msg, err := c.CommitMessage(ctx, h)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to get commit message for %s: %w",
				h.String(),
				err,
			)
		}
		parser := message.NewParser(msg)
		var kind commit.Kind
		if len(parents) > 2 {
			kind = commit.KindMerge
		} else if parser.IsRevert() {
			kind = commit.KindRevert
		}
		commits = append(
			commits,
			commit.NewCommit(
				h,
				msg,
				commit.WithKind(kind),
			),
		)
	}
	return commits, nil
}
