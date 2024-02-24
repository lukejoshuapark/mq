package example

import (
	"context"
	"testing"

	"github.com/lukejoshuapark/mq"
)

type CreateUserFixture struct {
	repository *MockUserRepository

	req *CreateUserRequest

	existingUser      *User
	existingUserError error

	createUserError error
}

func SetupCreateUserTest() *CreateUserFixture {
	fixture := &CreateUserFixture{
		repository: NewMockUserRepository(),

		req: &CreateUserRequest{
			Name:  "Luke",
			Email: "luke@test.com",
		},

		existingUser:      nil,
		existingUserError: nil,

		createUserError: nil,
	}

	fixture.repository.SetupFindByEmail(
		mq.IsAny[context.Context](),
		mq.IsExactly("luke@test.com"),
		mq.Returns(&fixture.existingUser),
		mq.Returns(&fixture.existingUserError))

	fixture.repository.SetupCreateUser(
		mq.IsAny[context.Context](),
		mq.IsAny[*User](),
		mq.Returns(&fixture.createUserError))

	return fixture
}

func TestCreateUserSuccess(t *testing.T) {
	// Arrange.
	fixture := SetupCreateUserTest()

	// Act.
	res, err := CreateUser(context.Background(), fixture.req, fixture.repository)

	// Assert.
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	if res == nil {
		t.Errorf("expected response to not be nil")
	}

	fixture.repository.VerifyCreateUser(
		mq.Once,
		mq.IsAny[context.Context](),
		mq.Is(func(x *User) bool {
			return x.Email == "luke@test.com" && x.Name == "Luke"
		}))
}

func TestCreateUserAlreadyExists(t *testing.T) {
	// Arrange.
	fixture := SetupCreateUserTest()
	fixture.existingUser = &User{}

	// Act.
	res, err := CreateUser(context.Background(), fixture.req, fixture.repository)

	// Assert.
	if err == nil {
		t.Errorf("expected error to not be nil")
	}

	if res != nil {
		t.Errorf("expected response to be nil")
	}

	if err.Error() != "a user with this email already exists" {
		t.Errorf("expected error to be 'a user with this email already exists', got %v", err)
	}

	fixture.repository.VerifyCreateUser(
		mq.Never,
		mq.IsAny[context.Context](),
		mq.IsAny[*User]())
}
