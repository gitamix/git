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

func TestClient_CommitMessage(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx  context.Context
		hash commit.Hash
	}
	type want struct {
		msg commit.Message
		err error
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
						NewCommitMessageCmdFixture(
							commit.NewHash("1234567"),
						),
						execution.NewExecution(
							[]byte(
								"feat(domain): add new feature\n\n"+
									"Added a new feature to the domain layer.\n\n"+
									"This feature allows users to perform advanced operations.\n",
							),
							nil,
						),
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				hash: commit.NewHash("1234567"),
			},
			want: want{
				msg: commit.NewMessage(
					commit.NewSubject(
						commit.NewType("feat"),
						commit.NewScope("domain"),
						commit.NewDescription("add new feature"),
					),
					commit.NewBody(
						[]byte(
							"Added a new feature to the domain layer.\n\n"+
								"This feature allows users to perform advanced operations.\n",
						),
					),
				),
				err: nil,
			},
		},
		{
			name: "failed to get",
			c: impl.NewClient(
				shellfk.NewShell(
					shellfk.WithResponse(
						NewCommitMessageCmdFixture(
							commit.NewHash("1234567"),
						),
						execution.NewExecution(
							nil,
							fake.Err,
						),
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				hash: commit.NewHash("1234567"),
			},
			want: want{
				msg: commit.Message{},
				err: errs.ErrGitFailed,
			},
		},
		{
			name: "got empty message",
			c: impl.NewClient(
				shellfk.NewShell(
					shellfk.WithResponse(
						NewCommitMessageCmdFixture(
							commit.NewHash("1234567"),
						),
						execution.NewExecution(
							[]byte(""),
							nil,
						),
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				hash: commit.NewHash("1234567"),
			},
			want: want{
				msg: commit.Message{},
				err: nil,
			},
		},
		{
			name: "failed to get",
			c: impl.NewClient(
				shellfk.NewShell(
					shellfk.WithResponse(
						NewCommitMessageCmdFixture(
							commit.NewHash("1234567"),
						),
						execution.NewExecution(
							nil,
							fake.Err,
						),
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				hash: commit.NewHash("1234567"),
			},
			want: want{
				msg: commit.Message{},
				err: errs.ErrGitFailed,
			},
		},
		{
			name: "empty hash",
			c: impl.NewClient(
				shellfk.NewShell(),
			),
			args: args{
				ctx:  context.Background(),
				hash: commit.NewHash(""),
			},
			want: want{
				msg: commit.Message{},
				err: errs.ErrEmptyHash,
			},
		},
		{
			name: "context canceled",
			c: impl.NewClient(
				shellfk.NewShell(
					shellfk.WithDuration(1*time.Second),
					shellfk.WithResponse(
						NewCommitMessageCmdFixture(
							commit.NewHash("1234567"),
						),
						execution.NewExecution(
							[]byte("feat(domain): add new feature\n"),
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
				hash: commit.NewHash("1234567"),
			},
			want: want{
				msg: commit.Message{},
				err: context.Canceled,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.c.CommitMessage(
				tt.args.ctx,
				tt.args.hash,
			)
			assert.Equalf(
				t,
				tt.want.msg,
				got,
				"got '%s', want '%s'",
				got,
				tt.want.msg,
			)
			assert.ErrorIs(t, err, tt.want.err)
		})
	}
}
