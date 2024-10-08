package product

import (
	"database/sql"

	"github.com/nadeem-baig/go-auth/types"
)

type ProductStore interface {
	GetProducts() ([]types.Product, error)
}

type Store struct {
	db *sql.DB
}

var _ ProductStore = (*Store)(nil)

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}
	return products, nil
}

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
	err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Image, &product.Price, &product.Quantity, &product.CreatedAt)
	if err != nil {
		return nil, err
	}
	return product, nil
}
