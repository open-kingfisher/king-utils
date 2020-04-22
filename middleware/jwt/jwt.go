package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/open-kingfisher/king-utils/common"
	"github.com/open-kingfisher/king-utils/common/log"
	"net/http"
	"strings"
	"time"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's malformed token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

type CustomClaims struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	ProductId string `json:"product_id"`
	jwt.StandardClaims
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if v, ok := err.(*jwt.ValidationError); ok {
			if v.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if v.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if v.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		responseData := common.ResponseData{
			Msg:  "Unauthorized",
			Data: "",
			Code: http.StatusUnauthorized,
		}
		// 从HTTP请求头获取”X-Signing“
		var token string
		if token = c.Request.Header.Get(common.HeaderSigning); token == "" {
			token = c.Request.Header.Get("Authorization")
		}
		if token == "" {
			c.JSON(http.StatusUnauthorized, responseData)
			c.Abort()
		} else {
			if s := strings.Split(token, " "); len(s) == 2 {
				token = s[1]
			}
			j := JWT{
				SigningKey: []byte(common.Signing),
			}
			// 进行Token解析
			claims, err := j.ParseToken(token)
			if err != nil {
				log.Errorf("Parse token error: %s", err.Error())
				responseData.Msg = "Unauthorized " + err.Error()
				c.JSON(http.StatusUnauthorized, responseData)
				c.Abort()
			}
			// gin上下文中定义用户变量
			c.Set("user", claims)
		}
	}
}
