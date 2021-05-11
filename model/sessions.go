package model

type ISessions interface {
	CreateTableIfNotExists() error
	SelectActiveByUserId(userId int64) ([]*Session, error)
	Insert(info *Session) (*Session, error)
	RefreshSession(id int64) error
	Delete(id int64) error
}

type Session struct {
	Id           uint64 `json:"id"` // id сессии
	UserId       uint64 `json:"user_id"`
	Device       string `json:"device"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
