package repository

type ProductStorage interface {
	SaveProduct(Product) error
	GetProduct(string) (*Product, error)
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
