package jwt

import (
	"crypto/rsa"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/mjah/jwt-auth/logger"
	"github.com/spf13/viper"
)

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
