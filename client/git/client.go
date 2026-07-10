package git

import (
	"github.com/sitnikovik/osxec/process"
)

// Client defines the client to interact with git repositories.
type Client struct {
	// shell is the shell instance used
	// for executing system commands.
	shell process.Shell
}

// NewClient creates a new Client instance
// with the provided shell to interact with git repositories.
func NewClient(exec process.Shell) *Client {
	return &Client{
		shell: exec,
	}
}
