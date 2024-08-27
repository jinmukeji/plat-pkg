package jwt

import "github.com/dgrijalva/jwt-go"

type Claims interface {
	jwt.Claims

	// GetIssuer 返回 iss
	GetIssuer() string
	// GetExpiresAt 返回 exp
	GetExpiresAt() int64
	// GetIssuedAt 返回 iat
	GetIssuedAt() int64
}

// StandardClaims is a wrapper for jwt.StandardClaims
type StandardClaims struct {
	jwt.StandardClaims
}

func (c *StandardClaims) GetIssuer() string {
	if c == nil {
		return ""
	}
	return c.Issuer
}

func (c *StandardClaims) GetExpiresAt() int64 {
	if c == nil {
		return 0
	}
	return c.ExpiresAt
}

func (c *StandardClaims) GetIssuedAt() int64 {
	if c == nil {
		return 0
	}
	return c.IssuedAt
}
