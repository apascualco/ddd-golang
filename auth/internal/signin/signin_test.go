package signin

import (
	"context"
	"testing"

	mock_auth "github.com/apascualco/apascualco-auth/internal/platform/storage/mockmysql"
	mock_event "github.com/apascualco/apascualco-auth/kit/mockkit"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignin(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Given a valid values should not return err", func(t *testing.T) {
		// Given
		ur := mock_auth.NewMockUserRepository(ctrl)
		b := mock_event.NewMockBus(ctrl)

		// When
		p := "12345"
		e := "apascualco@gmail.com"
		ur.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
		b.EXPECT().Publish(gomock.Any(), gomock.Any())
		signin := NewSignin(ur, b)

		err := signin.Signin(context.Background(), e, p)

		// Then
		assert.NoError(t, err)
	})

	t.Run("Given a invalid email should return err", func(t *testing.T) {
		// Given
		ur := mock_auth.NewMockUserRepository(ctrl)
		b := mock_event.NewMockBus(ctrl)

		// When
		p := "12345"
		e := "apascualcogmail.com"
		signin := NewSignin(ur, b)

		err := signin.Signin(context.Background(), e, p)

		// Then
		assert.Error(t, err)
		assert.EqualError(t, err, "The email address is invalid")
	})

	t.Run("Given a invalid email without domain should return err", func(t *testing.T) {
		// Given
		ur := mock_auth.NewMockUserRepository(ctrl)
		b := mock_event.NewMockBus(ctrl)

		// When
		p := "12345"
		e := "apascualco@.com"
		signin := NewSignin(ur, b)

		err := signin.Signin(context.Background(), e, p)

		// Then
		assert.Error(t, err)
		assert.EqualError(t, err, "The email address is invalid")
	})

	t.Run("Given a invalid email without username should return err", func(t *testing.T) {
		// Given
		ur := mock_auth.NewMockUserRepository(ctrl)
		b := mock_event.NewMockBus(ctrl)

		// When
		p := "12345"
		e := "@gmail.com"
		signin := NewSignin(ur, b)

		err := signin.Signin(context.Background(), e, p)

		// Then
		assert.Error(t, err)
		assert.EqualError(t, err, "The email address is invalid")
	})

	t.Run("Given a empty email should return err", func(t *testing.T) {
		// Given
		ur := mock_auth.NewMockUserRepository(ctrl)
		b := mock_event.NewMockBus(ctrl)

		// When
		p := "12345"
		e := ""
		signin := NewSignin(ur, b)

		err := signin.Signin(context.Background(), e, p)

		// Then
		assert.Error(t, err)
		assert.EqualError(t, err, "The email address is invalid")
	})

	t.Run("Given a empty password should return err", func(t *testing.T) {
		// Given
		ur := mock_auth.NewMockUserRepository(ctrl)
		b := mock_event.NewMockBus(ctrl)

		// When
		p := ""
		e := "apascualco@gmail.com"
		signin := NewSignin(ur, b)

		err := signin.Signin(context.Background(), e, p)

		// Then
		assert.Error(t, err)
		assert.EqualError(t, err, "The password should be null or empty for user apascualco@gmail.com")
	})
}
