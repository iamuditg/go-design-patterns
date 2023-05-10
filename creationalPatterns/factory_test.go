package creationalPatterns

import "testing"

func TestConcreteFactory_Create(t *testing.T) {
	// Test that ConcreteFactory creates ConcreteProduct
	factory := &ConcreteFactory{}
	product := factory.Create()

	if _, ok := product.(*ConcreteProduct); !ok {
		t.Errorf("Expected product to be of type *ConcreteProduct, but got %T instead", product)
	}

	if product.Use() != "Using ConcreteProduct" {
		t.Errorf("Expected product.Use() to return 'Using ConcreteProduct, but got %s instead", product.Use())
	}

}
