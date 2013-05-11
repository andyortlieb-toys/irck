#!/usr/bin/env bash
source $(cd $(dirname $0); pwd )/activate.sh

mkdir -p $GOPATH
echo $GOPATH

go get github.com/thoj/go-ircevent
