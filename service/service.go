package service

import (
	"auth/db/sessions"
	"auth/db/users"
	myJwt "auth/jwt"
	"auth/model"
	"auth/proto"
	"context"
	"errors"
	"strconv"
	"time"
)

type handler struct {
	userDb    model.IUsers
	sessionDb model.ISessions
}

func New(database model.Database) (proto.AuthServiceServer, error) {
	userConn, err := users.New(database)
	if err != nil {
		return nil, err
	}
	sessConn, err := sessions.New(database)
	if err != nil {
		return nil, err
	}

	ret := &handler{
		userDb:    userConn,
		sessionDb: sessConn,
	}
	if err = ret.userDb.CreateTableIfNotExists(); err != nil {
		return nil, err
	}
	if err = ret.sessionDb.CreateTableIfNotExists(); err != nil {
		return nil, err
	}

	return ret, nil
}

func (h *handler) Register(context context.Context, req *proto.RegisterRequest) (resp *proto.RegisterResponse, err error) {
	email := req.GetEmail()
	password := req.GetPassword()
	retypePassword := req.GetRetypedPassword()
	userAgent := req.GetUserAgent()

	if password != retypePassword {
		return nil, errors.New("password mismatch") // TODO это ошибка?
	}

	newUser := model.User{
		Id:        0,
		CreatedAt: time.Time{},
		Role:      0,
		Status:    0,
		Email:     email,
		Password:  password,
	}
	res, err := h.userDb.Insert(&newUser)
	if err != nil {
		return nil, err
	}
	tokens, err := myJwt.CreateTokens(res.Id)
	if err != nil {
		return nil, err
	}
	newSession := &model.Session{
		UserId:       res.Id,
		Device:       userAgent,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
	sessRes, err := h.sessionDb.Insert(newSession)
	if err != nil {
		return nil, err
	}
	resp = &proto.RegisterResponse{
		UserId:       res.Id,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		SessionId:    sessRes.Id,
	}
	return resp, nil
}

func (h *handler) Login(context context.Context, req *proto.LoginRequest) (resp *proto.LoginResponse, err error) {
	email := req.GetEmail()
	password := req.GetPassword()
	userAgent := req.GetUserAgent()

	user, err := h.userDb.SelectByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.Password != password { // тут должны сравниваться хэши
		return nil, errors.New("incorrect password")
	}
	tokens, err := myJwt.CreateTokens(user.Id)
	if err != nil {
		return nil, err
	}
	newSession := &model.Session{
		UserId:       user.Id,
		Device:       userAgent,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
	sessRes, err := h.sessionDb.Insert(newSession)
	if err != nil {
		return nil, err
	}
	resp = &proto.LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		SessionId:    sessRes.Id,
	}
	return resp, nil
}

func (h *handler) Check(context context.Context, req *proto.CheckRequest) (resp *proto.CheckResponse, err error) {
	resp = &proto.CheckResponse{
		TokenFine: false,
		Err:       "",
	}
	token := req.GetAccessToken()
	sessId := req.GetSessionId()
	session, err := h.sessionDb.SelectActiveByUserId(sessId)
	if err != nil {
		resp.Err = err.Error()
		return nil, err
	}
	err = myJwt.AccessTokenValidation(token, session.AccessToken, strconv.FormatUint(sessId, 10))
	if err != nil {
		resp.Err = err.Error()
		return resp, err
	}
	resp.TokenFine = true
	return resp, nil
}

func (h *handler) UpdAccess(context context.Context, req *proto.UpdTokenRequest) (resp *proto.UpdTokenResponse, err error) {
	resp = &proto.UpdTokenResponse{
		NewToken: "",
		Err:      "",
	}
	token := req.GetOldToken()
	sessId := req.GetSessionId()
	session, err := h.sessionDb.SelectActiveByUserId(sessId)
	if err != nil {
		resp.Err = err.Error()
		return resp, err
	}
	if session.AccessToken != token {
		resp.Err = errors.New("token doesn't match").Error()
		return resp, errors.New("token doesn't match")
	}
	newToken, err := myJwt.CreateAccessToken(session.UserId)
	if err != nil {
		resp.Err = err.Error()
		return resp, err
	}
	session, err := h.sessionDb.RefreshAccess(sessId, newToken.AccessToken)
	return nil, nil
}

func (h *handler) UpdRefresh(context context.Context, req *proto.UpdTokenRequest) (resp *proto.UpdTokenResponse, err error) {
	token := req.GetOldToken()
	return nil, nil
}
