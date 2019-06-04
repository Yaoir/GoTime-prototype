## Overview

This is a countdown timer that opens the Go Time web page when the show is scheduled to start.

It uses the Go Time API  https://changelog.com/slack/countdown/gotime

Currently, it is just a prototype that is implemented with some shell scripts and two Go programs.

Dependencies:

1. The **timer** program from my Chronograph repository: https://github.com/Yaoir/Chronograph
2. GNU **make**, if you want to use **make** for compiling and/or installing.
3. The **gotime** script relies on KDE and **konsole**. If you want to use Gnome or some other desktop and/or virtual terminal, you can modify the script.

It works on Linux at least. Maybe macOS?  But on Windows, I assume you will need something like Cygwin or Windows Subsystem for Linux. I haven't tried it.

This may all be rewritten in Go, using Fyne for the GUI, for proper cross-platform support.

## Introduction

## Quick Start

## Compiling and Installing

**gotime2dur** is written in Go. To compile it, you need to have Go installed.

To compile:

```
$ go build gotime2dur.go
```
or using GNU **make**:
```
$ make
```

To test:

Try the **check** program and see if it prints an RFC3339 string or "No show scheduled". If that works, it's accessing the API successfully. If that works, run **gotime-scraper** to make sure that works too. It's the script used by **gotime**.

Next, try **gotime-cli** to see if it works on your command line shell. You can then configure the **gotime** script to open a virtual terminal window with the timer running in it.

To install:

Before installing, edit Makefile and each shell script. Set the BINDIR and ELFDIR varibles the way you want them. Both need to be in your PATH environment variable. BINDIR is where you keep shell scripts, and ELFDIR is where you keep binary executable programs. It is ok if they are set to the same directory. Then run

```
$ make install
```
