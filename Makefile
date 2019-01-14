SHELL := /bin/bash
TMP:=$(shell /usr/bin/mktemp -d)

build-server:
	go build -o bin/predictor cmd/server/main.go

server: build-server
	./bin/predictor -grpc-port 50051 -http-port 8080


build-client:
	go build -o bin/crystal-ball cmd/client-grpc/main.go

client: build-client
	./bin/crystal-ball --grpc-addr localhost:50051 card -b 1986-04-16


proto:
	@protoc -I/usr/local/include -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:$(TMP) \
		api/proto/v1/predictor.proto
	@protoc -I/usr/local/include -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:$(TMP) \
		api/proto/v1/predictor.proto
	@protoc -I/usr/local/include -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:$(TMP) \
		api/proto/v1/predictor.proto
	@install -v -m 0644 $(TMP)/api/proto/v1/predictor.pb.go pkg/api/v1/
	@install -v -m 0644 $(TMP)/api/proto/v1/predictor.pb.gw.go pkg/api/v1/
	@install -v -m 0644 $(TMP)/api/proto/v1/predictor.swagger.json api/swagger/v1/
	@rm -rf $(TMP)

cpu-prof:
	go build -o predictor-main main.go
	go tool pprof --pdf ~/go/bin/yourbinary /var/path/to/cpu.pprof > file.pdf

mem-prof:
	PPROF_TMPDIR=$(TMP) go build -o $(TMP)/predictor-main main.go
	PPROF_TMPDIR=$(TMP) $(TMP)/predictor-main 1>/dev/null
	go tool pprof -top $(TMP)/predictor-main $(TMP)/mem.pprof
	@rm -rf $(TMP)

# .PHONY: all

# all: build

# # Package and name are determined from the directory layout
# PACKAGE       := $(shell go list 2>/dev/null)
# NAME          := $(shell basename ${PACKAGE})
# # Version is implied through signed git tags.
# # If the current branch is not exactly a signed tag, the current branch is used instead.
# VERSION       := $(shell git describe --exact-match 2>/dev/null || git symbolic-ref -q --short HEAD)
# BUILD_DIR      = ${GOPATH}/src/${PACKAGE}
# CURRENT_DIR    = $(shell pwd)
# BUILD_DIR_LINK = $(shell readlink ${BUILD_DIR})

# GO_SOURCES      := $(shell go list -f '{{join .GoFiles " "}}')
# GO_TEST_SOURCES := $(shell go list -f '{{join .TestGoFiles " "}}')
# GOHOSTOS        := $(shell go env GOHOSTOS 2>/dev/null)
# GOHOSTARCH      := $(shell go env GOHOSTARCH 2>/dev/null)

# BUILD_OPTS := -ldflags '-w -X main.Version=$(VERSION)' -a

# ifeq ($(BUILD_QUIET),1)
#     Q = @
# else
#     Q =
# endif

# .PHONY: build
# build: lint vet test $(NAME)

# $(NAME): $(GO_SOURCES) vendor rice
#   $(Q)go build $(BUILD_OPTS) -o $(NAME) $(PACKAGE)

# .PHONY: install
# install: $(GO_SOURCES) vendor rice
#   $(Q)go install $(BUILD_OPTS) $(PACKAGE)

# .PHONY: link
# link:
#   @BUILD_DIR=${BUILD_DIR}; \
#   BUILD_DIR_LINK=${BUILD_DIR_LINK}; \
#   CURRENT_DIR=${CURRENT_DIR}; \
#   if [ "$${BUILD_DIR_LINK}" != "$${CURRENT_DIR}" ]; then \
#     echo "Fixing symlinks for build"; \
#     rm -f $${BUILD_DIR}; \
#     mkdir -p $(shell dirname ${BUILD_DIR}); \
#     ln -s $${CURRENT_DIR} $${BUILD_DIR}; \
#   fi

# GODEP  := ${GOPATH}/bin/dep
# GOLINT := ${GOPATH}/bin/golint
# GOVET  := $(shell go env GOTOOLDIR)/vet
# GORICE := ${GOPATH}/bin/rice

# $(GODEP)  : ; $(Q)go get github.com/golang/dep/cmd/dep
# $(GOLINT) : ; $(Q)go get github.com/golang/lint/golint
# $(GOVET)  : ; $(Q)go get golang.org/x/tools/cmd/vet
# $(GORICE) : ; $(Q)go get github.com/GeertJohan/go.rice/rice

# # Just make this every time rather than tracking the dependencies....
# .PHONY: rice
# rice: $(GORICE)
# #	$(Q)cd cmd && rice embed-go

# vendor: $(GODEP) Gopkg.toml Gopkg.lock
#   $(Q)cd $(BUILD_DIR) && $(GODEP) ensure
#   $(Q)touch vendor

# .PHONY: fmt
# fmt:
#   $(Q)cd $(BUILD_DIR) && go fmt ./...

# .PHONY: lint
# lint: $(GOLINT)
#   $(Q)cd $(BUILD_DIR) && $(GOLINT) $$(cd $(BUILD_DIR) && go list ./...)

# .PHONY: vet
# vet: vendor $(GOVET)
#   $(Q)cd $(BUILD_DIR) && \
#   for src in $(GO_SOURCES) $(GO_TEST_SOURCES); do \
#     $(GOVET) -all=true $$src; \
#   done

# .PHONY: test
# test: vendor
#   $(Q)cd $(BUILD_DIR) && go test -v -timeout 30s -parallel=4 ./...

# .PHONY: testacc
# testacc: vendor
#   $(Q)cd $(BUILD_DIR) && TF_ACC=1 go test -v -timeout 10m ./...

# .PHONY: clean
# clean:
#   rm -f $(NAME)*

# .PHONY: clean-all
# clean-all: clean
#   rm -fr vendor/
#   rm -fr ${GOPATH}/bin/$(NAME)

# .PHONY: release-artifacts
# release-artifacts: $(NAME)_darwin_amd64 $(NAME)_darwin_amd64.asc $(NAME)_linux_amd64 $(NAME)_linux_amd64.asc

# $(NAME)_linux_amd64: vendor test $(GO_SOURCES) rice
#   $(Q)GOOS=linux GOARCH=amd64 go build $(BUILD_OPTS) -o $(NAME)_linux_amd64 $(PACKAGE)

# $(NAME)_darwin_amd64: vendor test $(GO_SOURCES) rice
#   $(Q)GOOS=darwin GOARCH=amd64 go build $(BUILD_OPTS) -o $(NAME)_darwin_amd64 $(PACKAGE)

# %_amd64.asc: %_amd64
#   $(Q)gpg --armor --output $@ --detach-sig $<

