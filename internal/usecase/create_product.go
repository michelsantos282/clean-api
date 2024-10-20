package usecase

import "github.com/michelsantos282/clean-api/internal/entity"

type CreateProductInputDto struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateProductOutputDto struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewCreateProductUseCase(productRepository entity.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		ProductRepository: productRepository,
	}
}

func (u *CreateProductUseCase) Execute(input CreateProductInputDto) (*CreateProductOutputDto, error) {
	product := entity.NewProduct(input.Name, input.Price)
	err := u.ProductRepository.Create(product)
	if err != nil {
		return nil, err
	}
	return &CreateProductOutputDto{
		product.ID,
		product.Name,
		product.Price,
	}, nil
}
