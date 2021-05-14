package users

import (
	"auth/db"
	"auth/model"
	"database/sql"
	"time"
)

type authInfo struct {
	conn      *sql.DB
	tableName string
}

func New(database model.Database) (model.IUsers, error) {
	conn, err := db.NewConnection(database)
	tableName := "user"
	if err != nil {
		return nil, err
	}
	return &authInfo{conn: conn, tableName: tableName}, nil
}

func (a *authInfo) CreateTableIfNotExists() error {
	_, err := a.conn.Exec(`create table if not exists $1 (
    id         bigserial   primary key,
    created_at timestamp   not null,
    email      text unique not null,
    password   text        not null,
    role       int         not null,
    status     int		   not null)`, a.tableName)
	if err != nil {
		return err
	}
	return nil
}

func (a *authInfo) SelectById(id int64) (*model.User, error) {
	row, err := a.conn.Query(`select id, created_at, email, password, role, status from $1 where id=$2`,
		a.tableName, id)
	if err != nil {
		return nil, err
	}
	authInformation := model.User{}
	err = row.Scan(&authInformation.Id, &authInformation.CreatedAt, &authInformation.Email,
		&authInformation.Password, &authInformation.Role, &authInformation.Status)
	if err != nil {
		return nil, err
	}
	return &authInformation, errors.New("not implemented")
}

func (a *authInfo) SelectByEmail(email string) (*model.User, error) {
	row, err := a.conn.Query(`select id, created_at, email, password, role, status from $1 where email=$2`,
		a.tableName, email)
	if err != nil {
		return nil, err
	}
	authInformation := model.User{}
	err = row.Scan(&authInformation.Id, &authInformation.CreatedAt, &authInformation.Email,
		&authInformation.Password, &authInformation.Role, &authInformation.Status)
	if err != nil {
		return nil, err
	}
	return &authInformation, errors.New("not implemented")
}

func (a *authInfo) Insert(info *model.User) (*model.User, error) {
	authInformation := model.User{}
	info.CreatedAt = time.Now()
	err := a.conn.QueryRow(`insert into $1 (created_at, email, password, role, status) 
		values ($2, $3, $4, $5, $6) returning id`,
		a.tableName, info.CreatedAt, info.Email, info.Password, info.Role, info.Status).Scan(&authInformation.Id,
		&authInformation.CreatedAt, &authInformation.Role, &authInformation.Status, &authInformation.Email,
		&authInformation.Password)
	if err != nil {
		return nil, err
	}
	return &authInformation, errors.New("not implemented")
}

func (a *authInfo) UpdatePassword(id int64, password string) (*model.User, error) {
	authInformation := model.User{}
	err := a.conn.QueryRow(`update $1 set password = $2 where id = $3 returning id, created_at, email, 
		password, role, status`,
		a.tableName, password, id).Scan(&authInformation.Id, &authInformation.CreatedAt, &authInformation.Email,
		&authInformation.Password, &authInformation.Role, &authInformation.Status)
	if err != nil {
		return nil, err
	}
	return &authInformation, errors.New("not implemented")
}

func (a *authInfo) UpdateEmail(id int64, email string) (*model.User, error) {
	authInformation := model.User{}
	err := a.conn.QueryRow(`update $1 set email = $2 where id = $3 returning id, created_at, email, 
		password, role, status`,
		a.tableName, email, id).Scan(&authInformation.Id, &authInformation.CreatedAt, &authInformation.Email,
		&authInformation.Password, &authInformation.Role, &authInformation.Status)
	if err != nil {
		return nil, err
	}
	return &authInformation, errors.New("not implemented")
}

func (a *authInfo) Delete(id int64) error {
	_, err := a.conn.Exec(`delete from $1 where id = $2`, a.tableName, id)
	if err != nil {
		return err
	}
	return errors.New("not implemented")
}
