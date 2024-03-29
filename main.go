package main

import (
	"auth/model"
	proto "auth/proto"
	"auth/service"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	mux := runtime.NewServeMux()

	micro, err := service.New(model.Database{
		Ip:       os.Getenv("PSQL_IP"),
		Port:     os.Getenv("PSQL_PORT"),
		Name:     os.Getenv("PSQL_NAME"),
		User:     os.Getenv("PSQL_USER"),
		Password: os.Getenv("PSQL_PASSWORD"),
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
