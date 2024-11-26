package usecase

import (
	"github.com/buwud/goNote/domain"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) (domain.UserUseCase, error) {
	return &userUseCase{userRepo: userRepo}, nil
}

func (u *userUseCase) SignUp(user *domain.UserSignup) (*mongo.InsertOneResult, error) {
	return u.userRepo.SignUp(user)
}

func (u *userUseCase) SignIn(user *domain.UserSignin, c *fiber.Ctx) (string, error) {
	return u.userRepo.SignIn(user, c)
}

func (u *userUseCase) SignOut(c *fiber.Ctx) {
	u.userRepo.SignOut(c)
}
