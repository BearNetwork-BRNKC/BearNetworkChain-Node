#!/bin/bash

set -e -x


xcodebuild -version

go run build/ci.go install -dlgo
go run build/ci.go archive -type tar

