package auth_info

import (
	"auth/service/db"
	"auth/service/model"
	"database/sql"
	"errors"
	"time"
)

type authInfo struct {
	conn *sql.DB
}

var tableName string = "auth"

func New(database model.Database) (model.IAuthInfo, error) {
	conn, err := db.NewConnection(database)
	if err != nil {
		return nil, err
	}
	return &authInfo{conn: conn}, nil
}

func (a *authInfo) CreateTableIfNotExists() error {
	_, err := a.conn.Exec(`create table if not exists $1 (
    id         int primary key,
    created_at timestamp   not null,
    email      text unique not null,
    password   text        not null,
    role       int         not null,
    status     int		   not null)`, tableName)
	if err != nil {
		return err
	}
	return nil
}

func (a *authInfo) SelectById(id int64) (*model.AuthInfo, error) {
	row, err := a.conn.Query(`select id, created_at, email, password, role, status from $1 where id=$2`, tableName, id)
	if err != nil {
		return nil, err
	}
	authInformation := model.AuthInfo{}
	err = row.Scan(&authInformation.Id, &authInformation.CreatedAt, &authInformation.Email,
		&authInformation.Password, &authInformation.Role, &authInformation.Status)
	if err != nil {
		return nil, err
	}
	return &authInformation, errors.New("not implemented")
}

func (a *authInfo) Insert(info *model.AuthInfo) (*model.AuthInfo, error) {
	authInformation := model.AuthInfo{}
	info.CreatedAt = time.Now()
	err := a.conn.QueryRow(`insert into $1 (created_at, email, password, role, status) 
		values ($2, $3, $4, $5, $6) returning id`,
		tableName, info.CreatedAt, info.Email, info.Password, info.Role, info.Status).Scan(&authInformation.Id)
	if err != nil {
		return nil, err
	}
	return &authInformation, errors.New("not implemented")
}

func (a *authInfo) UpdatePassword(id int64, password string) (*model.AuthInfo, error) {
	authInformation := model.AuthInfo{}
	err := a.conn.QueryRow(`update $1 set password = $2 where id = $3 returning id, created_at, email, 
		password, role, status`,
		tableName, password, id).Scan(&authInformation.Id, &authInformation.CreatedAt, &authInformation.Email,
		&authInformation.Password, &authInformation.Role, &authInformation.Status)
	if err != nil {
		return nil, err
	}
	return &authInformation, errors.New("not implemented")
}

func (a *authInfo) UpdateEmail(id int64, email string) (*model.AuthInfo, error) {
	authInformation := model.AuthInfo{}
	err := a.conn.QueryRow(`update $1 set email = $2 where id = $3 returning id, created_at, email, 
		password, role, status`,
		tableName, email, id).Scan(&authInformation.Id, &authInformation.CreatedAt, &authInformation.Email,
		&authInformation.Password, &authInformation.Role, &authInformation.Status)
	if err != nil {
		return nil, err
	}
	return &authInformation, errors.New("not implemented")
}

func (a *authInfo) Delete(id int64) error {
	_, err := a.conn.Exec(`delete from $1 where id = $2`, tableName, id)
	if err != nil {
		return err
	}
	return errors.New("not implemented")
}
