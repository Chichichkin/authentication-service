package main

import (
	_ "google.golang.org/grpc"
)

//go:generate protoc -I. -I$GOPATH\src -I=$GOPATH\pkg\mod\github.com\grpc-ecosystem\grpc-gateway@v1.16.0\third_party\googleapis --go_out=plugins=grpc,paths=source_relative:./ --grpc-gateway_out=logtostderr=true:./ --swagger_out=allow_merge=true,merge_file_name=api:./proto ./proto/*.proto
