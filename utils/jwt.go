package utils

import (
	"fmt"
	"os"
	"server/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct{
	ID   uint  `json:"id"`
	Name string	 `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
func GenerateToken(user dto.UserResponse) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	claims := &Claims{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "",err
	}
	return tokenString,err
}

func VerifyToken(tokenString string)(*Claims, error){
	secretKey := []byte("hellodek")
	claims := &Claims{}
	// tokens,err := jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {

	// 	return secretKey,nil
	// })
	tokens, err := jwt.ParseWithClaims(tokenString,claims,func(t *jwt.Token) (interface{}, error) {
		return secretKey,nil
	})
	if err != nil{
		return nil,err
	}
	if !tokens.Valid{
		return nil,fmt.Errorf("invalid token")
	}
	//fmt.Println("ini verifikasi jwt")
	//fmt.Println(tokens.Claims.GetSubject())
	return claims,nil
}
