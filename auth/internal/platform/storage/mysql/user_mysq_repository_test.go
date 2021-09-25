package mysql

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	auth "github.com/apascualco/apascualco-auth/internal"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMysqlUserRepository(t *testing.T) {

	dbr, sqlMockr, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)

	dbw, sqlMockw, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)

	t.Run("Given a valid user should not return err when save user", func(t *testing.T) {
		// Given
		idString := "d85ecb79-403c-443a-99cb-9ae3916a4961"
		id, err := uuid.Parse(idString)
		assert.NoError(t, err)
		uuid := auth.NewUUID(id)
		e := "apascualco@gmail.com"
		email, err := auth.NewEmail(e)
		assert.NoError(t, err)
		p := "1234meomfeofm"
		password, err := auth.NewPassword(p)
		assert.NoError(t, err)

		// When
		user, err := auth.NewUser(uuid, email, password)
		assert.NoError(t, err)
		binaryId, err := user.ID().MarshalBinary()
		assert.NoError(t, err)

		sqlMockw.ExpectExec(
			"INSERT INTO user (id, email, password, is_active, updated_at, created_at) VALUES (?, ?, ?, ?, ?, ?)").
			WithArgs(binaryId, e, sqlmock.AnyArg(), true, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(0, 1))
		r := NewMysqlUserRepository(dbr, dbw, 1*time.Millisecond)

		// Then
		err = r.Save(context.Background(), user)
		assert.NoError(t, err)
	})

	t.Run("Given a valid email should search the user", func(t *testing.T) {
		email := "apascualco@gmail.com"

		rows := sqlmock.NewRows([]string{
			"id",
			"email",
			"password",
			"is_active",
			"updated_at",
			"created_at"})

		sqlMockr.ExpectQuery(
			"SELECT user.id, user.email, user.password, user.is_active, user.updated_at, user.created_at FROM user WHERE email = ?").
			WithArgs(email).
			WillReturnRows(rows)
		r := NewMysqlUserRepository(dbr, dbw, 1*time.Millisecond)

		_, err = r.SearchUserByEmail(context.Background(), email)

		assert.NoError(t, sqlMockr.ExpectationsWereMet())
		assert.NoError(t, err)

	})
}
