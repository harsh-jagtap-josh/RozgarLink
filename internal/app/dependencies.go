package app

import "database/sql"

type Dependencies struct {
	ProductService product.Service
}

func NewServices(db *sql.DB) Dependencies {
	productRepo := repository.NewProductRepo(db)

	//initialize service dependencies
	productService := product.NewService(productRepo)

	return Dependencies{
		// OrderService:   orderService,
		ProductService: productService,
	}
}
