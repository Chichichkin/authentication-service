package service

import (
	"auth/db/users"
	"auth/model"
	"auth/proto"
	"context"
	"errors"
)

type handler struct {
	userDb    model.IUsers
	sessionDb model.ISessions
}

func New(database model.Database) (proto.AuthServiceServer, error) {
	conn, err := users.New(database)
	if err != nil {
		return nil, err
	}

	ret := &handler{userDb: conn}
	if err = ret.userDb.CreateTableIfNotExists(); err != nil {
		return nil, err
	}

	return ret, nil
}

func (h *handler) Register(context.Context, *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	return nil, errors.New("not implemented")
}

func (h *handler) Login(context.Context, *proto.LoginRequest) (*proto.LoginResponse, error) {
	return nil, errors.New("not implemented")
}
