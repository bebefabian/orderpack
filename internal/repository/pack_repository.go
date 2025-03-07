package repository

type PackRepository interface {
	GetPacks() []int
	UpdatePacks(newPacks []int)
}
