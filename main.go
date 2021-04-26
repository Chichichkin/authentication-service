package main

import (
	"auth/model"
	"auth/proto"
	"auth/service"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	mux := runtime.NewServeMux()

	micro, err := service.New(model.Database{
		Ip:       "",
		Port:     "",
		Name:     "",
		User:     "",
		Password: "",
	})
	if err != nil {
		panic(err)
	}

	if err = proto.RegisterAuthServiceHandlerServer(context.Background(), mux, micro); err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		panic(err.Error())
	}

	proto.RegisterAuthServiceServer(grpcServer, micro)

	log.Println(fmt.Sprintf("Listening on %s port", "5300"))

	go func() {
		grpcServer.Serve(listener)
	}()

	log.Println(fmt.Sprintf("HTTP API Listening on %s port", "3000"))
	http.ListenAndServe(fmt.Sprintf(":%s", "3000"), mux)
}
