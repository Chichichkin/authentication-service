package model

import (
	"database/sql"
	"time"
)

type ISessions interface {
	CreateTableIfNotExists() error
	SelectActiveByUserId(userId int64) ([]*Session, error)
	Insert(info *Session) (*Session, error)
	RefreshSession(id int64) error
	Delete(id int64) error
}

type Session struct {
	Id        uint64       `json:"id"` // id сессии
	UserId    uint64       `json:"user_id"`
	Device    string       `json:"device"`
	Token     string       `json:"token"`
	DeletedAt sql.NullTime `json:"deleted_at"`
	ExpiredAt time.Time    `json:"expired_at"` // до какого времени сессия не протухла
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
