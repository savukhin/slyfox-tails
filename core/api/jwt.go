package api

import (
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

type PointClaims struct {
	PointID   uint64    `json:"point_id,omitempty"`
	ExpiresAt time.Time `json:"exp,omitempty"`
	jwt.RegisteredClaims
}

func checkUserToken(tokenValue string, key interface{}) (*jwt.Token, error) {
	userClaims := &UserClaims{}
	userToken, err := jwt.ParseWithClaims(tokenValue, userClaims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, fiber.ErrUnauthorized
	}

	// check for time expired
	diff := time.Until(userClaims.ExpiresAt)
	if diff < 0 {
		return nil, fiber.ErrForbidden
	}

	if !userToken.Valid {
		return nil, fiber.ErrForbidden
	}

	return userToken, nil
}

func checkPointToken(tokenValue string, key interface{}) (*jwt.Token, error) {
	pointClaims := &PointClaims{}
	pointToken, err := jwt.ParseWithClaims(tokenValue, pointClaims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, fiber.ErrUnauthorized
	}

	// check for time expired
	diff := time.Until(pointClaims.ExpiresAt)
	if diff < 0 {
		return nil, fiber.ErrForbidden
	}

	if !pointToken.Valid {
		return nil, fiber.ErrForbidden
	}

	return pointToken, nil
}

func JWTChooser(db *gorm.DB, key interface{}, userHandler fiber.Handler, pointHandler fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenHeader := c.GetReqHeaders()["Authorization"]

		if len(tokenHeader) < 8 || tokenHeader[:7] != "Bearer " {
			return fiber.ErrBadRequest
		}

		tokenHeader = tokenHeader[7:]

		userToken, err := checkUserToken(tokenHeader, key)
		if err == nil {
			c.Locals("user", userToken)
			return userHandler(c)
		}

		pointToken, err := checkPointToken(tokenHeader, key)
		if err == nil {
			c.Locals("point", pointToken)
			return pointHandler(c)
		}

		return fiber.ErrForbidden
	}
}
