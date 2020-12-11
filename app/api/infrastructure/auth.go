package infrastructure

import (
	"log"
	"os"
	"time"

	"github.com/Code0716/clean_architecture/app/api/interfaces/controllers"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetUuid() string {
	u, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)

	}
	uuID := u.String()
	return uuID
}

// GetNewToken is get new token
func getNewToken(id, name, email string) (tokenString string) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = id
	claims["name"] = name
	claims["email"] = email
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// 電子署名
	tokenString, _ = token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

	return
}

func validateJWT(c *gin.Context, executionFunc func(controllers.Context)) {
	//  署名の検証
	_, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(os.Getenv("SIGNINGKEY"))
		return b, nil
	})
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized."})
		return
	}
	executionFunc(c)
}

// passwordHash make hash
func passwordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// PasswordVerify check hash
func passwordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}
