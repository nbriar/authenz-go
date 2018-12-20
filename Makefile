COMPLEXITY_THRESHOLD ?= 25
GO_PACKAGES ?= $$(go list ./... | grep -v /vendor/ | grep -v /tmp )

export
cover:
	bash ./scripts/coverage.sh

cyclo:
	find . -type f -name '*.go' -not -path './vendor/*' | grep -v health | xargs -I {} gocyclo -over $(COMPLEXITY_THRESHOLD) {}

lint:
	bash ./scripts/lint.sh

tests:
	go test $(GO_PACKAGES)

vet:
	go vet $(GO_PACKAGES)

# This will run go vet and a bunch of other checks at the same time (so we don't need them as
# separate pipeline steps)
check:
	golangci-lint run

guard-%:
	@ if [ "${${*}}" = "" ]; then \
	echo "Environment variable $* not set"; \
	exit 1; \
    fi
