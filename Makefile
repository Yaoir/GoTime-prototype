# Makefile for building the gotime2dur command and installing the shell scripts.

# CONFIGURATION:

# set ELFDIR to the directory in your PATH where you want binary executables (gotime2dur) installed
ELFDIR=/home/jay/.bin/elf

# set BINDIR to the directory in your PATH where you want shell scripts installed
BINDIR=/home/jay/.bin

# END OF CONFIGURATION

SCRIPTS=gotime-autostart gotime gotime-cli
COMMANDS=gotime-fetch gotime2dur

gotime-fetch: gotime-fetch.go
	@go build -o gotime-fetch gotime-fetch.go

gotime2dur: gotime2dur.go
	@go build -o gotime2dur gotime2dur.go

vet:
	@go vet gotime2dur.go

clean:
	@rm -f gotime2dur

install:
	@cp $(COMMANDS) $(ELFDIR)
	@cp $(SCRIPTS)  $(BINDIR)

wc:
	@wc -l *.go

backup back bak:
	@cp -a $(SCRIPTS) *.go Makefile README.md TODO .bak
