go build: ss_pinger

all: clean dependencies build

linux: clean dependencies build-linux

ss_pinger:
	@echo "Building pinger"
	@echo "Building..."
	$(GOBUILD) ./...
	$(GOBUILD) -o $(BINARY_NAME) -v

clean:
	@echo "Cleanning..."
	-rm -f $(BINARY_NAME)
	-rm -f $(BINARY_UNIX)
	-find . -type d -name -exec rm -rf \{} +
	-$(GOCLEAN) -i
	@echo "Done cleanning."

dependencies:
	@cat /etc/resolv.conf | grep nameserver
	$(GOMOD) tidy
	$(GOMOD) download
	$(GOMOD) vendor
	@echo "Done getting dependencies."

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

GOCMD=go1.18
GOMOD=GO111MODULE=on $(GOCMD) mod
GOGET=GO111MODULE=on $(GOCMD) "get -u"
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=ss_pinger
BINARY_UNIX=$(BINARY_NAME)_unix
