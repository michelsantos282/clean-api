package usecase

import "github.com/michelsantos282/clean-api/internal/entity"

type ListProductsOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type ListProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductUseCase(productRepository entity.ProductRepository) *ListProductUseCase {
	return &ListProductUseCase{
		ProductRepository: productRepository,
	}
}

func (u *ListProductUseCase) Execute() ([]*ListProductsOutputDto, error) {
	products, err := u.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var productsOutput []*ListProductsOutputDto

	for _, product := range products {
		productsOutput = append(productsOutput, &ListProductsOutputDto{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}
	return productsOutput, nil
}
