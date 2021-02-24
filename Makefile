BINARY="app"
#VERSION=1.0.0
BUILD=`date +%FT%T%z`

CONFIG=./config/*.yaml
PACKAGES=`go list ./... | grep -v /vendor/`
VETPACKAGES=`go list ./... | grep -v /vendor/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`
VERSION?=$(shell date +%s)
CI_PROJECT_PATH_SLUG?=$(BINARY)

default:
	docker build . -t dockerhub.piggy.xiaozhu.com/${CI_PROJECT_PATH_SLUG}/${CI_PROJECT_PATH_SLUG}:${VERSION}

list:
	@echo ${PACKAGES}
	@echo ${VETPACKAGES}
	@echo ${GOFILES}

fmt:
	@gofmt -s -w ${GOFILES}

fmt-check:
	@diff=?(gofmt -s -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

install:
	@govendor sync -v

test:
	@go test -cpu=1,2,4 -v ./...

vet:
	@go vet $(VETPACKAGES)

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

run:
	docker build -t ${BINARY} .
	docker run -p $(lcport):80 ${BINARY}

.PHONY: default fmt fmt-check install test vet clean