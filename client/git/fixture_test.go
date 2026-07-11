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

// NewCommitsFromCmdFixture creates a new command fixture
// for retrieving commits by the specified commit hash.
//
// Used in tests to simulate the behavior of the git log command.
func NewCommitsFromCmdFixture(
	hash commit.Hash,
) command.Command {
	return command.NewCommand(
		"git",
		"rev-list",
		"--abbrev-commit",
		hash.String()+"..HEAD",
	)
}
