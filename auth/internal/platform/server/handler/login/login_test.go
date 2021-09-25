package login

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	auth "github.com/apascualco/apascualco-auth/internal"
	"github.com/apascualco/apascualco-auth/internal/login"
	"github.com/apascualco/apascualco-auth/internal/platform/bus/event"
	"github.com/apascualco/apascualco-auth/internal/platform/bus/query"
	mock_auth "github.com/apascualco/apascualco-auth/internal/platform/storage/mockmysql"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	secret := "secret"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Given a valid user and password should return status ok", func(t *testing.T) {
		// Given
		gin.SetMode(gin.TestMode)
		g := gin.New()
		em := "apascualco@gmail.com"
		pass := "12345"
		lr := loginRequest{
			Email:    em,
			Password: pass,
		}
		id, err := uuid.Parse("2877235f-d853-4aea-b72a-07ea512a9cfa")
		assert.NoError(t, err)
		email, err := auth.NewEmail(em)
		assert.NoError(t, err)
		passw, err := auth.NewPassword(pass)
		assert.NoError(t, err)
		user, err := auth.NewUser(auth.NewUUID(id), email, passw)
		assert.NoError(t, err)
		j, err := json.Marshal(lr)
		assert.NoError(t, err)

		ur := mock_auth.NewMockUserRepository(ctrl)
		ur.EXPECT().SearchUserByEmail(gomock.Any(), em).Return(user, nil)
		b := query.NewQueryBus()
		e := event.NewInMemoryEventBus()

		l := login.NewLogin(ur, e, secret)
		loginHandler := login.NewLoginQueryHandler(l)
		b.Register(login.LoginQueryType, loginHandler)

		g.POST("/login", Login(b))

		// When
		req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(j))
		assert.NoError(t, err)
		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		// Then
		assert.Equal(t, http.StatusOK, res.StatusCode)
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		assert.True(t, len(buf.String()) > 10)
	})

	t.Run("Given a valid user and invalid user should return 401 code", func(t *testing.T) {
		// Given
		gin.SetMode(gin.TestMode)
		g := gin.New()
		em := "apascualco@gmail.com"
		pass := "12345"
		lr := loginRequest{
			Email:    em,
			Password: "badpass",
		}
		id, err := uuid.Parse("2877235f-d853-4aea-b72a-07ea512a9cfa")
		assert.NoError(t, err)
		email, err := auth.NewEmail(em)
		assert.NoError(t, err)
		passw, err := auth.NewPassword(pass)
		assert.NoError(t, err)
		user, err := auth.NewUser(auth.NewUUID(id), email, passw)
		assert.NoError(t, err)
		j, err := json.Marshal(lr)
		assert.NoError(t, err)

		ur := mock_auth.NewMockUserRepository(ctrl)
		ur.EXPECT().SearchUserByEmail(gomock.Any(), em).Return(user, nil)
		b := query.NewQueryBus()
		e := event.NewInMemoryEventBus()

		l := login.NewLogin(ur, e, secret)
		loginHandler := login.NewLoginQueryHandler(l)
		b.Register(login.LoginQueryType, loginHandler)

		g.POST("/login", Login(b))

		// When
		req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(j))
		assert.NoError(t, err)
		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		// Then
		assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	})

	t.Run("Given a valid user but not found should return 401", func(t *testing.T) {
		// Given
		gin.SetMode(gin.TestMode)
		g := gin.New()
		em := "apascualco@gmail.com"
		pass := "12345"
		lr := loginRequest{
			Email:    em,
			Password: pass,
		}
		j, err := json.Marshal(lr)
		assert.NoError(t, err)

		ur := mock_auth.NewMockUserRepository(ctrl)
		ur.EXPECT().SearchUserByEmail(gomock.Any(), em).Return(auth.User{}, nil)
		b := query.NewQueryBus()
		e := event.NewInMemoryEventBus()

		l := login.NewLogin(ur, e, secret)
		loginHandler := login.NewLoginQueryHandler(l)
		b.Register(login.LoginQueryType, loginHandler)

		g.POST("/login", Login(b))

		// When
		req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(j))
		assert.NoError(t, err)
		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		// Then
		assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	})
}
