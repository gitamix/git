# git

[![Go Reference](https://pkg.go.dev/badge/github.com/gitamix/git.svg)](https://pkg.go.dev/github.com/gitamix/git)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gitamix/git)
[![Release](https://img.shields.io/github/v/release/gitamix/git?style=flat)](https://github.com/gitamix/git/releases)

A Go library with a client for interacting with Git repositories.
It executes git commands through a pluggable shell
and returns git domain types from
[gitamix/types](gitamix/types) —
so instead of parsing raw git output yourself,
you get typed branches, commits, and messages
ready to pass to other tools.

> [!NOTE]
> The project is not a CLI tool — it is a Go library you embed into your own tools.
> It never hard-codes how commands are executed:
> any `process.Shell` implementation can be plugged in —
> the local shell from
> [sitnikovik/osxec](https://github.com/sitnikovik/osxec),
> a Docker container, or a fake in tests.

## Features

- **Branch retrieval** — `CurrentBranch` returns the current branch
  as a typed `branch.Branch`.
- **Commit access** — `CommitMessage` returns the parsed `commit.Message`
  for a given hash; `Commits` returns every commit reachable from `HEAD`
  but not from the given hash — for example, all commits after the merge base.
- **Merge base** — `MergeBase` finds the common ancestor
  of two branches or commits, e.g. to diff a feature branch against `main`.
- **Typed results** — every method returns domain types from
  [gitamix/types](https://github.com/gitamix/types)
  (`Branch`, `Commit`, `Hash`, `Message`),
  ready to pass to other gitamix tools.
- **Pluggable shell** — git commands run through the `process.Shell` interface
  from [sitnikovik/osxec](https://github.com/sitnikovik/osxec):
  use the built-in local shell, execute inside a Docker container,
  or substitute a fake in tests.
- **Context-aware** — every call accepts a `context.Context`
  and respects cancellation and deadlines.
- **Typed errors** — shared domain errors `ErrEmptyHash` and `ErrGitFailed`,
  plus `errs.IsContextError` to distinguish context failures.

## Installation

Run the command in terminal:

```sh
go get github.com/gitamix/git
```

Both runtime dependencies —
[gitamix/types](https://github.com/gitamix/types) for the git domain types
and [/sitnikovik/osxec](https://github.com/sitnikovik/osxec)
for command execution — are pulled in automatically.

## Getting started

Create the client with any shell implementation
and query the repository you are in:

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/sitnikovik/osxec/shell"

    "github.com/gitamix/git/client/git"
)

func main() {
    c := git.NewClient(shell.NewShell())
    ctx := context.Background()

    // Current branch.
    br, err := c.CurrentBranch(ctx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("branch:", br)

    // Commits between the merge base and HEAD.
    base, err := c.MergeBase(ctx, "origin/main", br.String())
    if err != nil {
        log.Fatal(err)
    }
    cmts, err := c.Commits(ctx, base)
    if err != nil {
        log.Fatal(err)
    }
    for _, cm := range cmts {
        fmt.Println(
            cm.Hash().ShortString(), 
            m.Message().Subject(),
        )
    }
}
```

## Documentation

Full API reference is available on [pkg.go.dev](https://pkg.go.dev/github.com/gitamix/git)

## Ecosystem

The client is the execution layer of the
[gitamix](https://github.com/gitamix) tooling:
it fetches real branches and commits as
[gitamix/types](https://github.com/gitamix/types) values,
which every gitamix project consumes
instead of running git commands itself —
for example, [gitamix/lint](https://github.com/gitamix/lint)
lints the values this client retrieves.

## Requirements

- [Go](https://go.dev/) 1.25.7 or later.
- Git — commands are executed through the shell,
  so the `git` binary must be resolvable in the execution environment.
- [gitamix/types](https://github.com/gitamix/types) —
  git domain types returned by the client.
- [sitnikovik/osxec](https://github.com/sitnikovik/osxec) —
  command execution abstraction the client is built on.

The client does not decide where git commands run:
it delegates every command to the `process.Shell`
passed to `git.NewClient` — the local shell from `osxec`,
a Docker container, or your own implementation.

## Contributing

Want to contribute?
Read [CONTRIBUTING.md](CONTRIBUTING.md)
for the full workflow, repository requirements, and Pull Request process.

Please open an issue to discuss large or breaking changes before implementing.

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

## Author / Contact

Maintained by [Ilya Sitnikov](https://github.com/sitnikovik)
