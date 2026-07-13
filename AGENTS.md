# AGENTS.md

## Project Overview

- **Language**: Go 1.25.7
- **Module**: `github.com/gitamix/git`
- **Structure**: `internal/`, `client/`, `docs/`, `scripts/`, `errs/`
- **Testing**: testify + testcontainers-go
- **Documentation**: GoDoc (pkg.go.dev)

## Architectural Principles

This project follows a strict declarative programming style with these key rules:

- **Immutability**: Objects cannot be changed after creation. Create new instances instead of mutating.
- **No setters**: Avoid setter methods. Values should be fixed at construction time.
- **No getters**: Use direct method names (e.g., `Price()` instead of `GetPrice()`).
- **Functional options**: All constructor parameters use the functional options pattern.
- **One constructor per type**: Single constructor `New<TypeName>` with variadic options.
- **No logic in constructors**: Constructors only assign values, no validation or side effects.
- **Unexported fields**: All struct fields must be unexported for encapsulation.
- **Exported functions/types**: All public APIs must be exported.

**See [STYLE.MD](STYLE.MD) for complete style guide.**

## Project Structure

| Directory | Purpose |
|-----------|---------|
| `client/` | Public client API |
| `internal/` | Internal implementation |
| `internal/test/` | Test infrastructure (fixtures, fakes) |
| `internal/test/integration/` | Integration tests |
| `docs/` | Documentation |
| `scripts/` | Utility scripts |
| `errs/` | Shared domain errors |

## Development Workflow

### Setup

```bash
# Install Git hooks and personalize template placeholders
make setup
```

### Testing

```bash
# Run all tests (unit + integration)
make test

# Run unit tests only
make unit-test

# Run integration tests only
make integration-test
```

### Code Quality

```bash
# Run all checks (test + coverage + lint)
make check

# Run linters only
make lint

# Generate coverage report
make coverage
```

### Commit Process

1. Follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
2. Use `feat:`, `fix:`, `docs:`, `chore:`, etc. prefixes
3. Update `CHANGELOG.md` for user-visible changes
4. Run `make check` before committing/pushing

### Release Process

1. Update `CHANGELOG.md` with the new version
2. Run the `Release` GitHub Actions workflow
3. Provide the version number (without `v` prefix)
4. The workflow creates a tag and publishes the release

## Linting Configuration

The project uses [golangci-lint](https://golangci-lint.run/) with the following linters enabled:

- `govet`, `staticcheck`, `unused`, `errcheck`, `revive`
- `misspell`, `unparam`, `gocyclo`, `paralleltest`
- `godoclint`, `godot`, `godox`, `dogsled`, `bodyclose`
- `testpackage`

**See [.golangci.yml](.golangci.yml) for full configuration.**

## Git Hooks

- **pre-commit**: Runs `make unit-test`
- **pre-push**: Runs `make check`

## Documentation Standards

- All public and private functions, types, and fields must have GoDoc comments
- First sentence should be a short summary starting with the function name
- Use full sentences and describe what, not how
- Document parameters and return values clearly

## Testing Standards

- Tests go in separate `_test` packages
- Use [testify](https://github.com/stretchr/testify) for assertions
- Follow the [canonical test template](STYLE.MD#canonical-test-template)
- Use `t.Parallel()` where appropriate
- Test fixtures in `internal/test/fixture/`
- Fake implementations in `internal/test/fake/`

## Important Commands Reference

| Command | Description |
|---------|-------------|
| `make setup` | Install hooks and personalize template |
| `make test` | Run all tests |
| `make unit-test` | Run unit tests |
| `make integration-test` | Run integration tests |
| `make lint` | Run linters |
| `make coverage` | Generate coverage report |
| `make check` | Run all checks (test + coverage + lint) |

## contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for the full contribution workflow.

## License

MIT License - see [LICENSE](LICENSE) for details.
