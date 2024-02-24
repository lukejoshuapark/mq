//go:generate mq generate --input $GOFILE
package example

import "context"

type User struct {
	Id    int
	Name  string
	Email string
}

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, user *User) error
}
