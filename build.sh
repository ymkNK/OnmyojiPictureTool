#!/bin/bash

RUN_NAME="otool"

mkdir -p output/bin/config output/bin/templates
cp -r templates/. output/bin/templates
cp config/config.yaml output/bin/config/

go build -o output/bin/${RUN_NAME} main.go

