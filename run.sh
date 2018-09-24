#!/bin/sh

export GOPATH=$GOPATH/ShortURL/vender:$GOPATH

mkdir -p output/bin 

go build -a -o output/bin/work
./output/bin/work
