package service

import (
	repository "github.com/Mpetrato/startupconnect/src/repositories"
	"github.com/Mpetrato/startupconnect/src/types"
	"github.com/Mpetrato/startupconnect/src/utils"

	"github.com/google/uuid"
)

func RegisterUserService(user *types.User) (string, error) {
	id := uuid.New().String()
	user.Id = id

	EncryptedPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return "bug", err
	}
	user.Password = EncryptedPass

	existingUser, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		return "Error Email already exists!", err
	}

	if existingUser != nil {
		return "Email already exists!", err
	}

	err = repository.Register(user)
	if err != nil {
		return "Error creating account", err
	}

	return "Account has been created!", err
}
