package api

import "github.com/gofiber/fiber/v2"

func NotImplemented(c *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

func Pong(c *fiber.Ctx) error {
	return c.SendString("pong")
}
