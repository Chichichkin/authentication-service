package sessions

import (
	"auth/db"
	"auth/model"
	"database/sql"
	"errors"
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

func (s *sessionInfo) CreateTableIfNotExists() error {
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

func (s *sessionInfo) SelectActiveByUserId(userId uint64) (userSession *model.Session, err error) {
	userSession = &model.Session{}
	row, err := s.conn.Query(`
	select id, user_id, device, access_token, refresh_token 
	from $1 where user_id=$2`,
		s.tableName, userId)
	if err != nil {
		return nil, errors.New("can't e")
	}
	err = row.Scan(&userSession.Id, &userSession.UserId, &userSession.Device,
		&userSession.AccessToken, &userSession.RefreshToken)
	if err != nil {
		return nil, err
	}
	return userSession, errors.New("not implemented")
}

func (s sessionInfo) Insert(info *model.Session) (*model.Session, error) {
	panic("implement me")
}

func (s *sessionInfo) RefreshSession(id uint64, newRefreshToken, newAccessToken string) (userSession *model.Session, err error) {
	userSession = &model.Session{}
	if newRefreshToken == "" && newAccessToken == "" {
		return nil, errors.New("no tokens")
	}

	err = s.conn.QueryRow(`
		update $1 set (access_token, refresh_token) = ($2, $3) 
		where id = $4 
		returning id, user_id, device, access_token, refresh_token`,
		s.tableName, newAccessToken, newRefreshToken, id).Scan(&userSession.Id, &userSession.UserId, &userSession.Device,
		&userSession.AccessToken, &userSession.RefreshToken)
	if err != nil {
		return nil, err
	}
	return userSession, nil
}

func (s *sessionInfo) RefreshAccess(id uint64, newAccessToken string) (userSession *model.Session, err error) {
	userSession = &model.Session{}
	if newAccessToken == "" {
		return nil, errors.New("no token")
	}

	err = s.conn.QueryRow(`
		update $1 set (access_token) = $2 
		where id = $3 
		returning id, user_id, device, access_token, refresh_token`,
		s.tableName, newAccessToken, id).Scan(&userSession.Id, &userSession.UserId, &userSession.Device,
		&userSession.AccessToken, &userSession.RefreshToken)
	if err != nil {
		return nil, err
	}
	return userSession, nil
}
