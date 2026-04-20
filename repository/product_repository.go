package repository

import (
	"database/sql"
	"fmt"

	"github.com/misaelmodesto/simple-crud-go/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GETProducts() ([]model.Product, error) {
	query := "select id, product_name, price from product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productOBJ model.Product

	for rows.Next() {
		err = rows.Scan(
			&productOBJ.ID,
			&productOBJ.Name,
			&productOBJ.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productOBJ)

	}

	rows.Close()

	return productList, nil
}
