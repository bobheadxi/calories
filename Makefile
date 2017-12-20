.PHONY: test deps clean glide db

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

glide:
	sh scripts/install_glide.sh

db:
	sh scripts/create_db.sh