package app

type Dependencies struct {
	OrderService   order.Service
	ProductService product.Service
}

func NewServices(db *storm.DB) Dependencies {
	//initialize repo dependencies
	orderRepo := repository.NewOrderRepo(db)
	orderItemsRepo := repository.NewOrderItemRepo(db)
	productRepo := repository.NewProductRepo(db)

	//initialize service dependencies
	productService := product.NewService(productRepo)
	orderService := order.NewService(orderRepo, orderItemsRepo, productService)

	return Dependencies{
		OrderService:   orderService,
		ProductService: productService,
	}
}
