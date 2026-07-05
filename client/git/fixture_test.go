package git_test

import (
	"github.com/gitamix/types/commit"
	"github.com/sitnikovik/osxec/command"
)

// NewCommitMessageCmdFixture creates a new command fixture
// for retrieving the commit message by the specified commit hash.
//
// Returns the instance to be used
// in tests to simulate the behavior of the git log command.
func NewCommitMessageCmdFixture(
	hash commit.Hash,
) command.Command {
	return command.NewCommand(
		"git",
		"log",
		"-1",
		"--pretty=%B",
		hash.String(),
	)
}
