# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/2.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Go module configuration for the project.
- Base project template files, including `Makefile` and `golangci-lint` configuration.
- `errs` package with shared domain errors for empty commit hashes and failed git execution.
- `client/git` package with a Git client and `CommitMessage` retrieval by commit hash.
- Test fixtures and fake shell execution coverage for the Git client.
- Integration tests run with a Docker container via testcontainers

### Fixed

- Adjusted `golangci-lint` configuration and enabled additional linters.
- Fixed `make unit-test`, `make coverage`.
