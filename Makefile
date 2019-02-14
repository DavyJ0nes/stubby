.DEFAULT_TARGET=help
.PHONY: all
all: help

# VARIABLES
USERNAME = davyj0nes
APP_NAME = stubby

APP_PORT = 8080
LOCAL_PORT = 8080

VERSION = 0.1.0
COMMIT = $(shell git rev-parse HEAD | cut -c 1-6)
BUILD_TIME = $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

BUILD_PREFIX = CGO_ENABLED=0 GOOS=linux
BUILD_FLAGS = -a -tags netgo --installsuffix netgo
LDFLAGS = -ldflags "-s -w -X ${GO_PROJECT_PATH}/cmd.Release=${VERSION} -X ${GO_PROJECT_PATH}/cmd.Commit=${COMMIT} -X ${GO_PROJECT_PATH}/cmd.BuildTime=${BUILD_TIME}"
GO_BUILD_STATIC = $(BUILD_PREFIX) go build $(BUILD_FLAGS) $(LDFLAGS) main.go

DOCKER_RUN_CMD = docker run -it --rm -p ${LOCAL_PORT}:${APP_PORT} --name ${APP_NAME} ${USERNAME}/${APP_NAME}:${VERSION} "\$$@"

# COMMANDS

## binary: builds a statically linked binary of the application (used in Docker image)
.PHONY: binary
binary:
	$(call blue, "# Building Golang Binary...")
	@go get && ${GO_BUILD_STATIC} -o ${APP_NAME} main.go

## image: builds a docker image for the application
.PHONY: image
image:
	$(call blue, "# Building Docker Image...")
	@docker build --label APP_VERSION=${VERSION} --label BUILT_ON=${BUILD_TIME} --label GIT_HASH=${COMMIT} -t ${USERNAME}/${APP_NAME}:${VERSION} .
	@docker tag ${USERNAME}/${APP_NAME}:${VERSION} ${USERNAME}/${APP_NAME}:latest
	@$(MAKE) clean

## publish: pushes the tagged docker image to docker hub
.PHONY: publish
publish: image
	$(call blue, "# Publishing Docker Image...")
	@docker push docker.io/${USERNAME}/${APP_NAME}:${VERSION}

## run_image: builds and runs the docker image locally
.PHONY: run_image
run_image:
	$(call blue, "# Running Docker Image Locally...")
	@docker run -it --rm --name ${APP_NAME} -v ${PWD}/config.yaml:/config.yaml -p ${LOCAL_PORT}:${APP_PORT} ${USERNAME}/${APP_NAME}:${VERSION}

## clean: remove binary from non release directory
.PHONY: clean
clean: 
	@rm -f ${APP_NAME} 

## help: Show this help message
.PHONY: help
help: Makefile
	@echo "${APP_NAME} - v${VERSION}"
	@echo
	@echo " Choose a command run in "$(APP_NAME)":"
	@echo
	@sed -n 's/^## //p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

# FUNCTIONS
define blue
	@tput setaf 4
	@echo $1
	@tput sgr0
endef
