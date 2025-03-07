package service

import (
	"github.com/bebefabian/orderpack/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPackRepository is a mock implementation of PackRepository
type MockPackRepository struct {
	mock.Mock
}

func (m *MockPackRepository) GetPacks() []int {
	args := m.Called()
	return args.Get(0).([]int)
}

func (m *MockPackRepository) UpdatePacks(newPacks []int) {
	m.Called(newPacks)
}

func TestGetPackSizes_ShouldReturnPackSizes(t *testing.T) {
	mockRepo := new(MockPackRepository)
	service := NewPackService(mockRepo)

	// Expected data
	expectedPacks := []int{500, 1000, 2000}

	// Mock repository behavior
	mockRepo.On("GetPacks").Return(expectedPacks)

	// Test service method
	result := service.GetPackSizes()

	// Verify
	assert.Equal(t, expectedPacks, result, "Pack sizes should match expected values")
	mockRepo.AssertExpectations(t)
}

func TestUpdatePackSizes_ShouldUpdatePackSizes(t *testing.T) {
	mockRepo := new(MockPackRepository)
	service := NewPackService(mockRepo)

	// New pack sizes to update
	newPackSizes := []int{100, 300, 700}

	// Expect UpdatePacks to be called with the new values
	mockRepo.On("UpdatePacks", newPackSizes).Return()

	// Test service method
	service.UpdatePackSizes(newPackSizes)

	// Verify method was called correctly
	mockRepo.AssertExpectations(t)
}

// Test for Exact Match Cases
func TestCalculatePacks_ExactMatch(t *testing.T) {
	mockRepo := new(MockPackRepository)
	mockRepo.On("GetPacks").Return([]int{250, 500, 1000, 2000, 5000})

	service := NewPackService(mockRepo)

	tests := []struct {
		orderQuantity int
		expected      models.CalculateResponse
	}{
		{
			orderQuantity: 1000,
			expected: models.CalculateResponse{
				OrderQuantity: 1000,
				Packs:         []models.PackResult{{PackSize: 1000, Quantity: 1}},
			},
		},
		{
			orderQuantity: 5000,
			expected: models.CalculateResponse{
				OrderQuantity: 5000,
				Packs:         []models.PackResult{{PackSize: 5000, Quantity: 1}},
			},
		},
	}

	for _, test := range tests {
		result := service.CalculatePacks(test.orderQuantity)
		assert.Equal(t, test.expected, result)
	}

	mockRepo.AssertExpectations(t)
}

// Test for Multiple Pack Combinations
func TestCalculatePacks_MultiplePacks(t *testing.T) {
	mockRepo := new(MockPackRepository)
	mockRepo.On("GetPacks").Return([]int{250, 500, 1000, 2000, 5000})

	service := NewPackService(mockRepo)

	result := service.CalculatePacks(12001)
	expected := models.CalculateResponse{
		OrderQuantity: 12001,
		Packs: []models.PackResult{
			{PackSize: 5000, Quantity: 2},
			{PackSize: 2000, Quantity: 1},
			{PackSize: 250, Quantity: 1},
		},
	}

	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

//Test for Order Quantity Less than Smallest Pack
func TestCalculatePacks_OrderTooSmall(t *testing.T) {
	mockRepo := new(MockPackRepository)
	mockRepo.On("GetPacks").Return([]int{250, 500, 1000})

	service := NewPackService(mockRepo)

	result := service.CalculatePacks(100)
	expected := models.CalculateResponse{
		OrderQuantity: 100,
		Packs:         []models.PackResult{{PackSize: 250, Quantity: 1}}, // Uses smallest available pack
	}

	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}
