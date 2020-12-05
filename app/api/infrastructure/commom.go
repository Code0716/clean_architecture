package infrastructure

import (
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetUuid() string {
	u, err := uuid.NewRandom()
	if err != nil {
		err.Error()
	}
	uuID := u.String()
	return uuID
}

// GetNewToken is get new token
func GetNewToken(email, id string) (tokenString string) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = uuid.New()
	claims["email"] = email
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// 電子署名
	tokenString, _ = token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

	return
}

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

// パスワードハッシュを作る
func passwordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// パスワードがハッシュにマッチするかどうかを調べる
func passwordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}
