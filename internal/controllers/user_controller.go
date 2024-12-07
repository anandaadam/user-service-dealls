package controllers

import "github.com/gofiber/fiber/v2"

type UserController interface {
	CreateUser(ctx *fiber.Ctx) error
	CreateUserBio(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}
