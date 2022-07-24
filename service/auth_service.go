package service

import (
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
)

func VerifyUser(tokenString string) (int, interface{}) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(("your-256-bit-secret")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// fmt.Println(claims["userId"])
		i, _ := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))	

		return i, nil
	} else {
		// fmt.Println(err)
		return 0, err

	}

}
