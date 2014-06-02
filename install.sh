#!/usr/bin/env bash

VERSION='0.0.1'
CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"


go install sworld
go install test


export GOPATH="$OLDGOPATH"
echo “installed”
echo $VERSION