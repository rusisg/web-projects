package types

import "fmt"

const minProductNamelen = 3

type Product struct {
	SKU  string `bson:"sku" json:"sku"`
	Name string `bson:"name" json:"name"`
	Slug string `bson:"slug" json:"slug"`
}

type CreateProductRequest struct {
	SKU  string `json:"sku"`
	Name string `json:"name"`
}

func NewProductFromRequest(req *CreateProductRequest) (*Product, error) {
	if err := ValidateCreateProductRequest(req); err != nil {
		return nil, err
	}
	return &Product{
		SKU:  req.SKU,
		Name: req.Name,
	}, nil
}

func ValidateCreateProductRequest(req *CreateProductRequest) error {
	if len(req.Name) < minProductNamelen {
		fmt.Errorf("the name of the product is too short")
	}
	return nil
}
