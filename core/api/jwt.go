package api

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserClaims struct {
	UserID    uint64    `json:"user_id,omitempty"`
	ExpiresAt time.Time `json:"exp,omitempty"`
	jwt.RegisteredClaims
}

// SigningKey holds information about the recognized cryptographic keys used to sign JWTs by this program.
type SigningKey struct {
	// JWTAlg is the algorithm used to sign JWTs. If this value is a non-empty string, this will be checked against the
	// "alg" value in the JWT header.
	//
	// https://www.rfc-editor.org/rfc/rfc7518#section-3.1
	JWTAlg string
	// Key is the cryptographic key used to sign JWTs. For supported types, please see
	// https://github.com/golang-jwt/jwt.
	Key interface{}
}

func signingKeyFunc(key SigningKey) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if key.JWTAlg != "" {
			alg, ok := token.Header["alg"].(string)
			if !ok {
				return nil, fmt.Errorf("unexpected jwt signing method: expected: %q: got: missing or unexpected JSON type", key.JWTAlg)
			}
			if alg != key.JWTAlg {
				return nil, fmt.Errorf("unexpected jwt signing method: expected: %q: got: %q", key.JWTAlg, alg)
			}
		}
		return key.Key, nil
	}
}

func JWTMiddleware(db *gorm.DB, key interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenHeader := c.GetReqHeaders()["Authorization"]
		fmt.Println("Header is ", tokenHeader)

		if len(tokenHeader) < 8 || tokenHeader[:7] != "Bearer " {
			return fiber.ErrBadRequest
		}

		tokenHeader = tokenHeader[7:]
		fmt.Println("Header after filter is ", tokenHeader)

		claims := &UserClaims{}
		tokenValue := tokenHeader
		// jwt.SigningMethodRS256
		token, err := jwt.ParseWithClaims(tokenValue, claims, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})

		if err != nil {
			fmt.Println("userclaims err", err)
			return fiber.ErrUnauthorized
		}
		fmt.Println("token", token)

		// check for time expired
		diff := time.Until(claims.ExpiresAt)
		if diff < 0 {
			fmt.Println("expired diff < 0", diff)
			return fiber.ErrForbidden
		}

		if !token.Valid {
			fmt.Println("token not valid")
			return fiber.ErrForbidden
		}

		c.Locals("user", token)
		return c.Next()
	}
}
