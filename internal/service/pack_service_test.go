package service

import (
	"github.com/bebefabian/orderpack/internal/models"
	"github.com/bebefabian/orderpack/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPackRepository is a mock implementation of PackRepository
type MockPackRepository struct {
	mock.Mock
}

// GetPacks mocks the GetPacks method.
func (m *MockPackRepository) GetPacks() ([]int, error) {
	args := m.Called()
	// You can assert that the first argument is of type []int and return it along with an error.
	return args.Get(0).([]int), args.Error(1)
}

// UpdatePacks mocks the UpdatePacks method.
func (m *MockPackRepository) UpdatePacks(newPacks []int) error {
	args := m.Called(newPacks)
	return args.Error(0)
}

func TestGetPackSizes_ShouldReturnPackSizes(t *testing.T) {
	mockRepo := new(MockPackRepository)
	service := NewPackService(mockRepo)

	// Expected data
	expectedPacks := []int{500, 1000, 2000}

	// Mock repository behavior
	mockRepo.On("GetPacks").Return(expectedPacks, nil)

	// Test service method
	result, err := service.GetPackSizes()
	assert.Nil(t, err)

	assert.Equal(t, expectedPacks, result, "Pack sizes should match expected values")
	mockRepo.AssertExpectations(t)
}

func TestUpdatePackSizes_ShouldUpdatePackSizes(t *testing.T) {
	mockRepo := new(MockPackRepository)
	service := NewPackService(mockRepo)

	// New pack sizes to update
	newPackSizes := []int{100, 300, 700}

	// Expect UpdatePacks to be called with the new values
	mockRepo.On("UpdatePacks", newPackSizes).Return(nil)

	// Test service method
	service.UpdatePackSizes(newPackSizes)

	// Verify method was called correctly
	mockRepo.AssertExpectations(t)
}

// Test for Exact Match Cases
func TestCalculatePacks_ExactMatch(t *testing.T) {
	mockRepo := new(MockPackRepository)
	mockRepo.On("GetPacks").Return([]int{250, 500, 1000, 2000, 5000}, nil)

	service := NewPackService(mockRepo)

	tests := []struct {
		orderQuantity int
		expected      models.CalculateResponse
	}{
		{
			orderQuantity: 1,
			expected: models.CalculateResponse{
				OrderQuantity: 1,
				Packs:         []models.PackResult{{PackSize: 250, Quantity: 1}},
			},
		},
		{
			orderQuantity: 250,
			expected: models.CalculateResponse{
				OrderQuantity: 250,
				Packs:         []models.PackResult{{PackSize: 250, Quantity: 1}},
			},
		},
		{
			orderQuantity: 251,
			expected: models.CalculateResponse{
				OrderQuantity: 251,
				Packs:         []models.PackResult{{PackSize: 500, Quantity: 1}},
			},
		},
		{
			orderQuantity: 501,
			expected: models.CalculateResponse{
				OrderQuantity: 501,
				Packs: []models.PackResult{{PackSize: 500, Quantity: 1},
					{PackSize: 250, Quantity: 1}},
			},
		},
		{
			orderQuantity: 12001,
			expected: models.CalculateResponse{
				OrderQuantity: 12001,
				Packs: []models.PackResult{{PackSize: 5000, Quantity: 2},
					{PackSize: 2000, Quantity: 1}, {PackSize: 250, Quantity: 1}},
			},
		},
	}

	for _, test := range tests {
		result, err := service.CalculatePacks(test.orderQuantity)
		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
	}

	mockRepo.AssertExpectations(t)
}

func TestCalculatePacks_EdgeCase_500000(t *testing.T) {
	mockRepo := repository.NewMemoryPackRepository()
	mockRepo.UpdatePacks([]int{23, 31, 53})
	service := NewPackService(mockRepo)

	result, err := service.CalculatePacks(500000)
	assert.Nil(t, err)

	expected := []models.PackResult{
		{PackSize: 53, Quantity: 9429},
		{PackSize: 31, Quantity: 7},
		{PackSize: 23, Quantity: 2},
	}

	assert.Equal(t, 500000, result.OrderQuantity)
	assert.ElementsMatch(t, expected, result.Packs)
}
