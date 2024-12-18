package controller

import (
	responser "github.com/buwud/goNote/api/errors"
	"github.com/buwud/goNote/domain"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserUseCase domain.UserUseCase
}

func (userController *UserController) SignUp(c *fiber.Ctx) error {
	user := new(domain.UserSignup)
	if err := c.BodyParser(user); err != nil {
		return responser.InvalidBody(c)
	}

	err := userController.UserUseCase.SignUp(user)
	if err != nil {
		return responser.FailedLogin(c)
	}
	return responser.SuccessfulSignup(c)
}
func (userController *UserController) SignIn(c *fiber.Ctx) error {
	user := new(domain.UserSignin)
	if err := c.BodyParser(user); err != nil {
		return responser.InvalidBody(c)
	}

	err := userController.UserUseCase.SignIn(user, c)
	if err != nil {
		return responser.FailedLogin(c)
	}

	return responser.SuccessfulLogin(c)
}

func (userController *UserController) SignOut(c *fiber.Ctx) error {
	userController.UserUseCase.SignOut(c)
	return responser.SuccessfulLogout(c)
}
