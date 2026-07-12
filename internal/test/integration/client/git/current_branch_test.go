//go:build integration
// +build integration

package git_test

import (
	"context"
	"testing"
	"time"

	"github.com/gitamix/types/branch"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	impl "github.com/gitamix/git/client/git"
	shfx "github.com/gitamix/git/internal/test/fixture/shell"
)

// TestClient_CurrentBranch tests the CurrentBranch method of the Git client.
//
// It verifies that the method correctly retrieves the current branch name
// from the fixture repository and returns an error for a canceled context.
func TestClient_CurrentBranch(t *testing.T) {
	t.Parallel()
	t.Run("main branch", func(t *testing.T) {
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
			CurrentBranch(ctx)
		require.NoError(t, err)
		assert.Equal(
			t,
			branch.NewBranch(branch.NewName("main")),
			got,
		)
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
			CurrentBranch(ctx)
		assert.ErrorIs(t, err, context.Canceled)
		assert.Equal(t, branch.Branch{}, got)
	})
}
