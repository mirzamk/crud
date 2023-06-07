package user

import (
	"crud/entity"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	args := m.Called(user)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(user *entity.User) (interface{}, error) {
	args := m.Called(user)
	return args.Get(0), args.Error(1)
}

func (m *MockUserRepository) DeleteUser(email string) (interface{}, error) {
	args := m.Called(email)
	return args.Get(0), args.Error(1)
}

func (m *MockUserRepository) GetUserById(id uint) (entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(entity.User), args.Error(1)
}

func TestGetUserById(t *testing.T) {
	// Create a mock user repository
	mockRepo := &MockUserRepository{}

	// Create the use case with the mock repository
	useCase := useCaseUser{userRepo: mockRepo}

	// Define the expected user
	expectedUser := entity.User{
		Id:   20,
		Name: "John Doe",
		// Add other fields as necessary
	}

	// Set up the mock repository to return the expected user
	mockRepo.On("GetUserById", expectedUser.Id).Return(entity.User{}, errors.New("mock error"))

	// Call the GetUserById function
	user, err := useCase.GetUserById(expectedUser.Id)

	// Assert that the returned user and error match the expected values
	assert.NoError(t, err)              // No error should occur
	assert.Equal(t, expectedUser, user) // The returned user should match the expected user

	// Optionally, you can assert that the methods on the mock repository were called as expected
	mockRepo.AssertExpectations(t)
}
