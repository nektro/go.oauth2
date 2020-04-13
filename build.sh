#!/usr/bin/env bash

go get -v -u github.com/rakyll/statik
$GOPATH/bin/statik -src="./www/"
