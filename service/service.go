package service

import (
	"auth/db/users"
	"auth/jwt"
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

func (h *handler) Register(context context.Context, registerRequest *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	email := registerRequest.GetEmail()
	password := registerRequest.GetPassword()
	retypePassword := registerRequest.GetRetypedPassword()

	if password != retypePassword {
		return nil, errors.New("password mismatch") // TODO это ошибка?
	}

	newUser := model.User{Role: 0, Email: email, Password: password, Status: 0} // TODO статус не всегда должен быть 0
	h.userDb.Insert(&newUser)

	return nil, errors.New("not implemented") // TODO тут должен возвращаться токен
}

func (h *handler) Login(context context.Context, loginRequest *proto.LoginRequest) (*proto.LoginResponse, error) {
	email := loginRequest.GetEmail()
	password := loginRequest.GetPassword()

	user, err := h.userDb.SelectByEmail(email)
	if err != nil {
		return nil, errors.New("error in h.userDb.SelectByEmail(email)") // TODO выяснить какие ошибки должны быть
	}

	if user.Password != password { // тут должны сравниваться хэши
		return nil, errors.New("incorrect password")
	}

	return nil, errors.New("not implemented") // TODO тут должен возвращаться токен
}

// TODO реализовать
func (h *handler) Check(context context.Context, CheckRequest *proto.CheckRequest) (*proto.CheckResponse, error) {
	return nil, nil
}

func (h *handler) UpdAccess(context context.Context, UpdTokenRequest *proto.UpdTokenRequest) (*proto.UpdTokenResponse, error) {
	return nil, nil
}

func (h *handler) UpdRefresh(context context.Context, UpdTokenRequest *proto.UpdTokenRequest) (*proto.UpdTokenResponse, error) {
	return nil, nil
}
