
# The path from makefile to project
PROJECT_PATH=.

# The folder where the builds are stored - usually bin or build
PROJECT_BUILD_FOLDER=build

# This how we want to name the binary output
BINARY=gutils

# These are the values we want to pass for VERSION and BUILD
VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`

# The build flags
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"

BUILD_FULL_PATH=${PROJECT_BUILD_FOLDER}
ENTRY_FILE_PATH=./main.go

build:
	rm -rf ${PROJECT_PATH}/${PROJECT_BUILD_FOLDER}
	cd ${PROJECT_PATH} && go mod tidy
	cd ${PROJECT_PATH} && go mod verify
	cd ${PROJECT_PATH} && env go build ${LDFLAGS} -o ${BUILD_FULL_PATH}/${BINARY} ${ENTRY_FILE_PATH}
	cd ${PROJECT_PATH} && env GOOS=darwin go build ${LDFLAGS} -o ${BUILD_FULL_PATH}/darwin/${BINARY} ${ENTRY_FILE_PATH}
	cd ${PROJECT_PATH} && env GOOS=linux go build ${LDFLAGS} -o ${BUILD_FULL_PATH}/linux/${BINARY} ${ENTRY_FILE_PATH}
	cd ${PROJECT_PATH} && env GOOS=windows GOARCH=386 go build ${LDFLAGS} -o ${BUILD_FULL_PATH}/windows/${BINARY}.exe ${ENTRY_FILE_PATH}
	@echo "\n\nBinary built at ${PROJECT_PATH}/${BUILD_FULL_PATH}"

build-dev:
	rm -rf ${PROJECT_PATH}/${BUILD_FULL_PATH}/${BINARY}
	cd ${PROJECT_PATH} && go mod tidy
	cd ${PROJECT_PATH} && go mod verify
	cd ${PROJECT_PATH} && env go build ${LDFLAGS} -o ${BUILD_FULL_PATH}/${BINARY} ${ENTRY_FILE_PATH}

run:
	cd ${PROJECT_PATH} && ${BUILD_FULL_PATH}/${BINARY} $(filter-out $@,$(MAKECMDGOALS)) ${ARGS}

run-dev:
	cd ${PROJECT_PATH} && APP_ENV=development ${BUILD_FULL_PATH}/${BINARY} $(filter-out $@,$(MAKECMDGOALS)) ${ARGS}

clean:
	rm -rf ${PROJECT_PATH}/${PROJECT_BUILD_FOLDER} ${PROJECT_PATH}/vendor
	cd ${PROJECT_PATH} && go mod vendor

test:
	cd ${PROJECT_PATH}/pkg/calculate/scale/length && go test

%:
	@:

.PHONY: build clean
