## Overview

This is a countdown timer that opens the Go Time web page when the show is scheduled to start.

It uses the Go Time API  https://changelog.com/slack/countdown/gotime to get the time the next show starts.

Currently, it is just a prototype that is implemented with some shell scripts and two Go programs.
This may all be rewritten in Go, using Fyne for the GUI, for proper cross-platform support.

This README was written quickly and is very rough. Don't expect too much from it.

Dependencies:

1. The **timer** program from my Chronograph repository: https://github.com/Yaoir/Chronograph
2. **bash(1)**, **curl(1)**, **sed(1)**, **grep(1)**, **echo(1)**, web browser (Firefox assumed)
2. GNU **make**, if you want to use **make** for compiling and/or installing.
3. The **gotime** script relies on KDE and **konsole**. If you want to use Gnome or some other desktop and/or virtual terminal, you can modify the script.
4. The optional **gotime-autostart** script is for adding an Autostart program to KDE. I don't know if it works with any other desktop.

This works on Linux at least. Maybe macOS? But on Windows, I assume you will need something like Cygwin or Windows Subsystem for Linux. I haven't tried it.

## Introduction

Here is a quick run down of the programs:

check

A shell script just for testing. It is for checking to make sure the API is running and is reporting a time for the next scheduled show. It will print either an RFC3339 time string or "No show scheduled" if it can't find a time string at the API URL.

gotime-fetch.go

A Go program that grabs the RFC3339 time string from the Go Time API and prints it. If no show is currently scheduled, it reports that by printing the string provided by the API, and exits with an exit code of 1.

gotime2dur.go

A Go program that converts a Go Time value (in this case, the RFC3339 string) to a Go Duration value. It's needed because the **timer** program requires a Duration. (I could have used **alarm** from the Chronograph repository to accept a Time value directly, but I wanted to display a count down to the show start, not an alarm clock.)

gotime-cli

A shell script that runs the countdown timer in a terminal session, on the command line. To use it, you will probably want to open a new virtual terminal and run the command in that. When the timer reaches zero, a web browser (Firefox by default) is started to open the Go Time web page. If you want to use Chrome or some other browser, you can modify the script.

gotime

This shell script opens a new virtual terminal (**konsole** by default), and runs gotime-cli in it. For it to run nicely, you will need to create a profile for konsole with an appropriate height and width, font size, etc. See the comments in the script for help with that.

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

Try the **check** program and see if it prints an RFC3339 string or "No show scheduled". If that works, it's accessing the API successfully. If that works, run **gotime-fetch** to make sure that works too. It's the script used by **gotime**.

Next, try **gotime-cli** to see if it works on your command line shell. You can then configure the **gotime** script to open a virtual terminal window with the timer running in it.

To install:

Before installing, edit Makefile and each shell script. Set the BINDIR and ELFDIR varibles the way you want them. Both need to be in your PATH environment variable. BINDIR is where you keep shell scripts, and ELFDIR is where you keep binary executable programs. It is ok if they are set to the same directory. Then run

```
$ make install
```

To install the **gotime-autostart** script, open KDE's system settings -> Startup and Shutdown -> Autostart. Click on the Add Script button and enter the full path spec of the gotime-autostart script. (You cannot count on your PATH variable to be set yet, so a full path spec is necessary.) Set the "Run On" field to "Startup".

Note: this script currently assumes the show is on Tuesdays, and runs the **gotime** script when KDE is started at any time on a Tuesday.
