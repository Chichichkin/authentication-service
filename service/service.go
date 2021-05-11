package service

import (
	"auth/db/auth_info"
	"auth/model"
	proto "auth/proto"
	"context"
	"errors"
)

type handler struct {
	db model.IAuthInfo
}

func New(database model.Database) (proto.AuthServiceServer, error) {
	conn, err := auth_info.New(database)
	if err != nil {
		return nil, err
	}

	ret := &handler{db: conn}
	if err = ret.db.CreateTableIfNotExists(); err != nil {
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
