
FLUX_SCHED_ROOT ?= /opt/flux-sched
INSTALL_PREFIX ?= /usr
LIB_PREFIX ?= /usr/lib
LOCALBIN ?= $(shell pwd)/bin
COMMONENVVAR=GOOS=$(shell uname -s | tr A-Z a-z)
BUILDENVVAR=CGO_CFLAGS="-I${FLUX_SCHED_ROOT} -I${FLUX_SCHED_ROOT}/resource/reapi/bindings/c" CGO_LDFLAGS="-L${LIB_PREFIX} -L${LIB_PREFIX}/flux -L${FLUX_SCHED_ROOT}/resource/reapi/bindings -lreapi_cli -lflux-idset -lstdc++ -lczmq -ljansson -lhwloc -lboost_system -lflux-hostlist -lboost_graph -lyaml-cpp"

REGISTRY=ghcr.io/converged-computing
IMAGE=fluxion:latest
RELEASE_VERSION?=v$(shell date +%Y%m%d)-$(shell git describe --tags --match "v*")

.PHONY: all
all: build

.PHONY: $(LOCALBIN)
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

.PHONY: build
build: 
	docker build --build-arg ARCH="amd64" --build-arg RELEASE_VERSION="$(RELEASE_VERSION)" -t $(REGISTRY)/$(IMAGE) .

.PHONY: server
server: 
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o bin/server cmd/main.go

.PHONY: protoc
protoc: $(LOCALBIN)
	GOBIN=$(LOCALBIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	GOBIN=$(LOCALBIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	
# You can use make protoc to download proto
.PHONY: proto
proto: protoc
	PATH=$(LOCALBIN):${PATH} protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/fluxion-grpc/fluxion.proto
	PATH=$(LOCALBIN):${PATH} protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/service-grpc/service.proto
