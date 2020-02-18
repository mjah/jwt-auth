package jwt

import (
	"crypto/rsa"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var publicKey *rsa.PublicKey
var privateKey *rsa.PrivateKey

// Setup ...
func Setup() error {
	if err := loadPublicKey(); err != nil {
		return err
	}

	if err := loadPrivateKey(); err != nil {
		return err
	}

	return nil
}

func loadPublicKey() error {
	publicKeyPem, err := ioutil.ReadFile(viper.GetString("token.public_key_path"))
	if err != nil {
		return err
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyPem)
	if err != nil {
		return err
	}

	return nil
}

func loadPrivateKey() error {
	privateKeyPem, err := ioutil.ReadFile(viper.GetString("token.private_key_path"))
	if err != nil {
		return err
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyPem)
	if err != nil {
		return err
	}

	return nil
}
