package dto

type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UpdateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductOutput struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
}


type ProductOutputId struct {
	ID string `json:"id"`
}
