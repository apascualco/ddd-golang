package login

import (
	"net/http"

	"github.com/apascualco/apascualco-auth/internal/login"
	"github.com/apascualco/apascualco-auth/kit/query"
	"github.com/gin-gonic/gin"
)

// swagger:parameters LoginRequest
type bodyLoginRequest struct {
	// in: body
	Body loginRequest

	// in: header
	// example: application/vnd.auth.v1+json
	Accept string `json:"accept" binding:"required"`
}

type loginRequest struct {
	// description: email
	// example: your@email.com
	Email string `json:"email" binding:"required"`

	// description: password
	// example: 1234
	Password string `json:"password" binding:"required"`
}

// swagger:route POST /login login LoginRequest
//
// Login
//
//	Consumes:
//  - application/json
//
//	Schemes: http, https
//
//	Responses:
//		200:
//		400:
//		500:
func Login(bus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req loginRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		token, err := bus.Dispatch(ctx, login.NewLoginQuery(req.Email, req.Password))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, nil)
			return
		}
		ctx.JSON(http.StatusOK, token)
	}
}
