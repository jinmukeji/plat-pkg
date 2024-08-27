package main

import (
	"crypto/rsa"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/jinmukeji/plat-pkg/v4/auth/jwt"
)

const (
	// DefaultMaxExpirationInterval 默认最大的过期时间间隔（10分钟）
	DefaultMaxExpirationInterval = 10 * time.Minute
)

type MyClaims struct {
	jwt.StandardClaims

	AccessToken string `json:"access_token"`
}

func main() {
	var keyFile, tokenString string

	flag.StringVar(&keyFile, "key", "", "the RSA public key file with PEM format")
	flag.StringVar(&tokenString, "token", "", "the encoded JWT token")
	flag.Parse()

	key, err := jwt.LoadRSAPublicKeyFromPEM(keyFile)
	die(err)

	opt := jwt.VerifyOption{
		MaxExpInterval: DefaultMaxExpirationInterval,
		GetPublicKeyFunc: func(iss string) *rsa.PublicKey {
			// ignore iss check

			return key
		},
	}

	claims := MyClaims{}
	valid, err := jwt.RSAVerifyCustomJWT(tokenString, opt, &claims)
	fmt.Printf("IsValid: %v\n", valid)
	if err != nil {
		fmt.Printf("Validation Error: %v\n", err)
	}
	fmt.Println("Claims:", claims)
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
