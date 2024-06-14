package product

import (
	"errors"
	"go-technical-test-synapsis/src/entity"
	"go-technical-test-synapsis/src/helper"
	"testing"
)

type MockProductRepository struct {
	data map[int]entity.Product
}

func (m *MockProductRepository) FindAll(offset, limit int) ([]entity.Product, error) {
	var products []entity.Product
	for _, product := range m.data {
		products = append(products, product)
	}
	return products, nil
}

func (m *MockProductRepository) FindByCategory(category string) ([]entity.Product, error) {
	var products []entity.Product
	for _, product := range m.data {
		products = append(products, product)
	}
	return products, nil
}

func (m *MockProductRepository) Save(product entity.Product) (entity.Product, error) {
	if m.data == nil {
		m.data = make(map[int]entity.Product)
	}
	m.data[product.ID] = product
	return product, nil
}

func (m *MockProductRepository) FindByID(productID int) (entity.Product, error) {
	if product, ok := m.data[productID]; ok {
		return product, nil
	}
	return entity.Product{}, errors.New("product not found")
}

func (m *MockProductRepository) DeleteProductByID(productID int) (entity.Product, error) {
	product, ok := m.data[productID]
	if !ok {
		return entity.Product{}, errors.New("product not found")
	}
	delete(m.data, productID)
	return product, nil
}

func TestFindProducts(t *testing.T) {
	mockRepo := &MockProductRepository{}

	service := NewService(mockRepo)

	testProducts := []entity.Product{
		{ID: 1, ProductName: "Product 1", Price: 10000, Category: "FOOD", CreatedAt: helper.TimeNow()},
		{ID: 2, ProductName: "Product 2", Price: 20000, Category: "MOTORCYCLE", CreatedAt: helper.TimeNow()},
		{ID: 3, ProductName: "Product 3", Price: 30000, Category: "SHIRT", CreatedAt: helper.TimeNow()},
	}
	mockRepo.data = make(map[int]entity.Product)
	for _, product := range testProducts {
		mockRepo.data[product.ID] = product
	}

	// Test FindProducts method
	products, err := service.FindProducts(0, 10)
	if err != nil {
		t.Errorf("Error getting products: %v", err)
	}
	if len(products) != len(testProducts) {
		t.Errorf("Expected %d products, got %d", len(testProducts), len(products))
	}
}

func TestStoreProduct(t *testing.T) {
	mockRepo := &MockProductRepository{}

	service := NewService(mockRepo)

	request := ProductRequest{
		Name:        "Test Product",
		Price:       50000,
		Description: "Test Description",
		Category:    "MOTORCYCLE",
	}

	// Test StoreProduct method
	product, err := service.StoreProduct(request)
	if err != nil {
		t.Errorf("Error storing product: %v", err)
	}

	// Check if product is saved in the repository
	_, err = mockRepo.FindByID(product.ID)
	if err != nil {
		t.Errorf("Stored product not found in the repository")
	}
}

func TestFindProductByID(t *testing.T) {
	mockRepo := &MockProductRepository{}

	service := NewService(mockRepo)

	testProduct := entity.Product{
		ID:          1,
		ProductName: "Test Product",
		Price:       50000,
		Description: "Test Description",
		Category:    "MOTORCYCLE",
		CreatedAt:   helper.TimeNow(),
	}
	mockRepo.data = map[int]entity.Product{
		1: testProduct,
	}

	product, err := service.FindProductByID(ProductIDRequest{ID: 1})
	if err != nil {
		t.Errorf("Error finding product by ID: %v", err)
	}
	if product.ID != testProduct.ID {
		t.Errorf("Expected product ID %s, got %s", testProduct.ID, product.ID)
	}
}

func TestUpdateProduct(t *testing.T) {
	mockRepo := &MockProductRepository{}

	service := NewService(mockRepo)

	testProduct := entity.Product{
		ID:          1,
		ProductName: "Test Product",
		Price:       50000,
		Description: "Test Description",
		Category:    "MOTORCYCLE",
		CreatedAt:   helper.TimeNow(),
	}
	mockRepo.data = map[int]entity.Product{
		1: testProduct,
	}

	// Test UpdateProduct method with correct seller ID
	updateRequest := ProductRequest{
		Name:        "Updated Product",
		Price:       60000,
		Description: "Updated Description",
	}
	updatedProduct, err := service.UpdateProduct(ProductIDRequest{ID: 1}, updateRequest)
	if err != nil {
		t.Errorf("Error updating product: %v", err)
	}
	if updatedProduct.ProductName != updateRequest.Name {
		t.Errorf("Expected updated product name %s, got %s", updateRequest.Name, updatedProduct.ProductName)
	}

	// Test UpdateProduct method with incorrect seller ID
	_, err = service.UpdateProduct(ProductIDRequest{ID: 1}, updateRequest)
	if err == nil {
		t.Errorf("Expected error for unauthorized seller update")
	}
}

func TestDeleteProductByID(t *testing.T) {
	mockRepo := &MockProductRepository{}

	service := NewService(mockRepo)

	testProduct := entity.Product{
		ID:          1,
		ProductName: "Test Product",
		Price:       50000,
		Description: "Test Description",
		Category:    "MOTORCYCLE",
		CreatedAt:   helper.TimeNow(),
	}
	mockRepo.data = map[int]entity.Product{
		1: testProduct,
	}

	deletedProduct, err := service.DeleteProductByID(ProductIDRequest{ID: 1})
	if err != nil {
		t.Errorf("Error deleting product by ID: %v", err)
	}
	if deletedProduct.ID != testProduct.ID {
		t.Errorf("Expected deleted product ID %s, got %s", testProduct.ID, deletedProduct.ID)
	}
}
