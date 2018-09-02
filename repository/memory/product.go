package memory

import (
	"strconv"

	"github.com/alyssonggs/golang-store-api/repository"
)

type MemoryStorage struct {
}

func (m *MemoryStorage) GetProduct(id string) (*repository.Product, error) {
	ID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return &repository.Product{ID: ID}, nil
}

func (m *MemoryStorage) SaveProduct(p repository.Product) error {
	return nil
}
