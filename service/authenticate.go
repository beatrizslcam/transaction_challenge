package service

import (
	"time"
	"transactions/config"

	"github.com/golang-jwt/jwt/v5"
)
type Auth interface{
	GenerateToken(string) string
	ValidateToken(string)bool
}
func GenerateToken(accountID string) (string, error) {
	//Recovering the secret of the env var
	jwtKey  := getJWTSecretKey()
	
	
	//creating token
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
		IssuedAt:jwt.NewNumericDate(time.Now()),
		Subject: accountID,

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenSigned, err := token.SignedString(jwtKey)
	if err != nil{
		print("Log adequado aqui")
		return "", err
	}

return tokenSigned, nil
}

func ValidateToken(token string)bool{
	jwtKey  := getJWTSecretKey()
	if jwtKey != nil{
		print("Log adequado aqui")
		return false
	}

	return true
}

func getJWTSecretKey() ([]byte) {
	cfg,err := config.GetConfig()
	if err != nil{
		print("Log adequado aqui")
		return nil
	}
	
	jwtKey := []byte(cfg.JWTSecret)

	return jwtKey

}