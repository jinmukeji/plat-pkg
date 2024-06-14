package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jinmukeji/plat-pkg/v4/auth/jwt"
)

func main() {
	var keyFile, appId, interval, passFile string

	flag.StringVar(&keyFile, "key", "", "the RSA private key file with PEM format")
	flag.StringVar(&appId, "iss", "", "the issuer")
	flag.StringVar(&interval, "inr", "120s", "the time interval between iat and exp")
	flag.StringVar(&passFile, "pass", "", "the private key password file")
	flag.Parse()

	dur, err := time.ParseDuration(interval)
	die(err)

	key, err := jwt.LoadRSAPrivateKeyFromPEM(keyFile, passFile)
	die(err)

	claims := jwt.CreateClaims(appId, dur)

	ss, err := jwt.RS256SignJWT(claims, key)
	die(err)

	fmt.Println(ss)
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
