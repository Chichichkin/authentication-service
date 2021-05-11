package sessions

import (
	"auth/db"
	"auth/model"
	"database/sql"
)

type sessionInfo struct {
	conn      *sql.DB
	tableName string
}

func New(database model.Database) (model.ISessions, error) {
	conn, err := db.NewConnection(database)
	tableName := "session"
	if err != nil {
		return nil, err
	}
	return &sessionInfo{conn: conn, tableName: tableName}, nil
}

func (s sessionInfo) CreateTableIfNotExists() error {
	_, err := s.conn.Exec(`create table if not exists $1 (
    id         	bigserial primary key,
    userid		bigserial,
    device 		text,
    token 		text,
    deletedat 	timestamp,
    expiredat 	timestamp not null,
    createdat 	timestamp not null,
    upadtedat 	timestamp not null,
    created_at 	timestamp not null)`, s.tableName)
	if err != nil {
		return err
	}
	return nil
}

func (s sessionInfo) SelectActiveByUserId(userId int64) ([]*model.Session, error) {
	panic("implement me")
}

func (s sessionInfo) Insert(info *model.Session) (*model.Session, error) {
	panic("implement me")
}

func (s sessionInfo) RefreshSession(id int64) error {
	panic("implement me")
}

func (s sessionInfo) Delete(id int64) error {
	panic("implement me")
}
