# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/2.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2026-07-13

### Added

- Go module configuration for the project
- Base project template files, including `Makefile` and `golangci-lint` configuration
- `errs` package with shared domain errors for empty commit hashes and failed git execution
- `client/git` package with a Git client to interact with Git repositories
- Implemented `CommitMessage()` to retrieve the commit message for a given commit hash
- Implemented `Commits()` to retrieve commits that come from the provided commit hash
- Implemented `CurrentBranch()` to retrieve the current branch with its name
- Implemented `MergeBase()` to retrieve the merge base commit hash between two branches or commits
- Test fixtures and fake shell execution coverage for the Git client
- Integration tests run with a Docker container via testcontainers
- Added `AGENTS.md` documentation for AI assistants workflow integration

### Fixed

- Adjusted `golangci-lint` configuration and enabled additional linters.
- Fixed `make unit-test`, `make coverage`.
