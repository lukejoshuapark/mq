// This file was automatically generated by mq. DO NOT EDIT.
package example

import (
	"context"
	"github.com/lukejoshuapark/mq"
)

type MockUserRepository struct {
	findByEmailSetups []MockUserRepositoryFindByEmailSetup
	findByEmailCalls []MockUserRepositoryFindByEmailCall
	createUserSetups []MockUserRepositoryCreateUserSetup
	createUserCalls []MockUserRepositoryCreateUserCall
}

var _ UserRepository = &MockUserRepository{}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

type MockUserRepositoryFindByEmailSetup struct {
	ctx mq.Input[context.Context]
	email mq.Input[string]
	o0 mq.Output[*User]
	o1 mq.Output[error]
}

type MockUserRepositoryFindByEmailCall struct {
	ctx context.Context
	email string
}

func (m *MockUserRepository) SetupFindByEmail(ctx mq.Input[context.Context], email mq.Input[string], o0 mq.Output[*User], o1 mq.Output[error]) {
	m.findByEmailSetups = append(m.findByEmailSetups, MockUserRepositoryFindByEmailSetup{
		ctx: ctx,
		email: email,
		o0: o0,
		o1: o1,
	})
}

func (m *MockUserRepository) VerifyFindByEmail(count mq.Count, ctx mq.Input[context.Context], email mq.Input[string]) {
	c := 0

	for _, call := range m.findByEmailCalls {
		if ctx.Compare(call.ctx) && email.Compare(call.email) {
			c++
		}
	}

	if !count.ShouldPass(c) {
		panic("unexpected numbers of calls were made with those arguments")
	}
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	for _, setup := range m.findByEmailSetups {
		if setup.ctx.Compare(ctx) && setup.email.Compare(email) {
			m.findByEmailCalls = append(m.findByEmailCalls, MockUserRepositoryFindByEmailCall{
				ctx: ctx,
				email: email,
			})

			return setup.o0.Value(), setup.o1.Value()
		}
	}

	panic("no setups were created with those arguments")
}

type MockUserRepositoryCreateUserSetup struct {
	ctx mq.Input[context.Context]
	user mq.Input[*User]
	o0 mq.Output[error]
}

type MockUserRepositoryCreateUserCall struct {
	ctx context.Context
	user *User
}

func (m *MockUserRepository) SetupCreateUser(ctx mq.Input[context.Context], user mq.Input[*User], o0 mq.Output[error]) {
	m.createUserSetups = append(m.createUserSetups, MockUserRepositoryCreateUserSetup{
		ctx: ctx,
		user: user,
		o0: o0,
	})
}

func (m *MockUserRepository) VerifyCreateUser(count mq.Count, ctx mq.Input[context.Context], user mq.Input[*User]) {
	c := 0

	for _, call := range m.createUserCalls {
		if ctx.Compare(call.ctx) && user.Compare(call.user) {
			c++
		}
	}

	if !count.ShouldPass(c) {
		panic("unexpected numbers of calls were made with those arguments")
	}
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *User) error {
	for _, setup := range m.createUserSetups {
		if setup.ctx.Compare(ctx) && setup.user.Compare(user) {
			m.createUserCalls = append(m.createUserCalls, MockUserRepositoryCreateUserCall{
				ctx: ctx,
				user: user,
			})

			return setup.o0.Value()
		}
	}

	panic("no setups were created with those arguments")
}
