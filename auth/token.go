package auth

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mjah/jwt-auth/logger"
	"github.com/spf13/viper"
)

const issuer = "auth-server"
const expireIn = time.Hour * 24

var publicKey *rsa.PublicKey
var privateKey *rsa.PrivateKey

// LoadPublicKey ...
func LoadPublicKey() {
	publicKeyPem, err := ioutil.ReadFile(viper.GetString("token.public_key_path"))
	if err != nil {
		logger.Log().Fatal("Error: ", err)
	}

	var err2 error
	publicKey, err2 = jwt.ParseRSAPublicKeyFromPEM(publicKeyPem)
	if err2 != nil {
		logger.Log().Fatal("Error: ", err2)
	}
}

// LoadPrivateKey ...
func LoadPrivateKey() {
	privateKeyPem, err := ioutil.ReadFile(viper.GetString("token.private_key_path"))
	if err != nil {
		logger.Log().Fatal("Error: ", err)
	}

	var err2 error
	privateKey, err2 = jwt.ParseRSAPrivateKeyFromPEM(privateKeyPem)
	if err2 != nil {
		logger.Log().Fatal("Error: ", err2)
	}
}

// ValidateToken ...
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return publicKey, nil
	})
}

// IssueToken ...
func IssueToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": issuer,
		"exp": time.Now().Add(expireIn).Unix(),
		"grp": "admin",
	})

	return token.SignedString(privateKey)
}

// RefreshToken ...
func RefreshToken() {

}

// RevokeToken ...
func RevokeToken() {

}
