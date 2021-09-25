package login

import (
	"context"
	"fmt"
	"testing"

	auth "github.com/apascualco/apascualco-auth/internal"
	mock_auth "github.com/apascualco/apascualco-auth/internal/platform/storage/mockmysql"
	mock_event "github.com/apascualco/apascualco-auth/kit/mockkit"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {

	secret := "secret"
	p := "12345"
	e := "apascualco@gmail.com"
	uid, err := uuid.Parse("ad856379-cea6-4859-857b-d8225dea206d")
	assert.NoError(t, err)

	uuid := auth.NewUUID(uid)
	email, err := auth.NewEmail(e)
	assert.NoError(t, err)

	password, err := auth.NewPassword(p)
	assert.NoError(t, err)

	user, err := auth.NewUser(uuid, email, password)
	assert.NoError(t, err)

	t.Run("Given valid values should return token", func(t *testing.T) {
		// Given
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ur := mock_auth.NewMockUserRepository(ctrl)
		b := mock_event.NewMockBus(ctrl)

		ur.EXPECT().SearchUserByEmail(gomock.Any(), e).Return(user, nil)
		b.EXPECT().Publish(gomock.Any(), gomock.Any())

		// When
		login := NewLogin(ur, b, secret)
		tk, err := login.Login(context.Background(), e, p)

		// Then
		assert.NoError(t, err)
		assert.NotNil(t, tk)
		assert.True(t, len(tk) > 10)
	})

	t.Run("Given valid values but bad password should not call event publish and return empty token", func(t *testing.T) {
		// Given
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ur := mock_auth.NewMockUserRepository(ctrl)
		b := mock_event.NewMockBus(ctrl)

		ur.EXPECT().SearchUserByEmail(gomock.Any(), e).Return(user, nil)

		// When
		login := NewLogin(ur, b, secret)
		tk, err := login.Login(context.Background(), e, "badpassword")

		// Then
		assert.NoError(t, err)
		assert.NotNil(t, tk)
		assert.Equal(t, "", tk)
	})

	t.Run("Given valid email but user not found should not call event publish and return empty token", func(t *testing.T) {
		// Given
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ur := mock_auth.NewMockUserRepository(ctrl)
		b := mock_event.NewMockBus(ctrl)

		ur.EXPECT().SearchUserByEmail(gomock.Any(), e).Return(auth.User{}, nil)

		// When
		login := NewLogin(ur, b, secret)
		tk, err := login.Login(context.Background(), e, "badpassword")

		// Then
		assert.NoError(t, err)
		assert.NotNil(t, tk)
		assert.Equal(t, "", tk)
	})

	t.Run("Given valid email but repository return error should return error", func(t *testing.T) {
		// Given
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ur := mock_auth.NewMockUserRepository(ctrl)
		b := mock_event.NewMockBus(ctrl)

		ur.EXPECT().SearchUserByEmail(gomock.Any(), e).Return(auth.User{}, fmt.Errorf("Random error"))

		// When
		login := NewLogin(ur, b, secret)
		_, err := login.Login(context.Background(), e, "badpassword")

		// Then
		assert.Error(t, err)
		assert.EqualError(t, err, "Random error")
	})
}
