package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	models "github.com/Billy278/assignment_project/modules/models/promo"
	modelsToken "github.com/Billy278/assignment_project/modules/models/token"
	modelsUser "github.com/Billy278/assignment_project/modules/models/user"
	repository "github.com/Billy278/assignment_project/modules/repository/promo"
	"github.com/Billy278/assignment_project/pkg/crypto"
	"github.com/Billy278/assignment_project/pkg/responses"
	"github.com/patrickmn/go-cache"
)

type PromoSrvImpl struct {
	RepoPromo repository.PromoRepo
}

var C = cache.New(5*time.Minute, 10*time.Minute)

func NewPromoSrvImpl(repopromo repository.PromoRepo) PromoSrv {
	return &PromoSrvImpl{
		RepoPromo: repopromo,
	}
}

func (srv *PromoSrvImpl) GetToken(ctx context.Context, promo string) (resToken string, err error) {
	fmt.Println("Services GetToken")
	// find token in cache
	res, ok := C.Get(promo)
	if ok {
		resToken = res.(string)
		return
	}
	resToken, err = srv.RepoPromo.GetToken(ctx, promo)
	if err != nil {
		return
	}
	// Save value to cache ttl
	C.Set(promo, resToken, 5*time.Minute)

	return
}
func (srv *PromoSrvImpl) CreatedPromo(ctx context.Context) (err error) {
	fmt.Println("CreatePromo Services")
	resUser, err := srv.GetUsersIsBirthday(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resUser)

	resToken := []string{}
	var res string
	// do loop apabila data yg diterima lebih dari 1
	for i := 0; i < len(resUser); i++ {
		res, err = srv.GeneratePromoCode(ctx, resUser[i].Id, resUser[i].Name)
		if err != nil {
			return
		}
		resToken = append(resToken, res)
	}
	// created and save data promo
	var Promo string
	for i := 0; i < len(resUser); i++ {
		Promo = fmt.Sprintf("HAPPYBIRTHDAY%v", resUser[i].Name)
		err = srv.SendNotification(ctx, resUser[i].Name, Promo, resUser[i].Gmail)
		if err != nil {
			return
		}
	}
	//promoIn := []models.Promo{}
	promoTemp := models.Promo{}
	tNow := time.Now()
	for i := 0; i < len(resToken); i++ {
		promoTemp.Created_at = &tNow
		promoTemp.Kode_Promo = fmt.Sprintf("HAPPYBIRTHDAY%v", resUser[i].Name)
		promoTemp.Token = resToken[i]
		err = srv.RepoPromo.Created(ctx, promoTemp)
		if err != nil {
			return
		}
	}

	return

}

func (srv *PromoSrvImpl) GetUsersIsBirthday(ctx context.Context) (resUser []modelsUser.User, err error) {
	client := http.Client{
		Timeout:   time.Second * 10,
		Transport: http.DefaultTransport,
	}
	//get data user
	urlGetUsers := fmt.Sprintf("http://%v:%v/api/user", os.Getenv("hostUserServices"), os.Getenv("PortUserServices"))
	req, err := http.NewRequest(http.MethodGet, urlGetUsers, nil)
	if err != nil {
		err = errors.New("Fail get data user")

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
		err = errors.New("Fail to unmarshal  data user")
		return
	}
	if !resData.Success {
		err = errors.New("Fail to get  data user")
		return
	}

	jsonData, err := json.Marshal(resData.Data)
	if err != nil {
		err = errors.New("Fail to Marsal  data user to slice user")
		return
	}

	json.Unmarshal(jsonData, &resUser)
	return
}

func (srv *PromoSrvImpl) GeneratePromoCode(ctx context.Context, idUser uint64, name string) (tokenRes string, err error) {
	// generate token
	tNow := time.Now()
	// set token expired berlaku 1 hari setelah di buat
	defaultClaim := modelsToken.DefaultClaimPromo{
		Expired:   int(tNow.Add(24 * time.Hour).Unix()),
		NotBefore: int(time.Now().Unix()),
		IssuedAt:  int(time.Now().Unix()),
		Issuer:    fmt.Sprint(idUser),
		Audience:  "assgiment_project",
		JTI:       "jti",
		Type:      modelsToken.ACCESS_TOKEN,
	}

	accessTokenClaim := struct {
		modelsToken.DefaultClaimPromo
		modelsToken.AccessClaimPromo
	}{
		DefaultClaimPromo: defaultClaim,
		AccessClaimPromo: modelsToken.AccessClaimPromo{
			UserId: fmt.Sprint(idUser),
			Name:   name,
		},
	}

	res, err := crypto.CreatedJWT(accessTokenClaim)
	if err != nil {
		err = errors.New("failed to create token")
		return
	}

	tokenRes = res
	return

}

func (srv *PromoSrvImpl) SendNotification(ctx context.Context, name, promo, receiver string) (err error) {
	client := http.Client{
		Timeout:   time.Second * 10,
		Transport: http.DefaultTransport,
	}
	//Send to gmail services
	body := fmt.Sprintf(`{
				"name":"%v",
				"promo" :"%v",
				"receiver":"%v"
				}`, name, promo, receiver)
	BodyGmail := strings.NewReader(body)
	urlPostGmail := fmt.Sprintf("http://%v:%v/api/gmail", os.Getenv("hostGmailServices"), os.Getenv("PortGmailServices"))
	req, err := http.NewRequest(http.MethodPost, urlPostGmail, BodyGmail)
	if err != nil {
		err = errors.New("Fail send data gmail")
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
		err = errors.New("Fail to unmarshal  data gmail")
		return
	}

	fmt.Println(resData)
	if !resData.Success {
		err = errors.New("Fail to send  data gmail")
		return
	}

	return
}
