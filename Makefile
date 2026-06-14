.PHONY: build install clean vet test all

BINARY=archseed
INSTALL_DIR=$(HOME)/.bin

build:
	go build -o $(BINARY) .

install: build
	cp $(BINARY) $(INSTALL_DIR)/$(BINARY)
	@echo "Installed $(BINARY) to $(INSTALL_DIR)/$(BINARY)"

clean:
	rm -f $(BINARY)

vet:
	go vet ./...

test:
	go test ./...

all: vet test build
