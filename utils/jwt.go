package utils

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

const signKey = "supersecret"

func GenerateJWT(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	signedValue, err := token.SignedString([]byte(signKey))
	if err != nil {
		log.Printf("Unable to generate jwt token: %v\n", err)
		return signedValue, err
	}
	return signedValue, nil
}

func ValidateToken(tokenString string) (parsedToken *jwt.Token, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			log.Printf("Unexpected signing method\n: %v", token.Header["alg"])
			return nil, errors.New("unexpected signing method")
		}
		return []byte(signKey), nil
	})
	if err != nil {
		log.Printf("\nError validating token: %v", token.Header["alg"])
		return nil, err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, err
	}
	return token, nil
}

func ExtractData(token *jwt.Token) int64 {
	return int64(token.Claims.(jwt.MapClaims)["userId"].(float64))
}
