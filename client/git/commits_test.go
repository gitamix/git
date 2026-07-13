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

func TestClient_Commits(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx  context.Context
		hash commit.Hash
	}
	type want struct {
		commits []commit.Commit
		err     error
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
						NewCommitsCmdFixture(
							commit.NewHash("1234567"),
						),
						execution.NewExecution(
							[]byte(
								"89abcdef\n"+
									"1234567\n",
							),
							nil,
						),
					),
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
					shellfk.WithResponse(
						NewCommitMessageCmdFixture(
							commit.NewHash("89abcdef"),
						),
						execution.NewExecution(
							[]byte(
								"fix(ui): resolve button alignment issue\n\n"+
									"Fixed the alignment issue of the button in the UI.\n\n"+
									"This fix ensures consistent button placement across different screen sizes.\n",
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
				commits: []commit.Commit{
					commit.NewCommit(
						commit.NewHash("89abcdef"),
						commit.NewMessage(
							commit.NewSubject(
								commit.NewType("fix"),
								commit.NewScope("ui"),
								commit.NewDescription("resolve button alignment issue"),
							),
							commit.NewBody(
								[]byte(
									"Fixed the alignment issue of the button in the UI.\n\n"+
										"This fix ensures consistent button placement across different screen sizes.\n",
								),
							),
						),
					),
					commit.NewCommit(
						commit.NewHash("1234567"),
						commit.NewMessage(
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
					),
				},
				err: nil,
			},
		},
		{
			name: "failed to get commits",
			c: impl.NewClient(
				shellfk.NewShell(
					shellfk.WithResponse(
						NewCommitsCmdFixture(
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
				commits: nil,
				err:     errs.ErrGitFailed,
			},
		},
		{
			name: "failed to get commit message",
			c: impl.NewClient(
				shellfk.NewShell(
					shellfk.WithResponse(
						NewCommitsCmdFixture(
							commit.NewHash("1234567"),
						),
						execution.NewExecution(
							[]byte(
								"89abcdef\n"+
									"1234567\n",
							),
							nil,
						),
					),
					shellfk.WithResponse(
						NewCommitMessageCmdFixture(
							commit.NewHash("1234567"),
						),
						execution.NewExecution(
							nil,
							fake.Err,
						),
					),
					shellfk.WithResponse(
						NewCommitMessageCmdFixture(
							commit.NewHash("89abcdef"),
						),
						execution.NewExecution(
							[]byte(
								"fix(ui): resolve button alignment issue\n\n"+
									"Fixed the alignment issue of the button in the UI.\n\n"+
									"This fix ensures consistent button placement across different screen sizes.\n",
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
				commits: nil,
				err:     fake.Err,
			},
		},
		{
			name: "ctx cancelled on getting commits",
			c: impl.NewClient(
				shellfk.NewShell(
					shellfk.WithResponse(
						NewCommitsCmdFixture(
							commit.NewHash("1234567"),
						),
						execution.NewExecution(
							nil,
							fake.Err,
						),
					),
					shellfk.WithDuration(1*time.Second),
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
				commits: nil,
				err:     context.Canceled,
			},
		},
		{
			name: "empty hash",
			c: impl.NewClient(
				shellfk.NewShell(),
			),
			args: args{
				ctx: func() context.Context {
					ctx, cancel := context.WithCancel(context.Background())
					cancel()
					return ctx
				}(),
				hash: commit.NewHash(""),
			},
			want: want{
				commits: nil,
				err:     errs.ErrEmptyHash,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.c.Commits(
				tt.args.ctx,
				tt.args.hash,
			)
			assert.Equal(t, tt.want.commits, got)
			assert.ErrorIs(t, err, tt.want.err)
		})
	}
}
