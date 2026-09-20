package git_test

import (
	"testing"

	"github.com/gitamix/types/commit"
	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/git/client/git"
)

func TestParams_Args(t *testing.T) {
	t.Parallel()
	type args struct {
		hash commit.Hash
		opts []impl.Option
	}
	type want struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "default",
			args: args{
				hash: commit.NewHash("1234567"),
				opts: nil,
			},
			want: want{
				args: []string{
					"rev-list",
					"--abbrev-commit",
					"1234567..HEAD",
				},
			},
		},
		{
			name: "first parent",
			args: args{
				hash: commit.NewHash("1234567"),
				opts: []impl.Option{
					impl.WithFirstParent(),
				},
			},
			want: want{
				args: []string{
					"rev-list",
					"--abbrev-commit",
					"--first-parent",
					"1234567..HEAD",
				},
			},
		},
		{
			name: "no merges",
			args: args{
				hash: commit.NewHash("1234567"),
				opts: []impl.Option{
					impl.WithNoMerges(),
				},
			},
			want: want{
				args: []string{
					"rev-list",
					"--abbrev-commit",
					"--no-merges",
					"1234567..HEAD",
				},
			},
		},
		{
			name: "first parent and no merges",
			args: args{
				hash: commit.NewHash("1234567"),
				opts: []impl.Option{
					impl.WithFirstParent(),
					impl.WithNoMerges(),
				},
			},
			want: want{
				args: []string{
					"rev-list",
					"--abbrev-commit",
					"--first-parent",
					"--no-merges",
					"1234567..HEAD",
				},
			},
		},
		{
			name: "empty hash",
			args: args{
				hash: commit.NewHash(""),
				opts: nil,
			},
			want: want{
				args: []string{
					"rev-list",
					"--abbrev-commit",
					"..HEAD",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := impl.
				NewParams(tt.args.opts...).
				Args(tt.args.hash)
			assert.Equal(
				t,
				tt.want.args,
				got,
			)
		})
	}
}
