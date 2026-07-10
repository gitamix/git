//go:build integration
// +build integration

package git_test

const (
	// PathToDockerfile is the path to the Dockerfile used
	// for building the test fixture image.
	PathToDockerfile = "docker/git/Dockerfile"

	// RepoDir is the path to the fixture repository inside the test container.
	RepoDir = "/opt/fixture/repo"

	// ImageRepo is the Docker repository name for the git fixture image.
	ImageRepo = "gitamix/gitamix"

	// ImageTag is the Docker tag for the git fixture image.
	ImageTag = "local"

	// PathToEnv is the path to the fixture hashes file inside the test container.
	//
	// It is set in the Docker Compose file and used to load the fixture environment variables
	// containing commit hashes for the fixture repository.
	PathToEnv = RepoDir + "/.hashes.fixture.env"
)
