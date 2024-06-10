#!/bin/bash
# Build and run the Go-BnB application
# Ignore test files to make running the application easier

go build -o Go-BnB cmd/web/*.go && ./Go-BnB