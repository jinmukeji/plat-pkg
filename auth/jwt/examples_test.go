package jwt_test

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/jinmukeji/plat-pkg/v2/auth/jwt"
)

func ExampleRSAVerifyCustomJWT() {
	// MyClaims is a custom claims
	type MyClaims struct {
		jwt.StandardClaims

		AccessToken string `json:"access_token"`
	}

	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTQwNTMyMTQsImlhdCI6MTU5NDA1MjYxNCwiaXNzIjoiYXBwLXRlc3QxIn0.Xj2bALCrcIMHLHmeeI7ipRddoxU21MmigH3EBr9T_wygkZiZyzOOs-KU2VKuwMhnVsI0vU1iQKs0lCoHt8hSUGddHBjQ4oXcgfo9LWeKl0mluAeVzuBVsI-cZqDAapn5vKRrHvw2IsF-luJNB9th9-HY3_4Nif7OOKGc7DoYkzy-gazKl1lqOH76cy9jQBZ_FNYyKKh28_FgBECxoOogAfakyclPLfXjIxqvpAMMYYp3x0Gbeb1NtRToLNEHeJBEAs1W3vgCQ9i3DF2F1PP3XKHWifUp6MANMgt3w1ghPxxUK2MRHe1oX6wnu652GtspKQ0EJq5GnWMTie0KdRZCfw"
	key, err := jwt.LoadRSAPublicKeyFromPEM("public_key.pem")
	if err != nil {
		panic(err)
	}

	opt := jwt.VerifyOption{
		MaxExpInterval: 10 * time.Minute,
		GetPublicKeyFunc: func(iss string) *rsa.PublicKey {
			// ignore iss check

			return key
		},
	}

	claims := MyClaims{}

	valid, err := jwt.RSAVerifyCustomJWT(token, opt, &claims)
	fmt.Printf("IsValid: %v\n", valid)
	if err != nil {
		fmt.Printf("Validation Error: %v\n", err)
	}
	fmt.Println("Claims:", claims)
}

func ExampleHMACVerifyCustomJWT() {
	// MyClaims is a custom claims
	type MyClaims struct {
		jwt.StandardClaims

		AccessToken string `json:"access_token"`
	}

	claims := &MyClaims{}

	m := make(map[string][]byte)
	m["a"] = []byte("BQysRAXxfa4MjD5ta6p51AULAdQc1bGHJJVWsRRNQCTvqZpztWm3sJErB7MgZYYeqQkdkxpT0xyjhXDoySZdraq7OHcqksQCccIHtDHqu0ujrug4qI78EGgPeeZASpKqxnVibqDLqvpnFrb8BTrIfRz8VXe4Ncv4DIZLyqUMoILflIJvabtfuv1i51km4BIPIDR6Vvw5pratnEqcLgNQipd25fHooEZtj1X70oF3A0uVFggnmljk6XEbSL3ZbEIs")
	opt := jwt.HMACVerifyOption{
		MaxExpInterval: 10 * time.Minute,
		SecretKeys:     m,
	}

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjb20uamlubXVoZWFsdGguaGprIiwiZXhwIjoxNjIyNDQ0MjkxLCJzdWIiOiJoamsiLCJpYXQiOjE2MjI0NDM5OTEsImFjY2Vzc190b2tlbiI6IjJhYzc4ZGFjLTY4YTMtNDZjYi1iNzYxLWZmMjFmMjEwMGI3MSJ9.4A26iyNXPAZWemIs5P68Z2dOSciAh7IkqX8ldsLyQas"
	ok, err := jwt.HMACVerifyCustomJWT(token, opt, claims)
	if err != nil {
		panic(err)
	}
	fmt.Println(ok)
	fmt.Println(claims)
}

func ExampleRSAVerifyJWTWithKid() {

	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImhteSJ9.eyJpc3MiOiJjb20uamlubXVoZWFsdGguaG15IiwiZXhwIjoxNjI0OTM2NDIwLCJzdWIiOiJobXkiLCJpYXQiOjE2MjQ5MzYxMjB9.rSEgdLncgTtec19dV7VDo0tr_nTbfXl2qVqW96ezRR7aM5MPHuppSVCs-bgFkBiEEXGqKPNxOYprEKlSmAXwQGhQ7HGc6vdCe1GE0GqK0j0Bs_kihicmUCAP9AZU-hoqN98wm4wBO-z51Tc1Sio8ZTRig7ICp3OvlCvA9ZkFg694WuCSJNBIG-8JEUzJxNY1kaXwlxN7jQLW_zyNrFAeIyOSTCeITgL9a7VOA85l0VB36mjBY30uZNyOmUOnAurukfYkQxlEpU9d0E0vVcvtcpszU-ahT53WoNHmSWhdfcTkU9eGUucV0RNUQKVHdkqU75gx5diCO5F8mQIfzAJ_Eg"
	key, err := jwt.LoadRSAPublicKeyFromPEM("./a.pem")
	if err != nil {
		panic(err)
	}

	opt := jwt.KidVerifyOption{
		MaxExpInterval: 10 * time.Minute,
		GetPublicKeyFunc: func(iss string) *rsa.PublicKey {
			// ignore iss check

			return key
		},
	}

	valid, claims, err := jwt.RSAVerifyJWTWithKid(token, opt)
	fmt.Printf("IsValid: %v\n", valid)
	if err != nil {
		fmt.Printf("Validation Error: %v\n", err)
	}
	fmt.Println("Claims:", claims)
}
