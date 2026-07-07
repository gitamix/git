# Run all tests in the project.
.PHONY: test
test: unit-test integration-test

# Run unit tests.
.PHONY: unit-test
unit-test:
	@echo 🧪 Running unit tests...
	@mkdir -p tmp
	@pkgs=$$(go list ./... | grep -v 'internal/test/'); \
	if [ -n "$$pkgs" ]; then \
		go test -race -count=1 \
		-covermode=atomic \
		-coverprofile=tmp/coverage_unit.out \
		$$pkgs; \
	else \
		echo "no packages to test"; \
	fi

# Run integration tests.
.PHONY: integration-test
integration-test:
	@echo 🧪 Running integration tests...
	@mkdir -p tmp
	@coverpkgs=$$(go list ./... | grep -v 'internal/test/integration' | grep -v 'internal/test/fake' | tr '\n' ',' | sed 's/,$$//'); \
		go test -tags=integration \
		-race -count=1 -v \
		-covermode=atomic \
		-coverprofile=tmp/coverage_integration.out \
		-coverpkg=$$coverpkgs \
		./internal/test/integration/...; \
		if [ -f tmp/coverage_integration.out ] && [ $$(wc -l < tmp/coverage_integration.out) -le 1 ]; then \
			rm -f tmp/coverage_integration.out; \
		fi
