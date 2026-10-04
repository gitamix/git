package message

import (
	"strings"

	"github.com/gitamix/types/commit"
)

// Parser parses a commit message to derive facts about it.
//
// Parser takes the message as-is
// and answers questions about the commit
// behind it without modifying it,
// such as whether the commit is a revert commit.
type Parser struct {
	// msg stores the commit message to parse.
	msg commit.Message
}

// NewParser creates a new Parser instance
// for the provided commit message.
func NewParser(msg commit.Message) Parser {
	return Parser{
		msg: msg,
	}
}

// IsRevert reports whether the parsed message
// belongs to a revert commit.
//
// The commit is a revert commit when the subject starts with "Revert "
// or when the commit type of the subject equals "revert" case-insensitively.
//
// The parser never reports merge commits,
// as merge commits are detected by the caller from the commit parents.
func (p Parser) IsRevert() bool {
	subj, _, _ := strings.Cut(p.msg.String(), "\n")
	return strings.HasPrefix(subj, "Revert ") ||
		strings.EqualFold(
			p.msg.
				Subject().
				Type().
				String(),
			"revert",
		)
}
