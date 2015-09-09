GITCOMMIT := $(shell git rev-parse --short HEAD 2> /dev/null)

coverage:
	@echo "Not Yet Implemented"
	@echo "This rule produces the test coverage information for the app."

test:
	@echo "Not Yet Implemented"
	@echo "This rule is called to instantiate the sanity tests for the app. The test includes unit-test, syntax test, and pylint."

prepare: prepare-venv

prepare-venv:
	rm -rf src pkg bin
	mkdir -p src/github.com/davidkbainbridge
	(cd src/github.com/davidkbainbridge; ln -s ../../.. bp2-template)
	GOPATH=$(abspath $(dir $(lastword $(MAKEFILE_LIST))))	\
	go get -d .

build:
	GOPATH=$(abspath $(dir $(lastword $(MAKEFILE_LIST)))) \
	go build -v -o bp2-service github.com/davidkbainbridge/bp2-template

cross-build:
	GOPATH=$(abspath $(dir $(lastword $(MAKEFILE_LIST))))	\
	CGO_ENABLED=0 \
	go build -v -o bp2-service-alpine github.com/davidkbainbridge/bp2-template

clean:
	rm -rf src pkg bin bp2-service bp2-service-alpine

enter:
	@echo "Not Yet Implemented"
	@echo "When the image is running (through start rule), this rule allows the user to get into the container shell."

image: cross-build
	docker build -t dbainbri/bp2-service:$(GITCOMMIT) .

start:
	docker run -tid --name=bp2-service dbainbri/bp2-service:$(GITCOMMIT)

stop:
	docker stop bp2-service
	docker rm bp2-service

dconfigure:
	@echo "Not Yet Implemented"
	@echo "This rule is called by TeamCity to setup the docker image required for the tests."

dutest:
	@echo "Not Yet Implemented"
	@echo "This rule is called by TeamCity to execute the unit-test."

ditest:
	@echo "Not Yet Implemented"
	@echo "This rule is called by TeamCity to execute the integration-test."
