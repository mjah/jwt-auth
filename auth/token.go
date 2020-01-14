package auth

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const issuer = "auth-server"
const expireIn = time.Hour * 24

var privateKey *rsa.PrivateKey

func loadPrivateKey(keyPtr *string) {
	privateKeyPem, err := ioutil.ReadFile(*keyPtr)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	var err2 error
	privateKey, err2 = jwt.ParseRSAPrivateKeyFromPEM(privateKeyPem)
	if err2 != nil {
		log.Fatal("Error: ", err2)
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
