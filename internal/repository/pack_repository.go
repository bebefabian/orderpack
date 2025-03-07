package repository

type PackRepository interface {
	GetPacks() ([]int, error)
	UpdatePacks(newPacks []int) error
}
