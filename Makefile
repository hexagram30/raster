VERSION = 1.0.0
DOCKER_ORG = hexagram30
REDIX_VERSION = 1.10
DOCKER_DIR = build/docker
REDIX_DIR = $(DOCKER_DIR)/redixdb

DEFAULT_GOPATH=$(shell echo $$GOPATH|tr ':' '\n'|awk '!x[$$0]++'|sed '/^$$/d'|head -1)
ifeq ($(DEFAULT_GOPATH),)
DEFAULT_GOPATH := ~/go
endif
DEFAULT_GOBIN=$(DEFAULT_GOPATH)/bin
export PATH:=$(PATH):$(DEFAULT_GOBIN)

default: build

build: protoc-gen

deps:
	@GO111MODULE=off go install github.com/golang/protobuf/protoc-gen-go

# To regen, you will need to:
# export GOPATH=~/go:~/lab/hexagram30/go
# export GOBIN=~/go/bin:~/lab/hexagram30/go/bin
# export PATH=$PATH:$GOBIN
protoc-gen: api/*.pb.go

api/%.pb.go: api/%.proto 
	@protoc -I api --go_out=plugins=grpc:api $<

redix-image:
	@docker build -t $(DOCKER_ORG)/redixdb $(REDIX_DIR)
	@docker tag $(DOCKER_ORG)/redixdb $(DOCKER_ORG)/redixdb:latest
	@docker tag $(DOCKER_ORG)/redixdb $(DOCKER_ORG)/redixdb:$(REDIX_VERSION)

redix-publish:
	@docker push $(DOCKER_ORG)/redixdb:latest
	@docker push $(DOCKER_ORG)/redixdb:$(REDIX_VERSION)

redix-run:
	@docker run -it \
		-p 7090:7090 -p 6380:6380 \
		-v `pwd`/data/redixdb:/data \
		$(DOCKER_ORG)/redixdb:latest \
		-engine boltdb \
		-storage /data \
		-verbose
