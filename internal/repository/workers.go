package repository

import (
	"context"
)

type ProductStorer interface {
	GetProductByID(ctx context.Context, tx Transaction, productID int64) (Product, error)
	ListProducts(ctx context.Context, tx Transaction) ([]Product, error)
	UpdateProductQuantity(ctx context.Context, tx Transaction, productsQuantityMap map[int64]int64) error
}
