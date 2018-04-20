PROJECT = golang-api-example
GOARCH = amd64

VERSION=0.0.1
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

PROJECT_DIR=${GOPATH}/src/github.com/madsilver/${PROJECT}
CMD_DIR=${PROJECT_DIR}/cmd
BUILD_DIR=${PROJECT_DIR}/build
VET_REPORT = ${BUILD_DIR}/vet.report
TEST_REPORT = ${BUILD_DIR}/tests.xml

LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

.DEFAULT_GOAL := linux

# Build the project
all: depend test vet linux darwin windows

depend:
	go get github.com/tebeka/go2xunit

linux:
	cd ${CMD_DIR}; \
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BUILD_DIR}/${PROJECT}-linux-${GOARCH} . ; \
	cd - >/dev/null

darwin:
	cd ${CMD_DIR}; \
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BUILD_DIR}/${PROJECT}-darwin-${GOARCH} . ; \
	cd - >/dev/null

windows:
	cd ${CMD_DIR}; \
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BUILD_DIR}/${PROJECT}-windows-${GOARCH}.exe . ; \
	cd - >/dev/null

test:
	cd ${PROJECT_DIR}; \
	go test -v ./... 2>&1 | go2xunit -output ${TEST_REPORT} ; \
	cd - >/dev/null

vet:
	-cd ${PROJECT_DIR}; \
	go vet ./... > ${VET_REPORT} 2>&1 ; \
	cd - >/dev/null

fmt:
	cd ${PROJECT_DIR}; \
	go fmt $$(go list ./... | grep -v /vendor/) ; \
	cd - >/dev/null

clean:
	@rm -fR ./build/ ./vendor/

.PHONY: linux darwin windows test vet fmt clean