DIST := dist
EXECUTABLE := api
GOFMT ?= gofmt "-s"
GO ?= go

TARGETS ?= linux darwin windows
ARCHS ?= amd64 386
PACKAGES ?= $(shell $(GO) list ./...)
SOURCES ?= $(shell find . -name "*.go" -type f)
TAGS ?=

ifneq ($(DRONE_TAG),)
	VERSION ?= $(DRONE_TAG)
else
	VERSION ?= $(shell git describe --tags --always || git rev-parse --short HEAD)
endif

LDFLAGS ?= -X 'main.Version=$(VERSION)' -X 'main.DroneBuildNumber=$(DRONE_BUILD_NUMBER)' -X 'main.DroneTag=$(DRONE_TAG)'

#ifneq ($(shell uname), Darwin)
#	EXTLDFLAGS = -extldflags "-static" $(null)
##	EXTLDFLAGS = -extldflags $(null)
#else
#	EXTLDFLAGS =
#endif
EXTLDFLAGS = -extldflags "-static" $(null)

all: build

fmt:
	$(GOFMT) -w $(SOURCES)

vet:
	$(GO) vet $(PACKAGES)

lint:
	@hash revive > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/mgechev/revive; \
	fi
	revive -config .revive.toml ./... || exit 1

.PHONY: misspell-check
misspell-check:
	@hash misspell > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/client9/misspell/cmd/misspell; \
	fi
	misspell -error $(SOURCES)

.PHONY: misspell
misspell:
	@hash misspell > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/client9/misspell/cmd/misspell; \
	fi
	misspell -w $(SOURCES)

.PHONY: fmt-check
fmt-check:
	@diff=$$($(GOFMT) -d $(SOURCES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

verify: vet misspell-check fmt-check

test: fmt-check
	@$(GO) test -v -cover -coverprofile coverage.txt ./... && echo "\n==>\033[32m Ok\033[m\n" || exit 1

build:
	$(GO) build -v -tags "$(TAGS)" -ldflags "$(EXTLDFLAGS) -s -w $(LDFLAGS)"  -o api ./apps/api/cmd/api.go

build-cron:
	$(GO) build -v -tags "$(TAGS)" -ldflags "$(EXTLDFLAGS) -s -w $(LDFLAGS)"  -o cron ./apps/cron/cmd/cron.go

release: release-dirs release-build release-copy release-check

release-dirs:
	rm -rf $(DIST); mkdir -p $(DIST)/binaries $(DIST)/release

release-build:
	@which gox > /dev/null; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/mitchellh/gox; \
	fi
	gox -os="$(TARGETS)" -arch="$(ARCHS)" -tags="$(TAGS)" -ldflags="-s -w $(LDFLAGS)" -output="$(DIST)/binaries/$(EXECUTABLE)-$(VERSION)-{{.OS}}-{{.Arch}}"

release-copy:
	$(foreach file,$(wildcard $(DIST)/binaries/$(EXECUTABLE)-*),cp $(file) $(DIST)/release/$(notdir $(file));)

release-check:
	cd $(DIST)/release; $(foreach file,$(wildcard $(DIST)/release/$(EXECUTABLE)-*),sha256sum $(notdir $(file)) > $(notdir $(file)).sha256;)
