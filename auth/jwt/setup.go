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

// Setup ...
func Setup() {
	loadPublicKey()
	loadPrivateKey()
}

func loadPublicKey() {
	publicKeyPem, err := ioutil.ReadFile(viper.GetString("token.public_key_path"))
	if err != nil {
		logger.Log().Fatal("Could not load public key. ", err)
	}

	var err2 error
	publicKey, err2 = jwt.ParseRSAPublicKeyFromPEM(publicKeyPem)
	if err2 != nil {
		logger.Log().Fatal("Could not load public key. ", err2)
	}
}

// loadPrivateKey ...
func loadPrivateKey() {
	privateKeyPem, err := ioutil.ReadFile(viper.GetString("token.private_key_path"))
	if err != nil {
		logger.Log().Fatal("Could not load private key. ", err)
	}

	var err2 error
	privateKey, err2 = jwt.ParseRSAPrivateKeyFromPEM(privateKeyPem)
	if err2 != nil {
		logger.Log().Fatal("Could not load private key. ", err2)
	}
}
