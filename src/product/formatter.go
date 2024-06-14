package product

import "go-technical-test-synapsis/src/entity"

type ProductFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

func FormatProduct(product entity.Product) ProductFormatter {
	formatter := ProductFormatter{
		ID:          product.ID,
		Name:        product.ProductName,
		Price:       product.Price,
		Description: product.Description,
		Category:    product.Category,
	}

	return formatter
}

func FormatProducts(products []entity.Product) []ProductFormatter {
	productsFormatter := []ProductFormatter{}

	for _, product := range products {
		productFormatter := FormatProduct(product)
		productsFormatter = append(productsFormatter, productFormatter)
	}

	return productsFormatter
}
