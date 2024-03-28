package services

import (
	"context"
	"fmt"
	"time"

	models "github.com/Billy278/assignment_project/modules/models/products"
	repository "github.com/Billy278/assignment_project/modules/repository/product"
)

type SrvProductImpl struct {
	ProductRepo repository.ProductRepo
}

func NewSrvProductImpl(productrepo repository.ProductRepo) SrvProduct {
	return &SrvProductImpl{
		ProductRepo: productrepo,
	}
}

func (srv *SrvProductImpl) SrvList(ctx context.Context) (resProduct []models.Product, err error) {
	fmt.Println("SrvList")
	resProduct, err = srv.ProductRepo.RepoList(ctx)
	if err != nil {
		return
	}
	return
}
func (srv *SrvProductImpl) SrvFindByid(ctx context.Context, id uint64) (resProduct models.Product, err error) {
	fmt.Println("SrvFindByid")
	resProduct, err = srv.ProductRepo.RepoFindByid(ctx, id)
	if err != nil {
		return
	}
	return
}

func (srv *SrvProductImpl) SrvCreate(ctx context.Context, productIn models.Product) (resProduct models.Product, err error) {
	fmt.Println("SrvCreate")

	tNow := time.Now()
	productIn.Created_At = &tNow
	productIn.Updated_At = &tNow
	resProduct, err = srv.ProductRepo.RepoCreate(ctx, productIn)
	if err != nil {
		return
	}
	return
}

func (srv *SrvProductImpl) SrvDelete(ctx context.Context, id uint64) (err error) {
	fmt.Println("SrvDelete")
	_, err = srv.ProductRepo.RepoFindByid(ctx, id)
	if err != nil {
		return
	}
	err = srv.ProductRepo.RepoDelete(ctx, id)
	if err != nil {
		return
	}
	return

}
