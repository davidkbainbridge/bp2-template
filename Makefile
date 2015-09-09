GITCOMMIT := $(shell git rev-parse --short HEAD 2> /dev/null)
PACKAGES=github.com/davidkbainbridge/bp2-template
ALL_PACKAGES=github.com/davidkbainbridge/bp2-template github.com/davidkbainbridge/bp2-template/service
SERVICE=bp2-service
DOCKER_FOLDER=davidkbainbridge

coverage:
	@echo "Not Yet Implemented"
	@echo "This rule produces the test coverage information for the app."

test:
	GOPATH=$(abspath $(dir $(lastword $(MAKEFILE_LIST)))) \
	go vet $(ALL_PACKAGES)
	GOPATH=$(abspath $(dir $(lastword $(MAKEFILE_LIST)))) \
	go test -v -cover $(ALL_PACKAGES)

prepare: prepare-venv

prepare-venv:
	rm -rf src pkg bin
	mkdir -p src/github.com/davidkbainbridge
	(cd src/github.com/davidkbainbridge; ln -s ../../.. bp2-template)
	GOPATH=$(abspath $(dir $(lastword $(MAKEFILE_LIST))))	\
	go get -d .

build:
	GOPATH=$(abspath $(dir $(lastword $(MAKEFILE_LIST)))) \
	go build -v -o $(SERVICE) $(PACKAGES)

cross-build:
	GOPATH=$(abspath $(dir $(lastword $(MAKEFILE_LIST))))	\
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build -v -o $(SERVICE)-alpine $(PACKAGES)

clean:
	rm -rf src pkg bin $(SERVICE) $(SERVICE)-alpine

enter:
	@echo "Unable to access container shell, please use 'docker exec' command."

image: cross-build
	docker build -t $(DOCKER_FOLDER)/$(SERVICE):$(GITCOMMIT) .

start:
	docker run -tid --name=bp2-service $(DOCKER_FOLDER)/$(SERVICE):$(GITCOMMIT)

stop:
	docker stop $(SERVICE)
	docker rm $(SERVICE)

dconfigure:
	@echo "Not Yet Implemented"
	@echo "This rule is called by TeamCity to setup the docker image required for the tests."

dutest:
	@echo "Not Yet Implemented"
	@echo "This rule is called by TeamCity to execute the unit-test."

ditest:
	@echo "Not Yet Implemented"
	@echo "This rule is called by TeamCity to execute the integration-test."
