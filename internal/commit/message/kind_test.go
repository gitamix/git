package message_test

import (
	"testing"

	"github.com/gitamix/types/commit"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/git/internal/commit/message"
)

func TestParser_IsRevert(t *testing.T) {
	t.Parallel()
	t.Run("git revert subject", func(t *testing.T) {
		t.Parallel()
		p := impl.NewParser(
			commit.ParseMessage([]byte(
				"Revert \"update docs\"\n\n" +
					"This reverts commit 1234567.\n",
			)),
		)
		assert.True(t, p.IsRevert())
	})
	t.Run("revert type with scope", func(t *testing.T) {
		t.Parallel()
		p := impl.NewParser(
			commit.ParseMessage([]byte(
				"revert(ui): change button alignment\n\n" +
					"Changed the alignment of the button back.\n",
			)),
		)
		assert.True(t, p.IsRevert())
	})
	t.Run("revert type oversized", func(t *testing.T) {
		t.Parallel()
		p := impl.NewParser(
			commit.ParseMessage([]byte("REVERT: update docs")),
		)
		assert.True(t, p.IsRevert())
	})
	t.Run("revert type without raw", func(t *testing.T) {
		t.Parallel()
		p := impl.NewParser(
			commit.NewMessage(
				commit.NewSubject(
					commit.NewType("revert"),
					commit.NewScope(""),
					commit.NewDescription("update docs"),
				),
				commit.NewBody(nil),
			),
		)
		assert.True(t, p.IsRevert())
	})
	t.Run("default subject", func(t *testing.T) {
		t.Parallel()
		p := impl.NewParser(
			commit.ParseMessage([]byte(
				"feat(ui): add new button\n\n" +
					"Added a new button to the UI.\n",
			)),
		)
		assert.False(t, p.IsRevert())
	})
	t.Run("merge subject", func(t *testing.T) {
		t.Parallel()
		p := impl.NewParser(
			commit.ParseMessage([]byte(
				"merge(ui): integrate 2 commits\n\n" +
					"Merged 2 feature commits into the UI.\n",
			)),
		)
		assert.False(t, p.IsRevert())
	})
	t.Run("empty message", func(t *testing.T) {
		t.Parallel()
		p := impl.NewParser(
			commit.ParseMessage([]byte("")),
		)
		assert.False(t, p.IsRevert())
	})
}
