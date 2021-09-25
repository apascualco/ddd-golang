package signin

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apascualco/apascualco-auth/internal/platform/bus/command"
	"github.com/apascualco/apascualco-auth/internal/platform/bus/event"
	mock_auth "github.com/apascualco/apascualco-auth/internal/platform/storage/mockmysql"
	"github.com/apascualco/apascualco-auth/internal/signin"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignin(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Given a valid user and password should return status ok", func(t *testing.T) {
		// Given
		gin.SetMode(gin.TestMode)
		g := gin.New()
		em := "apascualco@gmail.com"
		pass := "12345"
		lr := signinRequest{
			Email:    em,
			Password: pass,
		}
		j, err := json.Marshal(lr)
		assert.NoError(t, err)

		ur := mock_auth.NewMockUserRepository(ctrl)
		ur.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
		b := command.NewCommandBus()
		e := event.NewInMemoryEventBus()

		s := signin.NewSignin(ur, e)
		sh := signin.NewSigninCommandHandler(s)
		b.Register(signin.SiginCommandType, sh)

		g.POST("/signin", Signin(b))

		// When
		req, err := http.NewRequest(http.MethodPost, "/signin", bytes.NewBuffer(j))
		assert.NoError(t, err)
		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		// Then
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("Given a invalid email should return bad request", func(t *testing.T) {
		// Given
		gin.SetMode(gin.TestMode)
		g := gin.New()
		em := "apascualcogmail.com"
		pass := "12345"
		lr := signinRequest{
			Email:    em,
			Password: pass,
		}
		j, err := json.Marshal(lr)
		assert.NoError(t, err)

		ur := mock_auth.NewMockUserRepository(ctrl)
		b := command.NewCommandBus()
		e := event.NewInMemoryEventBus()

		s := signin.NewSignin(ur, e)
		sh := signin.NewSigninCommandHandler(s)
		b.Register(signin.SiginCommandType, sh)

		g.POST("/signin", Signin(b))

		// When
		req, err := http.NewRequest(http.MethodPost, "/signin", bytes.NewBuffer(j))
		assert.NoError(t, err)
		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		// Then
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Given a empty password should return bad request", func(t *testing.T) {
		// Given
		gin.SetMode(gin.TestMode)
		g := gin.New()
		em := "apascualcogmail.com"
		pass := ""
		lr := signinRequest{
			Email:    em,
			Password: pass,
		}
		j, err := json.Marshal(lr)
		assert.NoError(t, err)

		ur := mock_auth.NewMockUserRepository(ctrl)
		b := command.NewCommandBus()
		e := event.NewInMemoryEventBus()

		s := signin.NewSignin(ur, e)
		sh := signin.NewSigninCommandHandler(s)
		b.Register(signin.SiginCommandType, sh)

		g.POST("/signin", Signin(b))

		// When
		req, err := http.NewRequest(http.MethodPost, "/signin", bytes.NewBuffer(j))
		assert.NoError(t, err)
		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		// Then
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}
