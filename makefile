#!/usr/bin/env make -f

SHELL=`which bash`
bin=bin
name=OpenPushOver

all: build

build:
	go build -v -o $(bin)/$(name)

clean:
	go clean -x

remove:
	go clean -i
