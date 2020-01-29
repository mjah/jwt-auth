package auth

import (
	"crypto/rsa"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mjah/jwt-auth/logger"
	"github.com/spf13/viper"
)

const issuer = "auth-server"
const expireIn = time.Hour * 24

var privateKey *rsa.PrivateKey

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

func issueToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": issuer,
		"exp": time.Now().Add(expireIn).Unix(),
		"grp": "admin",
	})

	return token.SignedString(privateKey)
}

func refreshToken() {

}

func revokeToken() {

}
