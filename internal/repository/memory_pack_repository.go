package repository

import (
	"sort"
	"sync"
)

// MemoryPackRepository is an in-memory implementation of PackRepository
type MemoryPackRepository struct {
	mu        sync.RWMutex
	packSizes []int
}

// NewMemoryPackRepository initializes an empty repository
func NewMemoryPackRepository() *MemoryPackRepository {
	return &MemoryPackRepository{
		packSizes: []int{}, // No default pack sizes
	}
}

// GetPacks returns the current available pack sizes
func (r *MemoryPackRepository) GetPacks() ([]int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.packSizes, nil
}

// UpdatePacks modifies the available pack sizes dynamically
func (r *MemoryPackRepository) UpdatePacks(newPacks []int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Sort pack sizes in descending order for better optimization
	sort.Sort(sort.Reverse(sort.IntSlice(newPacks)))
	r.packSizes = newPacks
	return nil
}
