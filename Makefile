.PHONY: test test-verbose deps clean glide db deploy

PACKAGES = `go list ./... | grep -v vendor/`

all: calories

calories:
	go build

test:
	make db
	go test $(PACKAGES) --cover

test-verbose:
	make db
	go test $(PACKAGES) -v --cover

clean: calories
	rm -f calories

db:
	sh scripts/create_db.sh

deploy: calories
	sh scripts/heroku_deploy.sh
