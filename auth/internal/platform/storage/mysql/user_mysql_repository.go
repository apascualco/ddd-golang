package mysql

import (
	"context"
	"database/sql"
	"time"

	auth "github.com/apascualco/apascualco-auth/internal"
	"github.com/huandu/go-sqlbuilder"
)

// MysqlUserRepository is a Mysql auth.UserRepository implementation
type MysqlUserRepository struct {
	dbReader  *sql.DB
	dbWriter  *sql.DB
	dbTimeout time.Duration
}

type sqlUser struct {
	Id        []byte    `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	IsActive  bool      `db:"is_active"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

func NewMysqlUserRepository(dbr *sql.DB, dbw *sql.DB, dbTimeout time.Duration) *MysqlUserRepository {
	return &MysqlUserRepository{
		dbReader:  dbr,
		dbWriter:  dbw,
		dbTimeout: dbTimeout,
	}
}

const sqlUserTable = "user"

func mapToSqlUser(c auth.User) (sqlUser, error) {
	uuid, err := c.ID().MarshalBinary()
	if err != nil {
		return sqlUser{}, err
	}
	hp, err := c.HashPassword()
	if err != nil {
		return sqlUser{}, err
	}
	return sqlUser{
		Id:        uuid,
		Email:     c.Email(),
		Password:  hp,
		IsActive:  true,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}, nil
}

func mapToUser(s sqlUser) (auth.User, error) {
	password, err := auth.NewHashPassword(s.Password)
	if err != nil {
		return auth.User{}, err
	}
	id, err := auth.NewUUIDByByte(s.Id)
	if err != nil {
		return auth.User{}, err
	}
	email, err := auth.NewEmail(s.Email)
	if err != nil {
		return auth.User{}, err
	}
	return auth.NewUser(id, email, password)
}

func (r *MysqlUserRepository) Save(ctx context.Context, user auth.User) error {
	userSQLStruct := sqlbuilder.NewStruct(new(sqlUser))
	sq, err := mapToSqlUser(user)
	if err != nil {
		return err
	}
	query, args := userSQLStruct.InsertInto(sqlUserTable, sq).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err = r.dbWriter.ExecContext(ctxTimeout, query, args...)
	return err
}

func (r *MysqlUserRepository) SearchUserByEmail(ctx context.Context, email string) (auth.User, error) {
	userSQLStruct := sqlbuilder.NewStruct(new(sqlUser))
	builder := userSQLStruct.SelectFrom(sqlUserTable)
	builder.Where(builder.Equal("email", email))
	query, args := builder.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	rows, err := r.dbReader.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		return auth.User{}, err
	}
	defer rows.Close()

	var us sqlUser
	rows.Next()
	rows.Scan(userSQLStruct.Addr(&us)...)
	if us.Id != nil {
		return mapToUser(us)
	}
	return auth.User{}, nil

}
