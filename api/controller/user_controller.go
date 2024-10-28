package controller

import (
	"github.com/buwud/goNote/domain"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	UserUseCase domain.UserUseCase
}

func (userController *UserController) SignUp(c *fiber.Ctx) error {
	user := new(domain.UserSignup)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	result, err := userController.UserUseCase.SignUp(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return c.Status(fiber.StatusOK).JSON(result)
}