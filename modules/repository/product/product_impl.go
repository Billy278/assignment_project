package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	models "github.com/Billy278/assignment_project/modules/models/products"
)

type RepoProductImpl struct {
	DB *sql.DB
}

func NewRepoProductImpl(db *sql.DB) ProductRepo {
	return &RepoProductImpl{
		DB: db,
	}
}
func (repo *RepoProductImpl) RepoList(ctx context.Context) (resProduct []models.Product, err error) {
	fmt.Println("RepoProductList")
	sqlList := "SELECT id,name,stock,price,created_at,updated_at FROM products"
	row, err := repo.DB.QueryContext(ctx, sqlList)
	if err != nil {
		return
	}
	defer row.Close()
	product := models.Product{}
	for row.Next() {
		err = row.Scan(&product.Id, &product.Name, &product.Stock, &product.Price, &product.Created_At, &product.Updated_At)
		if err != nil {
			return
		}
		resProduct = append(resProduct, product)
	}
	return
}
func (repo *RepoProductImpl) RepoFindByid(ctx context.Context, id uint64) (resProduct models.Product, err error) {
	fmt.Println("RepoFindByid")
	sqlFind := "SELECT id,name,stock,price,created_at,updated_at FROM products WHERE id=$1"
	row, err := repo.DB.QueryContext(ctx, sqlFind, id)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&resProduct.Id, &resProduct.Name, &resProduct.Stock, &resProduct.Price, &resProduct.Created_At, &resProduct.Updated_At)
		if err != nil {
			return
		}
	} else {
		err = errors.New("NOT FOUND")
	}
	return
}
func (repo *RepoProductImpl) RepoCreate(ctx context.Context, productIn models.Product) (resProduct models.Product, err error) {
	fmt.Println("RepoCreate")
	sqlCreate := "INSERT INTO products(name,stock,price,created_at,updated_at) VALUES($1,$2,$3,$4,$5) "
	_, err = repo.DB.ExecContext(ctx, sqlCreate, productIn.Name, productIn.Stock, productIn.Price, productIn.Created_At, productIn.Updated_At)
	if err != nil {
		return
	}
	return

}

func (repo *RepoProductImpl) RepoDelete(ctx context.Context, id uint64) (err error) {
	fmt.Println("RepoDelete")
	sqlDelete := "DELETE FROM products WHERE id=$1"
	_, err = repo.DB.ExecContext(ctx, sqlDelete, id)
	if err != nil {
		return
	}
	return
}
func (repo *RepoProductImpl) RepoUpdateTx(ctx context.Context, db *sql.Tx, productIn models.Product) (resProduct models.Product, err error) {
	fmt.Println("RepoUpdateProductTX")
	sqlUpdate := "UPDATE products set name=$1,stock=$2,price=$3,updated_at=$4 WHERE id=$5"

	_, err = db.ExecContext(ctx, sqlUpdate, productIn.Name, productIn.Stock, productIn.Price, productIn.Updated_At, productIn.Id)
	if err != nil {
		return
	}
	return
}
func (repo *RepoProductImpl) RepoFindByidTx(ctx context.Context, db *sql.Tx, id uint64) (resProduct models.Product, err error) {
	fmt.Println("RepoFindByidProductTx")
	sqlFind := "SELECT id,name,stock,price,created_at,updated_at FROM products WHERE id=$1 FOR UPDATE"
	row, err := db.QueryContext(ctx, sqlFind, id)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&resProduct.Id, &resProduct.Name, &resProduct.Stock, &resProduct.Price, &resProduct.Created_At, &resProduct.Updated_At)
		if err != nil {
			return
		}
	} else {
		err = errors.New("NOT FOUND")
	}

	return
}
