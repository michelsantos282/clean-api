package repository

import (
	"database/sql"

	"github.com/michelsantos282/clean-api/internal/entity"
)

type ProductRepositoryMysql struct {
	DB *sql.DB
}

// Como se fosse um construtor
func NewProductRepositoryMysql(db *sql.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{DB: db}
}

// Metodo do repositorio
func (r *ProductRepositoryMysql) Create(product *entity.Product) error {
	_, err := r.DB.Exec("INSERT INTO products (id, name, price) values (?, ?, ?)", product.ID, product.Name, product.Price)
	if err != nil {
		return nil
	}
	return nil
}

func (r *ProductRepositoryMysql) FindAll() ([]*entity.Product, error) {
	rows, err := r.DB.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
