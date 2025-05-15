package product

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	SKU         string  `json:"sku" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
}

type CreateProductResponse struct {
	Name string `json:"name"`
	SKU  string `json:"sku"`
}
