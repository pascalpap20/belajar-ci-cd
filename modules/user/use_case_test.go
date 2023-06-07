package user

import (
	"crud/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockUserRepository struct {
	mock.Mock
}

func (m MockUserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserRepository) GetUserById(id uint) (entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m MockUserRepository) UpdateUser(user entity.User) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockUserRepository) DeleteUser(email string) (any, error) {
	//TODO implement me
	panic("implement me")
}

func TestGetUserById(t *testing.T) {
	expectedUser := entity.User{
		ID:       1,
		Name:     "jhon",
		Email:    "email",
		Password: "test123",
	}

	// Create a mock repository instance
	mockRepo := &MockUserRepository{}

	// Set the expected behavior for GetUserById
	mockRepo.On("GetUserById", expectedUser.ID).Return(expectedUser, nil)

	// Create an instance of the use case with the mock repository
	useCase := useCaseUser{
		userRepo: mockRepo,
	}

	// Call the GetUserById function
	user, err := useCase.GetUserById(1)

	// Assertions
	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, expectedUser, user, "Expected user to match")

}
