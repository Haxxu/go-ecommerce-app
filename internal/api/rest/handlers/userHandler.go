package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-app/internal/api/rest"
)

type UserHandler struct {
	// svc UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	// create an instance of user service & inject to handler
	handler := UserHandler{}

	// Public endpoints
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// Private endpoints
	app.Get("/verify", handler.GetVerificationCode)
	app.Post("/verify", handler.Verify)
	app.Get("/profile", handler.GetProfile)
	app.Post("/profile", handler.CreateProfile)

	app.Get("/cart", handler.GetCart)
	app.Post("/cart", handler.AddToCart)

	app.Get("/order", handler.GetOrders)
	app.Post("/order", handler.CreateOrder)
	app.Get("/order/:id", handler.GetOrder)

	app.Post("/become-seller", handler.BecomeSeller)
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "register success",
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "login success",
	})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "verify success",
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "get verification code success",
	})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "create profile success",
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "get profile success",
	})
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "get cart success",
	})
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "add to cart success",
	})
}

func (h *UserHandler) CreateOrder(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "create order success",
	})
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "get orders success",
	})
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "get order by id success",
	})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "become seller success",
	})
}
