# Install and generate proto files

## Installation

go install google.golang.org/protobuf/cmd/protoc-gen-go
or
go get -u github.com/golang/protobuf/protoc-gen-go

## Generating protobuf files

protoc --go_out=plugins=grpc:. *.proto

The below command will generate _pb.go files in the directory nokia.com/srlinux/sdk/protos
