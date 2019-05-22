# Makefile for building the gotime2dur command and installing
# the gotime, gotime-cli, and gotime-scraper shell scripts.

# set BINDIR to the directory in your PATH where you want binary executables (gotime2dur) installed
BINDIR=/home/jay/.bin/elf

# set BINDIR to the directory in your PATH where you want shell scripts (gotime, gotime-cli, and gotime-scraper) installed
SSDIR=/home/jay/.bin

gotime2dur: gotime2dur.go
	@go build -o gotime2dur gotime2dur.go

vet:
	@go vet gotime2dur.go

clean:
	@rm -f gotime2dur

install:
	@cp gotime2dur $(BINDIR)
	@cp gotime gotime-scraper $(SSDIR)

wc:
	@wc -l gotime2dur.go

backup back bak:
	@cp -a gotime gotime-cli gotime-scraper *.go Makefile README.md TODO .bak
