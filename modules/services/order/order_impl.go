package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	models "github.com/Billy278/assignment_project/modules/models/orders"
	repository "github.com/Billy278/assignment_project/modules/repository/order"
	repoProduct "github.com/Billy278/assignment_project/modules/repository/product"
	"github.com/Billy278/assignment_project/pkg/responses"
)

type OrderSrvImpl struct {
	DB          *sql.DB
	OrderRepo   repository.OrderRepo
	ProductRepo repoProduct.ProductRepo
}

func NewOrderSrvImpl(db *sql.DB, orderrepo repository.OrderRepo, productrepo repoProduct.ProductRepo) OrderSrv {
	return &OrderSrvImpl{
		DB:          db,
		OrderRepo:   orderrepo,
		ProductRepo: productrepo,
	}
}
func (srv *OrderSrvImpl) SrvList(ctx context.Context) (resOrder []models.Order, err error) {
	fmt.Println("SrvList")
	resOrder, err = srv.OrderRepo.RepoList(ctx, srv.DB)
	if err != nil {
		return
	}
	return
}
func (srv *OrderSrvImpl) SrvGetTokenWithPromo(ctx context.Context, kode_promo string) (token string, err error) {
	client := http.Client{
		Timeout:   time.Second * 10,
		Transport: http.DefaultTransport,
	}
	//get data user
	urlGetToken := fmt.Sprintf("http://%v:%v/api/promo?kode_promo=%v", os.Getenv("hostUserServices"), os.Getenv("PortUserServices"), kode_promo)
	req, err := http.NewRequest(http.MethodGet, urlGetToken, nil)
	if err != nil {
		err = errors.New("Fail get promo token")

		return
	}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	resData := responses.Response{}
	err = json.Unmarshal(data, &resData)
	if err != nil {
		err = errors.New("Fail to unmarshal  data promo")
		return
	}
	if !resData.Success {
		err = errors.New("Fail to get  data promo")
		return
	}

	jsonData, err := json.Marshal(resData.Data)
	if err != nil {
		err = errors.New("Fail to Marsal token")
		return
	}
	json.Unmarshal(jsonData, &token)
	return
}
func (srv *OrderSrvImpl) SrvCreate(ctx context.Context, orderIn models.Order, isgetPromo bool) (resOrder models.Order, err error) {
	fmt.Println("Srv Create Product")
	tx, err := srv.DB.Begin()
	if err != nil {
		return
	}
	// get data Product
	dataProduct, err := srv.ProductRepo.RepoFindByidTx(ctx, tx, orderIn.ProductId)
	if err != nil {
		return
	}

	// cek apakah stock product >= qty
	if dataProduct.Stock < orderIn.Qty {
		msg := fmt.Sprintf("Stock Product with id=%v tidak Mencukupi", dataProduct.Id)
		err = errors.New(msg)
		tx.Rollback()
		return
	}

	dataProduct.Stock = dataProduct.Stock - orderIn.Qty
	var result, discon float64
	if isgetPromo {
		// diskon 20 %
		result = dataProduct.Price * float64(orderIn.Qty)
		discon = result * 0.2
		result = result - discon
		orderIn.Promo = discon

	} else {
		result = dataProduct.Price * float64(orderIn.Qty)
	}

	// cek paid ? >= result
	if orderIn.TotalPaid < result {
		tx.Rollback()
		err = errors.New("Saldo Anda kurang")
		return
	}

	// update data product
	tNow := time.Now()
	dataProduct.Updated_At = &tNow
	_, err = srv.ProductRepo.RepoUpdateTx(ctx, tx, dataProduct)
	if err != nil {
		tx.Rollback()
		return
	}

	// create data order
	orderIn.Created_At = &tNow
	orderIn.TotalPrize = result
	orderIn.TotalReturn = orderIn.TotalPaid - result
	resOrder, err = srv.OrderRepo.RepoCreate(ctx, tx, orderIn)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()

	return
}
