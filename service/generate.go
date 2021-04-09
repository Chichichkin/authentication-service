package service

import (
	_ "google.golang.org/grpc"
)

//go:generate protoc -I. -I$GOPATH\src -I$GOPATH\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis --go_out=plugins=grpc:./ ./proto/*.proto --grpc-gateway_out=logtostderr=true:./
//go:generate protoc  -I. -I$GOPATH\src -I$GOPATH\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis -I$GOPATH\src\github.com\grpc-ecosystem\grpc-gateway --swagger_out=allow_merge=true,merge_file_name=api:./proto --go_out=plugins=grpc:./ ./proto/*.proto ./proto/*.proto
