package example

import (
	"context"
	"errors"
	"math/rand"
)

type CreateUserRequest struct {
	Name  string
	Email string
}

type CreateUserResponse struct {
	Id int
}

func CreateUser(ctx context.Context, req *CreateUserRequest, repository UserRepository) (*CreateUserResponse, error) {
	existingUser, err := repository.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("a user with this email already exists")
	}

	user := &User{
		Id:    rand.Int(),
		Name:  req.Name,
		Email: req.Email,
	}

	if err = repository.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return &CreateUserResponse{Id: user.Id}, nil
}
