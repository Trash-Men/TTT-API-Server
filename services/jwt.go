package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/Trash-Men/api-server/configs"
	"github.com/dgrijalva/jwt-go"
)

type jwtObject struct {
	Id string
}

type JwtService struct{}

func (_ JwtService) CreateToken(id string, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	jwtConfig := token.Claims.(jwt.MapClaims)

	jwtConfig["id"] = id
	jwtConfig["role"] = role
	jwtConfig["exp"] = time.Now().Add(time.Minute * 50000000).Unix()

	accessToken, error := token.SignedString([]byte(configs.GetEnvironments().JWT_SECRET_KEY))

	if error != nil {
		return "", error
	}

	return accessToken, nil
}

func (_ JwtService) DecodeToken(tokenString string) (jwtObject, error) {
	token, error := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.GetEnvironments().JWT_SECRET_KEY), nil
	})

	if error != nil {
		return jwtObject{}, error
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	type MapClaims map[string]interface{}

	if ok && token.Valid {
		id, ok := claims["id"].(string)

		if !ok {
			return jwtObject{}, errors.New("token decode failed")
		}

		return jwtObject{
			Id: id,
		}, nil
	}

	return jwtObject{}, errors.New("failed to claim jwt")
}
