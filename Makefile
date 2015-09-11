GITCOMMIT := $(shell git rev-parse --short HEAD 2> /dev/null)
MAIN_PACKAGE=github.com/davidkbainbridge/bp2-template
ALL_PACKAGES=github.com/davidkbainbridge/bp2-template github.com/davidkbainbridge/bp2-template/service
SERVICE=bp2-service
DOCKER_REPO=davidkbainbridge

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
	mkdir -p src/$(dir $(MAIN_PACKAGE))
	(cd src/$(dir $(MAIN_PACKAGE)); ln -s ../../.. $(notdir $(MAIN_PACKAGE)))
	GOPATH=$(abspath $(dir $(lastword $(MAKEFILE_LIST))))	\
	go get -d .

build:
	GOPATH=$(abspath $(dir $(lastword $(MAKEFILE_LIST)))) \
	go build -v -o $(SERVICE) $(MAIN_PACKAGE)

cross-build:
	GOPATH=$(abspath $(dir $(lastword $(MAKEFILE_LIST))))	\
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build -v -o $(SERVICE)-docker $(MAIN_PACKAGE)

clean:
	rm -rf src pkg bin $(SERVICE) $(SERVICE)-docker

enter:
	@echo "Unable to access container shell, please use 'docker exec' command."

image: cross-build
	docker build -t $(DOCKER_REPO)/$(SERVICE):$(GITCOMMIT) .
	docker tag -f $(DOCKER_REPO)/$(SERVICE):$(GITCOMMIT) $(DOCKER_REPO)/$(SERVICE):build

start:
	docker run -tid --name=$(SERVICE) -p 8901:8901 $(DOCKER_REPO)/$(SERVICE):build

logs:
	docker logs bp2-service

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
