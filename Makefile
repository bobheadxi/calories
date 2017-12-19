.PHONY: test deps clean install_glide

PACKAGES = `go list ./... | grep -v vendor/`

all: calories

calories:
	go build

test:
	go test $(PACKAGES) --cover

deps:
	glide install

clean: calories
	rm -f calories

install-glide:
	sh scripts/install_glide.sh