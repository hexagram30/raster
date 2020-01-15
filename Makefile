VERSION = $(shell cat VERSION)
GO ?= go
GOFMT ?= $(GO)fmt

DEFAULT_GOPATH=$(shell echo $$GOPATH|tr ':' '\n'|awk '!x[$$0]++'|sed '/^$$/d'|head -1)
ifeq ($(DEFAULT_GOPATH),)
DEFAULT_GOPATH := ~/go
endif
DEFAULT_GOBIN=$(DEFAULT_GOPATH)/bin
export PATH:=$(PATH):$(DEFAULT_GOBIN)

GOLANGCI_LINT=$(DEFAULT_GOBIN)/golangci-lint
RICH_GO = $(DEFAULT_GOBIN)/richgo

GODOC=godoc -index -links=true -notes="BUG|TODO|XXX|ISSUE"

DVCS_HOST = github.com
DVCS_ORG = hexagram30
DOCKER_ORG = $(DVCS_ORG)
PROJ = raster
BIN=rasterd

FQ_PROJ = $(DVCS_HOST)/$(DVCS_ORG)/$(PROJ)

COMMIT_ID = $(shell git rev-parse --short HEAD)

LD_VERSION = -X $(FQ_PROJ)/common.vsn=$(VERSION)
LD_BUILDDATE = -X $(FQ_PROJ)/common.buildDate=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LD_GITCOMMIT = -X $(FQ_PROJ)/common.gitCommit=$(COMMIT_ID)
LD_GITBRANCH = -X $(FQ_PROJ)/common.gitBranch=$(shell git rev-parse --abbrev-ref HEAD)
LD_GITSUMMARY = -X $(FQ_PROJ)/common.gitSummary=$(shell git describe --tags --dirty --always)

LDFLAGS = -w -s $(LD_VERSION) $(LD_BUILDDATE) $(LD_GITBRANCH) $(LD_GITSUMMARY) $(LD_GITCOMMIT)

default: all

all: clean protoc-gen lint test build

#############################################################################
###   Custom Installs   #####################################################
#############################################################################

GOLANGCI_LINT = $(DEFAULT_GOBIN)/golangci-lint
RICH_GO = $(DEFAULT_GOBIN)/richgo
GODA = $(DEFAULT_GOBIN)/goda

$(GOLANGCI_LINT):
	@echo ">> Couldn't find $(GOLANGCI_LINT); installing ..."
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
	sh -s -- -b $(DEFAULT_GOBIN) v1.21.0

$(RICH_GO):
	@echo ">> Couldn't find $(RICH_GO); installing ..."
	@GOPATH=$(DEFAULT_GOPATH) \
	GOBIN=$(DEFAULT_GOBIN) \
	GO111MODULE=on \
	$(GO) get -u github.com/kyoh86/richgo

$(GODA):
	@echo ">> Couldn't find $(GODA); installing ..."
	@GOPATH=$(DEFAULT_GOPATH) \
	GOBIN=$(DEFAULT_GOBIN) \
	GO111MODULE=on \
	$(GO) get -u github.com/loov/goda

#############################################################################
###   Protobuf Support   ####################################################
#############################################################################

# To generate protobuf files, you will need to:
# export GOPATH=~/go:~/lab/hexagram30/go
# export GOBIN=~/go/bin:~/lab/hexagram30/go/bin
# export PATH=$PATH:$GOBIN

PROTOBUF_SRC = api
PROTOBUF_GO = $(PROTOBUF_SRC)
THIRD_PARTY = ./vendor
PROTO_HXGM30_BASE = github.com/hexagram30
PROTO_HXGM30_PROTO = $(PROTO_HXGM30_BASE)/protocols
PROTO_HXGM30_PROTO_REPO = https://$(PROTO_HXGM30_PROTO).git

$(THIRD_PARTY):
	@mkdir -p $(THIRD_PARTY)/$(PROTO_HXGM30_BASE)

$(THIRD_PARTY)/$(PROTO_HXGM30_PROTO):
	@-cd $(THIRD_PARTY)/$(PROTO_HXGM30_BASE) && git clone $(PROTO_HXGM30_PROTO_REPO)

