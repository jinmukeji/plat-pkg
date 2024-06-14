package jwt

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrParseClaimsFailed = errors.New("failed to parse not standard claims")
	ErrNoPublicKey       = errors.New("no public key to verity JWT")
	ErrEmptyToken        = errors.New("token is empty")
)

// CreateClaims 根据 issuer 与过期时间间隔创建一个 JWT Claims.
// 例如，issuer 可以是一个 APP ID.
func CreateClaims(issuer string, inr time.Duration) jwt.Claims {
	now := time.Now()
	claims := jwt.StandardClaims{
		IssuedAt:  now.Unix(),          // iat
		ExpiresAt: now.Add(inr).Unix(), // exp
		Issuer:    issuer,              // iss
	}
	return &claims
}

// RS256SignJWT 使用 RS256 算法对 claims 进行签名
func RS256SignJWT(claims jwt.Claims, key *rsa.PrivateKey) (string, error) {
	return signJWT(jwt.SigningMethodRS256, key, claims)
}

// RS512SignJWT 使用 RS512 算法对 claims 进行签名
func RS512SignJWT(claims jwt.Claims, key *rsa.PrivateKey) (string, error) {
	return signJWT(jwt.SigningMethodRS512, key, claims)
}

func signJWT(method jwt.SigningMethod, key interface{}, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(method, claims)
	ss, err := token.SignedString(key)
	return ss, err
}

// GetPublicKeyFunc 根据 iss 获取一个 rsa.PublicKey
type GetPublicKeyFunc func(iss string) *rsa.PublicKey

// GetPublicKeyByKidFunc 根据 kid 获取一个 rsa.PublicKey
type GetPublicKeyByKidFunc func(kid string) *rsa.PublicKey

// VerifyOption 验证参数
type VerifyOption struct {
	MaxExpInterval   time.Duration    // 最大过期时间间隔，单位为秒
	GetPublicKeyFunc GetPublicKeyFunc // PublicKey 查找函数
}

// KidVerifyOption 验证参数
type KidVerifyOption struct {
	MaxExpInterval   time.Duration         // 最大过期时间间隔，单位为秒
	GetPublicKeyFunc GetPublicKeyByKidFunc // PublicKey 查找函数
}

// HMACVerifyOption 验证参数
type HMACVerifyOption struct {
	MaxExpInterval time.Duration // 最大过期时间间隔，单位为秒
	// app-key与密钥的对应关系
	// key: app-key
	// value: 密钥的 byte 数组
	SecretKeys map[string][]byte
}

// RSAVerifyCustomJWT 使用 RSA 算法 (RS256/RS384/RS512) 对包含自定义 Claims 的 JWT Token 进行验证。
func RSAVerifyCustomJWT(tokenString string, opt VerifyOption, claims Claims) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		if opt.GetPublicKeyFunc != nil {
			key := opt.GetPublicKeyFunc(claims.GetIssuer())

			if key != nil {
				return key, nil
			}
		}

		return nil, ErrNoPublicKey
	})

	if token == nil {
		return false, ErrEmptyToken
	}

	if token.Valid {
		inr := float64(claims.GetExpiresAt() - claims.GetIssuedAt())
		if inr > opt.MaxExpInterval.Seconds() {
			return false, fmt.Errorf("expiration interval exceeds the limit: %fs", opt.MaxExpInterval.Seconds())
		}
		return true, nil
	} else {
		return false, err
	}
}

// RSAVerifyJWT 使用 RSA 算法 (RS256/RS384/RS512) 对 JWT Token 进行验证。
func RSAVerifyJWT(tokenString string, opt VerifyOption) (bool, *jwt.StandardClaims, error) {
	var stdClaims *jwt.StandardClaims

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// 获取 claims
		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok {
			return nil, ErrParseClaimsFailed
		}

		stdClaims = claims

		if opt.GetPublicKeyFunc != nil {
			key := opt.GetPublicKeyFunc(claims.Issuer)

			if key != nil {
				return key, nil
			}
		}

		return nil, ErrNoPublicKey
	})

	if token == nil {
		return false, nil, ErrEmptyToken
	}

	if token.Valid {
		inr := float64(stdClaims.ExpiresAt - stdClaims.IssuedAt)
		if inr > opt.MaxExpInterval.Seconds() {
			return false, nil, fmt.Errorf("expiration interval exceeds the limit: %fs", opt.MaxExpInterval.Seconds())
		}
		return true, stdClaims, nil
	} else {
		return false, nil, err
	}
}

// RSAVerifyJWTWithKid 使用 RSA 算法 (RS256/RS384/RS512) 对 JWT Token 进行验证。
func RSAVerifyJWTWithKid(tokenString string, opt KidVerifyOption) (bool, *jwt.StandardClaims, error) {
	var stdClaims *jwt.StandardClaims

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// 获取 claims
		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok {
			return nil, ErrParseClaimsFailed
		}

		stdClaims = claims

		// 获取 header 中的 kid
		kkid, ok := token.Header["kid"]
		if !ok {
			return nil, fmt.Errorf("jwt header kid not found")
		}

		kid, ok := kkid.(string)
		if !ok {
			return nil, fmt.Errorf("can not convert kid to string")
		}

		if opt.GetPublicKeyFunc != nil {
			key := opt.GetPublicKeyFunc(kid)

			if key != nil {
				return key, nil
			}
		}

		return nil, ErrNoPublicKey
	})

	if token == nil {
		return false, nil, ErrEmptyToken
	}

	if token.Valid {
		inr := float64(stdClaims.ExpiresAt - stdClaims.IssuedAt)
		if inr > opt.MaxExpInterval.Seconds() {
			return false, nil, fmt.Errorf("expiration interval exceeds the limit: %fs", opt.MaxExpInterval.Seconds())
		}

		return true, stdClaims, nil
	} else {
		return false, nil, err
	}
}

func HMACVerifyCustomJWT(tokenString string, opt HMACVerifyOption, claims Claims) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// 获取 header 中的	kid
		kkid, ok := token.Header["kid"]
		if !ok {
			return nil, fmt.Errorf("jwt header kid not found")
		}

		kid, ok := kkid.(string)
		if !ok {
			return nil, fmt.Errorf("can not convert kid to string")
		}

		// 如果 kid 对应的 key 找不到则报错
		secretKey, ok := opt.SecretKeys[kid]
		if !ok {
			return nil, fmt.Errorf("unable to find matching key for kid: %s", kid)
		}

		return secretKey, nil
	})

	if token == nil {
		return false, ErrEmptyToken
	}

	if token.Valid {
		inr := float64(claims.GetExpiresAt() - claims.GetIssuedAt())
		if inr > opt.MaxExpInterval.Seconds() {
			return false, fmt.Errorf("expiration interval exceeds the limit: %fs", opt.MaxExpInterval.Seconds())
		}

		return true, nil
	} else {
		return false, err
	}
}
