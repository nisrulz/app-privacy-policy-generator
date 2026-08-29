#!/usr/bin/env bash
cd "$(dirname "$0")/.." || exit 1

go get -u ./...
go mod tidy
