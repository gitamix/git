package git_test

import (
	"context"
	"testing"
	"time"

	"github.com/gitamix/types/commit"
	"github.com/sitnikovik/osxec/process/execution"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/git/client/git"
	"github.com/gitamix/git/errs"
	"github.com/gitamix/git/internal/test/fake"
	shellfk "github.com/gitamix/git/internal/test/fake/shell"
)

func TestClient_MergeBase(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		target string
		curr   string
	}
	type want struct {
		hash commit.Hash
		err  error
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
						NewMergeBaseCmdFixture(
							"main",
							"feature",
						),
						execution.NewExecution(
							[]byte("abc1234\n"),
							nil,
						),
					),
				),
			),
			args: args{
				ctx:    context.Background(),
				target: "main",
				curr:   "feature",
			},
			want: want{
				hash: commit.NewHash("abc1234"),
				err:  nil,
			},
		},
		{
			name: "failed to get merge base",
			c: impl.NewClient(
				shellfk.NewShell(
					shellfk.WithResponse(
						NewMergeBaseCmdFixture(
							"main",
							"feature",
						),
						execution.NewExecution(
							nil,
							fake.Err,
						),
					),
				),
			),
			args: args{
				ctx:    context.Background(),
				target: "main",
				curr:   "feature",
			},
			want: want{
				hash: "",
				err:  errs.ErrGitFailed,
			},
		},
		{
			name: "context canceled",
			c: impl.NewClient(
				shellfk.NewShell(
					shellfk.WithDuration(1*time.Second),
					shellfk.WithResponse(
						NewMergeBaseCmdFixture(
							"main",
							"feature",
						),
						execution.NewExecution(
							nil,
							context.Canceled,
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
				target: "main",
				curr:   "feature",
			},
			want: want{
				hash: "",
				err:  context.Canceled,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			hash, err := tt.c.MergeBase(
				tt.args.ctx,
				tt.args.target,
				tt.args.curr,
			)
			assert.Equal(t, tt.want.hash, hash)
			assert.ErrorIs(t, err, tt.want.err)
		})
	}
}
