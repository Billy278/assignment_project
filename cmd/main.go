package main

import (
	"log"

	"github.com/Billy278/assignment_project/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// go func() {
	// 	// Membuat time ticker untuk menecek status birthday user dengan interval setiap 24
	// 	ticker := time.NewTicker(24 * time.Hour)
	// 	defer ticker.Stop()

	// 	// Memanggil fungsi untuk mengakses endpoint setiap kali ticker berdenyut
	// 	for range ticker.C {
	// 		client := http.Client{
	// 			Timeout:   time.Second * 10,
	// 			Transport: http.DefaultTransport,
	// 		}
	// 		//
	// 		urlPromo := fmt.Sprintf("http://%v:%v/api/promo", os.Getenv("hosPromoServices"), os.Getenv("PortPromoServices"))
	// 		req, err := http.NewRequest(http.MethodPost, urlPromo, nil)
	// 		if err != nil {
	// 			err = errors.New("Fail request promo services")
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 		res, err := client.Do(req)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 		fmt.Println(res)

	// 	}
	// }()

	server.NewServer()
}
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
