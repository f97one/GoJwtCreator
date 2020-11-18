package main

import (
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"path/filepath"
	"time"
)

func createJWT(issuer string, expire int, private string) (string, error) {
	fp, err := filepath.Abs(private)
	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadFile(fp)
	if err != nil {
		return "", err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(bytes)
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = issuer
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expire)).Unix()

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
