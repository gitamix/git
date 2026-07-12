package errs_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/git/errs"
)

func TestIsContextError(t *testing.T) {
	t.Parallel()
	type args struct {
		err error
	}
	type want struct {
		is bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "context canceled",
			args: args{
				err: context.Canceled,
			},
			want: want{
				is: true,
			},
		},
		{
			name: "context deadline exceeded",
			args: args{
				err: context.DeadlineExceeded,
			},
			want: want{
				is: true,
			},
		},
		{
			name: "non-context error",
			args: args{
				err: errs.ErrGitFailed,
			},
			want: want{
				is: false,
			},
		},
		{
			name: "nil error",
			args: args{
				err: nil,
			},
			want: want{
				is: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := errs.IsContextError(tt.args.err)
			if tt.want.is {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}
