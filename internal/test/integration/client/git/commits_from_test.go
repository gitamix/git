//go:build integration
// +build integration

package git_test

import (
	"context"
	"testing"

	"github.com/gitamix/types/commit"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/git/client/git"
	"github.com/gitamix/git/errs"
	shfx "github.com/gitamix/git/internal/test/fixture/shell"
)

// TestClient_CommitsFrom tests the CommitsFrom method of the Git client.
//
// It verifies that the method correctly retrieves commits
// from the specified hash to HEAD and returns appropriate errors
// for invalid inputs.
func TestClient_CommitsFrom(t *testing.T) {
	t.Parallel()
	t.Run("main first parent commit", func(t *testing.T) {
		t.Parallel()
		fx := sharedContainerFixture(t)
		ctx := context.Background()
		got, err := impl.
			NewClient(
				shfx.NewShell(
					fx.Container(),
					RepoDir,
				),
			).
			CommitsFrom(
				ctx,
				commit.NewHash(
					fx.
						Env().
						MustGet("MAIN_FIRST_PARENT_01"),
				),
			)
		assert.NoError(t, err)
		assert.Len(t, got, 16)
	})

	t.Run("not found commit", func(t *testing.T) {
		t.Parallel()
		fx := sharedContainerFixture(t)
		ctx := context.Background()
		got, err := impl.
			NewClient(
				shfx.NewShell(
					fx.Container(),
					RepoDir,
				),
			).
			CommitsFrom(
				ctx,
				commit.NewHash("thishashdoesnotexist"),
			)
		assert.ErrorIs(t, err, errs.ErrGitFailed)
		assert.Nil(t, got)
	})

	t.Run("empty commit hash", func(t *testing.T) {
		t.Parallel()
		fx := sharedContainerFixture(t)
		ctx := context.Background()
		got, err := impl.
			NewClient(
				shfx.NewShell(
					fx.Container(),
					RepoDir,
				),
			).
			CommitsFrom(
				ctx,
				commit.NewHash(""),
			)
		assert.ErrorIs(t, err, errs.ErrEmptyHash)
		assert.Nil(t, got)
	})
}
