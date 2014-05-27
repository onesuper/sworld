#!/usr/bin/env bash


CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"




export GOPATH="$OLDGOPATH"
echo “installed”