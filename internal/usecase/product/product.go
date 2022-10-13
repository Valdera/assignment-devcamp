package product

import (
	"context"
	"log"
	e "shop/internal/entity/product"
)

type productRepo interface {
	GetProductById(ctx context.Context, id int64) (e.Product, error)
	GetProductAll(ctx context.Context) ([]e.Product, error)
	UpdateProduct(ctx context.Context, id int64, product e.Product) (e.Product, error)
	DeleteProduct(ctx context.Context, id int64) error
	AddProduct(ctx context.Context, product e.Product) (e.Product, error)
}

type ProductUsecase struct {
	productRepo productRepo
}

func New(repo productRepo) *ProductUsecase {
	return &ProductUsecase{
		productRepo: repo,
	}
}

func (p *ProductUsecase) AddProduct(ctx context.Context, data e.Product) (result e.Product, err error) {
	err = data.Validate()
	if err != nil {
		log.Println("[ProductUsecase][AddProduct] bad request, err: ", err.Error())
		return
	}

	result, err = p.productRepo.AddProduct(ctx, data)
	if err != nil {
		log.Println("[ProductUsecase][AddProduct] problem in getting from storage, err: ", err.Error())
		return
	}

	return
}

func (p *ProductUsecase) GetProductById(ctx context.Context, id int64) (result e.Product, err error) {
	result, err = p.productRepo.GetProductById(ctx, id)
	if err != nil {
		log.Println("[ProductUsecase][GetProduct] problem getting storage data, err: ", err.Error())
		return
	}

	return
}

func (p *ProductUsecase) GetProductAll(ctx context.Context) (result []e.Product, err error) {
	result, err = p.productRepo.GetProductAll(ctx)
	if err != nil {
		log.Println("[ProductUsecase][GetProduct] problem getting storage data, err: ", err.Error())
		return
	}

	return
}

func (p *ProductUsecase) UpdateProduct(ctx context.Context, id int64, data e.Product) (result e.Product, err error) {
	result, err = p.productRepo.UpdateProduct(ctx, id, data)
	if err != nil {
		log.Println("[ProductUsecase][UpdateProduct] problem getting storage data, err: ", err.Error())
		return
	}

	return
}
func (p *ProductUsecase) DeleteProduct(ctx context.Context, id int64) (err error) {
	err = p.productRepo.DeleteProduct(ctx, id)
	if err != nil {
		log.Println("[ProductUsecase][GetProduct] problem getting storage data, err: ", err.Error())
		return
	}

	return
}
