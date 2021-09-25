package signin

import (
	"net/http"

	"github.com/apascualco/apascualco-auth/internal/signin"
	"github.com/apascualco/apascualco-auth/kit/command"
	"github.com/gin-gonic/gin"
)

// swagger:parameters SigninRequest
type bodySigninRequest struct {
	// in: body
	Body signinRequest

	// in: header
	// example: application/vnd.auth.v1+json
	Accept string `json:"accept" binding:"required"`
}

type signinRequest struct {
	// description: email
	// example: your@email.com
	Email string `json:"email" binding:"required"`

	// description: user password
	// example: 1234
	Password string `json:"password" binding:"required"`
}

// swagger:route POST /signin signin SigninRequest
//
// Login
//
//	Consumes:
//	- application/json
//
//	Schemes: http, https
//
//	Responses:
//		200:
//		400:
//		500:
func Signin(bus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req signinRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err := bus.Dispatch(ctx, signin.NewSigninCommand(req.Email, req.Password)); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, nil)
	}
}
