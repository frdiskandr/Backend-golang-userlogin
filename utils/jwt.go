package utils

import (
    "os"
    "time"

    "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
    UserId uint `json:"user_id"`
    jwt.StandardClaims
}

func GenerateToken(userId uint) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserId: userId,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ValidateToken(tknStr string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        return nil, err
    }
	println(token)
    return claims, nil
}
