package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/auth"
	"github.com/jak103/powerplay/internal/server"
)

func init() {
	server.RegisterHandler(fiber.MethodPost, "/auth", auth.NONE, postAuthHandler)

	server.RegisterHandler(fiber.MethodGet, "/user", auth.JWT, getCurrentUser)
	server.RegisterHandler(fiber.MethodPost, "/user", auth.NONE, createUserAccount)
}

func postAuthHandler(c *fiber.Ctx) error {
	return nil
}

func getCurrentUser(c *fiber.Ctx) error {
	return nil
}

func createUserAccount(c *fiber.Ctx) error {
	return nil
}
