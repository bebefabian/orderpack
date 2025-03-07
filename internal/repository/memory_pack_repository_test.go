package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMemoryPackRepository_ShouldStartEmpty(t *testing.T) {
	repo := NewMemoryPackRepository()
	assert.Equal(t, []int{}, repo.GetPacks(), "Pack sizes should start empty")
}

func TestUpdatePacks_ShouldUpdateAndSortDescending(t *testing.T) {
	repo := NewMemoryPackRepository()

	// New pack sizes
	newPackSizes := []int{100, 300, 700, 1500}
	repo.UpdatePacks(newPackSizes)

	// Expected sorted order
	expected := []int{1500, 700, 300, 100}
	assert.Equal(t, expected, repo.GetPacks(), "Pack sizes should be updated and sorted in descending order")
}
