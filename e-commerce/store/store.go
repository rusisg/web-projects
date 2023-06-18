package store

import "github.com/rusisg/e-commerce/types"

type ProductStorer interface {
	Insert(*types.Product) error
	GetByID(string) (*types.Product, error)
}
