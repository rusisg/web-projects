package store

import (
	"context"

	"github.com/rusisg/e-commerce/types"
)

type ProductStorer interface {
	Insert(context.Context, *types.Product) error
	GetByID(context.Context, string) (*types.Product, error)
}
