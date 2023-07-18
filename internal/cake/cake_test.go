package cake

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock implementation of the CakeRepository interface
type mockCakeRepository struct {
	mock.Mock
}

func (m *mockCakeRepository) Create(req *CreateCakeRequest) (*CreateCakeResponse, error) {
	args := m.Called(req)
	return args.Get(0).(*CreateCakeResponse), args.Error(1)
}

func (m *mockCakeRepository) GetAll() ([]*Cake, error) {
	args := m.Called()
	return args.Get(0).([]*Cake), args.Error(1)
}

func (m *mockCakeRepository) GetByID(id int) (*Cake, error) {
	args := m.Called(id)
	return args.Get(0).(*Cake), args.Error(1)
}

func (m *mockCakeRepository) Update(req *UpdateCakeRequest) (*Cake, error) {
	args := m.Called(req)
	return args.Get(0).(*Cake), args.Error(1)
}

func (m *mockCakeRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCakeService_CreateCake(t *testing.T) {
	mockRepo := &mockCakeRepository{}
	service := NewCakeService(mockRepo)

	req := &CreateCakeRequest{
		Title:       "Chocolate Cake",
		Description: "Delicious chocolate cake",
		Rating:      5,
		Image:       "chocolate_cake.jpg",
	}

	expectedCake := &CreateCakeResponse{
		ID:          1,
		Title:       "Chocolate Cake",
		Description: "Delicious chocolate cake",
		Rating:      5,
		Image:       "chocolate_cake.jpg",
	}

	mockRepo.On("Create", req).Return(expectedCake, nil)

	createdCake, err := service.CreateCake(req)

	assert.NoError(t, err)
	assert.Equal(t, expectedCake, createdCake)
	mockRepo.AssertExpectations(t)
}

func TestCakeService_GetAllCakes(t *testing.T) {
	mockRepo := &mockCakeRepository{}
	service := NewCakeService(mockRepo)

	expectedCakes := []*Cake{
		{ID: 1, Title: "Chocolate Cake", Rating: 5},
		{ID: 2, Title: "Vanilla Cake", Rating: 4},
	}

	mockRepo.On("GetAll").Return(expectedCakes, nil)

	cakes, err := service.GetAllCakes()

	assert.NoError(t, err)
	assert.Equal(t, expectedCakes, cakes)
	mockRepo.AssertExpectations(t)
}

func TestCakeService_GetCakeByID(t *testing.T) {
	mockRepo := &mockCakeRepository{}
	service := NewCakeService(mockRepo)

	expectedCake := &Cake{ID: 1, Title: "Chocolate Cake", Rating: 5}

	mockRepo.On("GetByID", 1).Return(expectedCake, nil)

	cake, err := service.GetCakeByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedCake, cake)
	mockRepo.AssertExpectations(t)
}

func TestCakeService_UpdateCake(t *testing.T) {
	mockRepo := &mockCakeRepository{}
	service := NewCakeService(mockRepo)

	req := &UpdateCakeRequest{
		ID:          1,
		Title:       "New Chocolate Cake",
		Description: "Delicious new chocolate cake",
		Rating:      5,
		Image:       "new_chocolate_cake.jpg",
	}

	expectedCake := &Cake{
		ID:          1,
		Title:       "New Chocolate Cake",
		Description: "Delicious new chocolate cake",
		Rating:      5,
		Image:       "new_chocolate_cake.jpg",
	}

	mockRepo.On("Update", req).Return(expectedCake, nil)

	updatedCake, err := service.UpdateCake(req)

	assert.NoError(t, err)
	assert.Equal(t, expectedCake, updatedCake)
	mockRepo.AssertExpectations(t)
}

func TestCakeService_DeleteCake(t *testing.T) {
	mockRepo := &mockCakeRepository{}
	service := NewCakeService(mockRepo)

	mockRepo.On("Delete", 1).Return(nil)

	err := service.DeleteCake(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCakeService_CreateCake_Error(t *testing.T) {
	mockRepo := &mockCakeRepository{}
	service := NewCakeService(mockRepo)

	req := &CreateCakeRequest{
		Title:       "Chocolate Cake",
		Description: "Delicious chocolate cake",
		Rating:      5,
		Image:       "chocolate_cake.jpg",
	}

	mockRepo.On("Create", req).Return(&CreateCakeResponse{}, errors.New("failed to create cake"))

	createdCake, err := service.CreateCake(req)

	assert.Error(t, err)
	assert.Nil(t, createdCake)
	mockRepo.AssertExpectations(t)
}

func TestCakeService_GetAllCakes_Error(t *testing.T) {
	mockRepo := &mockCakeRepository{}
	service := NewCakeService(mockRepo)

	expectedCakes := []*Cake{
		{ID: 1, Title: "Chocolate Cake", Rating: 5, Image: "chocolate_cake.jpg"},
		{ID: 2, Title: "Vanilla Cake", Rating: 4, Image: "vanilla_cake.jpg"},
	}

	mockRepo.On("GetAll").Return(expectedCakes, errors.New("failed to retrieve cakes"))

	cakes, err := service.GetAllCakes()

	assert.Error(t, err)
	assert.Nil(t, cakes)
	mockRepo.AssertExpectations(t)
}

func TestCakeService_GetCakeByID_Error(t *testing.T) {
	mockRepo := &mockCakeRepository{}
	service := NewCakeService(mockRepo)

	mockRepo.On("GetByID", 1).Return(&Cake{}, errors.New("failed to retrieve cake"))

	cake, err := service.GetCakeByID(1)

	assert.Error(t, err)
	assert.Nil(t, cake)
	mockRepo.AssertExpectations(t)
}

func TestCakeService_UpdateCake_Error(t *testing.T) {
	mockRepo := &mockCakeRepository{}
	service := NewCakeService(mockRepo)

	req := &UpdateCakeRequest{
		ID:          1,
		Title:       "New Chocolate Cake",
		Description: "Delicious new chocolate cake",
		Rating:      5,
		Image:       "new_chocolate_cake.jpg",
	}

	mockRepo.On("Update", req).Return(&Cake{}, errors.New("failed to update cake"))

	updatedCake, err := service.UpdateCake(req)

	assert.Error(t, err)
	assert.Nil(t, updatedCake)
	mockRepo.AssertExpectations(t)
}

func TestCakeService_DeleteCake_Error(t *testing.T) {
	mockRepo := &mockCakeRepository{}
	service := NewCakeService(mockRepo)

	mockRepo.On("Delete", 1).Return(errors.New("failed to delete cake"))

	err := service.DeleteCake(1)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
