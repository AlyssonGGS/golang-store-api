package memory_test

import (
	"testing"

	"github.com/alyssonggs/golang-store-api/repository"
	"github.com/alyssonggs/golang-store-api/repository/memory"
)

func TestGetProduct(t *testing.T) {
	t.Parallel()

	tctc := []struct {
		ID      string
		Name    string
		Product *repository.Product
	}{
		{
			Name:    "success",
			ID:      "1",
			Product: &repository.Product{ID: 1},
		},
	}

	memoryStorage := memory.MemoryStorage{}
	for _, tc := range tctc {
		t.Run(tc.Name, func(t *testing.T) {
			p, err := memoryStorage.GetProduct(tc.ID)
			if err != nil {
				t.Fatal("Error on testing memory", "err", err)
			}

			if p.ID != tc.Product.ID {
				t.Error("Wrong id while recovering product from memory")
			}
		})
	}
}
