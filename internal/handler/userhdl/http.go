package userhdl

import (
	"http-server/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

type HTTPHandler struct {
	userService ports.UserService
}

func NewHTTPHandler(userService ports.UserService) *HTTPHandler {
	return &HTTPHandler{
		userService: userService,
	}
}

func (h *HTTPHandler) Login(c *fiber.Ctx) error {
	return nil
}

func (h *HTTPHandler) Create(c *fiber.Ctx) error {
	return nil
}

func (h *HTTPHandler) Get(c *fiber.Ctx) error {
	return nil
}

func (h *HTTPHandler) List(c *fiber.Ctx) error {
	return nil
}

func (h *HTTPHandler) Update(c *fiber.Ctx) error {
	return nil
}

func (h *HTTPHandler) Delete(c *fiber.Ctx) error {
	return nil
}
