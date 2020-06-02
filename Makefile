

all: build

build: 
	@go build

clean:
	rm -f gtodo

help:
	@echo "make - build the source code"
	@echo "make clean - clear the binary file"

.PHONY: clean help
