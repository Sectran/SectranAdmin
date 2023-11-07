SHELL = /bin/bash
CMD_GIT ?= git
CMD_GO ?= go
CMD_GREP ?= grep
CMD_AWK ?= awk
CMD_CUT ?= cut
CMD_SED ?= sed
MAKE = make

#
# version
#
OS ?=$(shell go env GOOS)
ARCH ?=$(shell go env GOARCH)
COMMIT := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
DATE := $(shell git log -1 --format=%cd --date=format:"%Y%m%d")

VERSIONGO ?= cli/version/version.go

.check_%:
	@command -v $* >/dev/null
	if [ $$? -ne 0 ]; then
		echo "Sectran Makefile: missing required tool $*"
		exit 1
	else
		touch $@ # avoid target rebuilds due to inexistent file
	fi

GO_VERSION = $(shell $(CMD_GO) version 2>/dev/null | $(CMD_AWK) '{print $$3}' | $(CMD_SED) 's:go::g' | $(CMD_CUT) -d. -f1,2)
GO_VERSION_MAJ = $(shell echo $(GO_VERSION) | $(CMD_CUT) -d'.' -f1)
GO_VERSION_MIN = $(shell echo $(GO_VERSION) | $(CMD_CUT) -d'.' -f2)

.PHONY: .checkver_$(CMD_GO)
.checkver_$(CMD_GO):
	@if [ "$(GO_VERSION_MAJ)" -eq 1 ] && [ "$(GO_VERSION_MIN)" -lt 18 ]; then \
		echo "You MUST use golang 1.18 or newer, your current golang version is ${GO_VERSION}"; \
		exit 1; \
	fi

.PHONY: help
help:
	@echo "# environment"
	@echo "    $$ make env					# show makefile environment/variables"
	@echo ""
	@echo "# build"
	@echo "    $$ make build					# build Sectran in release mode"
	@echo "    $$ make debug					# build Sectran in debug mode"
	@echo "    $$ make release					# build Sectran int release mode"
	@echo ""
	@echo "# clean"
	@echo "    $$ make clean				# wipe ./bin/"
	@echo ""

.PHONY: env
env:
	@echo ---------------------------------------
	@echo "Sectran Makefile Environment:"
	@echo ---------------------------------------
	@echo "GO_VERSION              $(GO_VERSION)"
	@echo "OS                      $(OS)"
	@echo "ARCH                    $(ARCH)"
	@echo "GO_VERSION              $(GO_VERSION)"
	@echo "GO_VERSION_MAJ          $(GO_VERSION_MAJ)"
	@echo "GO_VERSION_MIN          $(GO_VERSION_MIN)"
	@echo ---------------------------------------

.PHONY: version
version:
	@echo "Building with commit: $(COMMIT), branch: $(BRANCH), time: $(DATE)"
	@echo "package version" > ${VERSIONGO}
	@echo "" >> ${VERSIONGO}
	@echo "var (" >> ${VERSIONGO}
	@echo "    Commit string = \"$(COMMIT)\"" >> ${VERSIONGO}
	@echo "    Branch string = \"$(BRANCH)\"" >> ${VERSIONGO}
	@echo "    BuildTime string = \"$(DATE)\"" >> ${VERSIONGO}
	@echo ")" >> ${VERSIONGO}

DCMAKE_BUILD_TYPE = Release
.PHONY: build
build: version .checkver_$(CMD_GO)
	@mkdir -p pkg
	@if [ -d ./backend/terminal/build ]; then rm -rf ./backend/terminal/build; fi
	@mkdir -p ./backend/terminal/build
	cd ./backend/terminal/build && cmake .. -DCMAKE_BUILD_TYPE=$(DCMAKE_BUILD_TYPE) && make && make install && cd -
	CGO_ENABLED=1 $(CMD_GO) build -ldflags "-w -s -extldflags=-Wl,-rpath,." -o pkg/sectran-${OS}-${ARCH}

debug: DCMAKE_BUILD_TYPE = Debug
debug: build

release: build

.PHONY: package
package: build
	@if [ -f pkg.tar.gz ]; then rm -f pkg.tar.gz; fi
	tar -zcvf pkg.tar.gz pkg
.PHONY: clean
clean:
	@if [ -e bin/sectran-${os}-${arch} ]; then rm -f bin/sectran-${os}-${arch}; fi
	@if [ -e cli/version/version.go ]; then rm -f cli/version/version.go; fi
	@if [ -e terminal.dump ]; then rm -f terminal.dump; fi
	@if [ -d build ]; then rm -f build; fi
	

# goctl api go --api ./apis/apis.api --dir ./apiservice --style goZero