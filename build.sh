#!/bin/bash

RUN_NAME="otool"

mkdir -p output/bin

go build -o output/bin/${RUN_NAME} main.go