proto-deps: $(THIRD_PARTY) $(THIRD_PARTY)/$(PROTO_HXGM30_PROTO)

protoc-gen: proto-deps protoc-gen-go

protoc-gen-go: clean-protobuf-go deps $(PROTOBUF_GO)/*.pb.go fix-pb-go-import

$(PROTOBUF_GO)/%.pb.go: $(PROTOBUF_SRC)/%.proto
	@protoc -I $(PROTOBUF_SRC) -I $(THIRD_PARTY) --go_out=plugins=grpc:$(PROTOBUF_GO) $<

PROTO_DEP_IMPORT = $(PROTO_HXGM30_PROTO)/src/proto
GOLANG_DEP_IMPORT = $(PROTO_HXGM30_PROTO)/src/golang/common

fix-pb-go-import:
	@sed -i.bak 's|$(PROTO_DEP_IMPORT)|$(GOLANG_DEP_IMPORT)|g' $(PROTOBUF_GO)/*.go && \
	sed -i.bak 's|json:"diceType|json:"dice-type|g' $(PROTOBUF_GO)/*.go && \
	rm $(PROTOBUF_GO)/*.go.bak

clean-protobuf-go:
	@rm -f $(PROTOBUF_GO)/*.pb.go

#############################################################################
###   Build   ###############################################################
#############################################################################

build: build-server

deps:
	@GO111MODULE=off go get github.com/golang/protobuf/protoc-gen-go
	@GO111MODULE=off go install github.com/golang/protobuf/protoc-gen-go

bin:
	@mkdir ./bin

bin/$(BIN): bin
	@echo '>> Building server binary'
	@GO111MODULE=on $(GO) build -ldflags "$(LDFLAGS)" -o bin/$(BIN) ./cmd/$(BIN)

build-server: | bin/$(BIN)

clean:
	@echo ">> Removing compiled binary files ..."
	@rm -f bin/*

clean-all: clean-protobuf-go clean

#############################################################################
###   Docker   ##############################################################
#############################################################################

redix-run:
	@docker run -it \
		-p 7090:7090 -p 6380:6380 \
		-v `pwd`/data/redixdb:/data \
		$(DOCKER_ORG)/redixdb:latest \
		-engine boltdb \
		-storage /data \
		-verbose

#############################################################################
###   Linting & Testing   ###################################################
#############################################################################

show-linter:
	@echo $(GOLANGCI_LINT)

lint-silent: $(GOLANGCI_LINT)
	@$(GOLANGCI_LINT) \
	--enable=typecheck \
	--enable=golint \
	--enable=gocritic \
	--enable=misspell \
	--enable=nakedret \
	--enable=unparam \
	--enable=lll \
	--enable=goconst \
	--out-format=colored-line-number \
	run ./...

lint:
	@echo '>> Linting source code'
	@GO111MODULE=on $(MAKE) lint-silent

test: $(RICH_GO)
	@echo '>> Running all tests'
	@GO111MODULE=on $(RICH_GO) test ./... -v

test-nocolor:
	@echo '>> Running all tests'
	@GO111MODULE=on $(GO) test ./... -v

#############################################################################
###   Release Process   #####################################################
#############################################################################

tag:
	@echo "Tags:"
	@git tag|tail -5
	@git tag "v$(VERSION)"
	@echo "New tag list:"
	@git tag|tail -6

tag-and-push: tag
	@git push --tags

tag-delete: VERSION ?= 0.0.0
tag-delete:
	@git tag --delete v$(VERSION)
	@git push --delete origin v$(VERSION)

#############################################################################
###   Misc   ################################################################
#############################################################################

version:
	@echo $(VERSION)

commit-id:
	@echo $(COMMIT_ID)

show-ldflags:
	@echo $(LDFLAGS)

clean-cache:
	@echo '>> Purging Go mod cahce ...'
	@$(GO) clean -cache
	@$(GO) clean -modcache

rebuild-all: clean-all clean-cache all

check-modules:
	@echo '>> Checking modules ...'
	@GO111MODULE=on $(GO) mod tidy
	@GO111MODULE=on $(GO) mod verify

update-modules:
	@GO111MODULE=on go get -u ./...

list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | \
	awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | \
	sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

.PHONY: default build
