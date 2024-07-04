#!/usr/bin/bash

rm -rf ./bin
rm -rf ./log
rm -rf ./output
mkdir log
mkdir output
chmod 777 output

go mod tidy
cd main
go build -o ../bin/mini_spider