package git_test

import (
	"context"
	"testing"
	"time"

	"github.com/gitamix/types/branch"
	"github.com/sitnikovik/osxec/process/execution"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/git/client/git"
	"github.com/gitamix/git/errs"
	"github.com/gitamix/git/internal/test/fake"
	shellfk "github.com/gitamix/git/internal/test/fake/shell"
)

func TestClient_CurrentBranch(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
	}
	type want struct {
		branch branch.Branch
		err    error
	}
	tests := []struct {
		name string
		c    *impl.Client
		args args
		want want
	}{
		{
			name: "ok",
			c: impl.NewClient(
				shellfk.NewShell(
					shellfk.WithResponse(
						NewCurrentBranchCmdFixture(),
						execution.NewExecution(
							[]byte("main\n"),
							nil,
						),
					),
				),
			),
			args: args{
				ctx: context.Background(),
			},
			want: want{
				branch: branch.NewBranch(
					branch.NewName("main"),
				),
				err: nil,
			},
		},
		{
			name: "failed to get current branch",
			c: impl.NewClient(
				shellfk.NewShell(
					shellfk.WithResponse(
						NewCurrentBranchCmdFixture(),
						execution.NewExecution(
							nil,
							fake.Err,
						),
					),
				),
			),
			args: args{
				ctx: context.Background(),
			},
			want: want{
				branch: branch.Branch{},
				err:    errs.ErrGitFailed,
			},
		},
		{
			name: "ctx canceled",
			c: impl.NewClient(
				shellfk.NewShell(
					shellfk.WithDuration(1*time.Second),
					shellfk.WithResponse(
						NewCurrentBranchCmdFixture(),
						execution.NewExecution(
							[]byte("main\n"),
							nil,
						),
					),
				),
			),
			args: args{
				ctx: func() context.Context {
					ctx, cancel := context.WithCancel(context.Background())
					cancel()
					return ctx
				}(),
			},
			want: want{
				branch: branch.Branch{},
				err:    context.Canceled,
			},
		},
		{
			name: "empty branch name",
			c: impl.NewClient(
				shellfk.NewShell(
					shellfk.WithResponse(
						NewCurrentBranchCmdFixture(),
						execution.NewExecution(
							[]byte("\n"),
							nil,
						),
					),
				),
			),
			args: args{
				ctx: context.Background(),
			},
			want: want{
				branch: branch.NewBranch(branch.NewName("")),
				err:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.c.CurrentBranch(
				tt.args.ctx,
			)
			assert.Equal(t, tt.want.branch, got)
			assert.ErrorIs(t, err, tt.want.err)
		})
	}
}
