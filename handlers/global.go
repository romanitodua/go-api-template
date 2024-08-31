package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v3"
	"go-api-template/domain"
)

func (h *Handler) Status(c fiber.Ctx) error {
	return c.JSON("ok")
}
func (h *Handler) MigrateDB(c fiber.Ctx) error {
	err := h.repos.MigrateDB()
	if err != nil {
		return err
	}
	return c.JSON("ok")
}

func (h *Handler) Error(c fiber.Ctx, err error) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	var fe *fiber.Error
	var e *domain.Error
	switch {
	case errors.As(err, &fe):
		return c.Status(fe.Code).SendString(fe.Error())
	case errors.As(err, &e):
		switch e.Type {
		case domain.ErrTypeValidation:
			c.Status(400)
		case domain.ErrTypePermissions:
			c.Status(403)
		case domain.ErrTypeNotFound:
			c.Status(404)
		default:
			c.Status(500)
		}
		return c.SendString(e.Error())
	}
	return c.Status(500).SendString(err.Error())
}
