# SOVD Server with gRPC Gateway

## Description
This project is designed to map RestAPIs/HTTP1.1 request to gRPC/Protobuf services. SOVD server is created with gateway between SOVD Client and gRPC server. SOVD APIs are created to diagnose the application manager in a distrbuted system.

## Project structure
- proto/application_manager/sovd-grpc-interface.proto: have the SOVD APIs and its correcspoing gRPC services.
- proxy/main.go: for registering gRPC server and route requests to various registered services 
- buf.gen.yaml: is buf configuration file

## Requirements
Following is the requirements and steps on Linux based systems:
- Download Homebrew:
`/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`

- Protocol buffer compiler:
`brew install protobuf`

- Ensure compiler version is 3+ : 
    `protoc --version`

- Install go (for sovd proxy)
    `sudo apt update && sudo apt install golang -y`

- Install buf CLI
    `brew install bufbuild/buf/buf`

## Run the SOVD Server
- `go run proxy/main.go`


