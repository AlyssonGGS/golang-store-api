package memory

import "github.com/alyssonggs/golang-store-api/repository"

type MemoryStorage struct {
}

func (m *MemoryStorage) GetProduct(id string) repository.Product {
	return repository.Product{}
}

func (m *MemoryStorage) SaveProduct(p repository.Product) error {
	return nil
}
