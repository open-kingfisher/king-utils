package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

var customClaims = CustomClaims{
	ID:    "123456",
	Name:  "test",
	Email: "test@example.kingfisher.com",
	StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
	},
}

var j = JWT{
	SigningKey: []byte("test"),
}

var tokenTmp string

func TestJWTCreateToken(t *testing.T) {
	token, err := j.CreateToken(customClaims)
	if err != nil {
		t.Errorf("Create Token error: %s", err)
	} else {
		tokenTmp = token
		t.Logf("Token: %s", token)
	}
}

func TestJWTParseToken(t *testing.T) {
	claims, err := j.ParseToken(tokenTmp)
	if err != nil {
		t.Errorf("Parse Token error: %s", err)
	} else {
		t.Logf("Token ID: %s", claims.ID)
		t.Logf("Token ExpiresAt: %d", claims.ExpiresAt)
	}
}

func TestJWTRefreshToken(t *testing.T) {
	time.Sleep(5 * time.Second)
	token, err := j.RefreshToken(tokenTmp)
	if err != nil {
		t.Errorf("Parse Token error: %s", err)
	} else {
		t.Logf("Token: %s", token)
		claims, _ := j.ParseToken(token)
		t.Logf("Token ID: %s", claims.ID)
		t.Logf("Token ExpiresAt: %d", claims.ExpiresAt)
	}
}
