# Makefile for time2dur command

# set BINDIR to the directory in your PATH where you want binary executables (time2dur) installed
BINDIR=/home/jay/.bin/elf
# set BINDIR to the directory in your PATH where you want shell scripts (gotime and getgotime) installed
SCRIPTDIR=/home/jay/.bin

time2dur: time2dur.go
	@go build -o time2dur time2dur.go

vet:
	@go vet time2dur.go

clean:
	@rm -f time2dur

install:
	@cp time2dur $(BINDIR)
	@cp gotime getgotime $(SCRIPTDIR)

wc:
	@wc -l time2dur.go

backup back bak:
	@cp -a gotime getgotime *.go Makefile README.md TODO .bak
