package middlewares

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"bff/order_web/global"
	"bff/order_web/models"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That not even a token")
	TokenInvalid     = errors.New("token invalid")
)

type JWT struct {
	SigningKey []byte
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "请登录",
			})
			c.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"msg": "授权已过期",
				})

			} else {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"msg": "重新登录",
				})
			}
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Set("userId", claims.ID)
		c.Next()
	}
}

func NewJWT() *JWT {
	return &JWT{SigningKey: []byte(global.ServerConfig.JWTInfo.SigningKey)}
}

func (j *JWT) CreateToken(claims models.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*models.CustomClaims, error) {
	tokenCustomClaims, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
		return nil, err
	}
	if tokenCustomClaims != nil {
		if claims, ok := tokenCustomClaims.Claims.(*models.CustomClaims); ok && tokenCustomClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
