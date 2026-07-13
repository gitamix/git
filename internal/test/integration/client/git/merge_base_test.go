//go:build integration
// +build integration

package git_test

import (
	"context"
	"testing"
	"time"

	"github.com/gitamix/types/commit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	impl "github.com/gitamix/git/client/git"
	"github.com/gitamix/git/errs"
	shfx "github.com/gitamix/git/internal/test/fixture/shell"
)

// TestClient_MergeBase tests the MergeBase method of the Git client.
//
// It verifies that the method correctly finds the common ancestor
// between two branches or commits in a real Git repository.
func TestClient_MergeBase(t *testing.T) {
	t.Parallel()
	t.Run("merge base between main and feature-merge branches", func(t *testing.T) {
		t.Parallel()
		fx := sharedContainerFixture(t)
		ctx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()
		got, err := impl.
			NewClient(
				shfx.NewShell(
					fx.Container(),
					RepoDir,
				),
			).
			MergeBase(
				ctx,
				"main",
				"feature-merge",
			)
		require.NoError(t, err)
		assert.NotEmpty(t, got)
	})

	t.Run("merge base between two commits", func(t *testing.T) {
		t.Parallel()
		fx := sharedContainerFixture(t)
		ctx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()
		// Get merge base between the first and third main commits
		got, err := impl.
			NewClient(
				shfx.NewShell(
					fx.Container(),
					RepoDir,
				),
			).
			MergeBase(
				ctx,
				fx.Env().MustGet("MAIN_FIRST_PARENT_01"),
				fx.Env().MustGet("MAIN_FIRST_PARENT_03"),
			)
		require.NoError(t, err)
		want := commit.NewHash(
			fx.Env().MustGet("MAIN_FIRST_PARENT_01"),
		)
		assert.Equal(t, want, got)
	})

	t.Run("merge base between main and release branches", func(t *testing.T) {
		t.Parallel()
		fx := sharedContainerFixture(t)
		ctx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()
		got, err := impl.
			NewClient(
				shfx.NewShell(
					fx.Container(),
					RepoDir,
				),
			).
			MergeBase(
				ctx,
				"main",
				"release/1.0.0",
			)
		require.NoError(t, err)
		assert.NotEmpty(t, got)
	})

	t.Run("merge base with non-existent branch", func(t *testing.T) {
		t.Parallel()
		fx := sharedContainerFixture(t)
		ctx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()
		got, err := impl.
			NewClient(
				shfx.NewShell(
					fx.Container(),
					RepoDir,
				),
			).
			MergeBase(
				ctx,
				"main",
				"nonexistent-branch",
			)
		assert.ErrorIs(t, err, errs.ErrGitFailed)
		assert.Empty(t, got)
	})

	t.Run("canceled context", func(t *testing.T) {
		t.Parallel()
		fx := sharedContainerFixture(t)
		ctx, cancel := context.WithCancel(
			context.Background(),
		)
		cancel()
		got, err := impl.
			NewClient(
				shfx.NewShell(
					fx.Container(),
					RepoDir,
				),
			).
			MergeBase(
				ctx,
				"main",
				"feature-merge",
			)
		assert.ErrorIs(t, err, context.Canceled)
		assert.Empty(t, got)
	})
}
