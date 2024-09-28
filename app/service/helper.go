package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func gensalt() (string, error) {
	salt := make([]byte, 8)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(salt), nil
}

func encryptPassword(password, salt string) string {
	passwordStr := password + salt
	hash := sha256.Sum256([]byte(passwordStr))
	return fmt.Sprintf("%x", hash)
}

type authCustomClaims struct {
	id uint
	jwt.StandardClaims
}

func generateToken(id uint, oAuthKey string) string {
	claims := &authCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(oAuthKey))
	if err != nil {
		panic(err)
	}
	return t
}
