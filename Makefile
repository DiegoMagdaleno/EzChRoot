GOPATH=$(shell pwd)/vendor:$(shell pwd)
GOBIN=$(shell pwd)/bin
GOFILES=$(wildcard *.go)
GONAME=$(shell basename "$(PWD)")
PID=/tmp/go-$(GONAME).pid
DESTDIR=$
PREFIX=/usr/local

build:
	@echo "Building $(GOFILES) to ./bin"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o bin/$(GONAME) $(GOFILES)

get:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get .

install:
	$(MAKE) -C .
	@install -d $(DESTDIR)$(PREFIX)/bin
	@install -m 777 ./bin/EzChRoot $(DESTDIR)$(PREFIX)/bin

run:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOFILES)

watch:
	@$(MAKE) restart &
	@fswatch -o . -e 'bin/.*' | xargs -n1 -I{}  make restart

	restart: clear stop clean build start

start:
	@echo "Starting bin/$(GONAME)"
	@./bin/$(GONAME) & echo $$! > $(PID)

stop:
	@echo "Stopping bin/$(GONAME) if it's running"
	@-kill `[[ -f $(PID) ]] && cat $(PID)` 2>/dev/null || true

clear:
	@clear

clean:
	@echo "Cleaning"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

.PHONY: build get install run watch start stop restart clean