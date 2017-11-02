package auth

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func AuthWrapper(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate token
		req := c.Request
		token, err := request.ParseFromRequest(req, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			return []byte(Secret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
		} else {
			if token.Valid == false {
				c.JSON(http.StatusUnauthorized, "Unauthorized access")
			}
		}
		return
		// token is validate
		// execute function passed by argument
		handlerFunc(c)
	}

}
