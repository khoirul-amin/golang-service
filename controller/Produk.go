package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi/Library"
	"restapi/config"
	"restapi/structs"

	_ "github.com/go-sql-driver/mysql"
)


func GetProduk(w http.ResponseWriter, r *http.Request){
	var produk structs.Produk
	var arr_produk []structs.Produk
	var response structs.ResponseProduk

	ua := r.Header.Get("User-Agent")
	db := config.Connect()
	defer db.Close()


	if ua == "123" {
		rows, err := db.Query("SELECT id,nama_produk,status FROM produk WHERE status = 'Active'")
		if err != nil {
			log.Print(err)
		}
		for rows.Next() {
			if err := rows.Scan(&produk.Id, &produk.Nama, &produk.Status); err != nil {
				log.Fatal(err.Error())

			} else {
				arr_produk = append(arr_produk, produk)
			}
		}

		response.ErrNumber = 0
		response.Status = "SUCCESS"
		response.Message = "Daftar User"
		response.Data = arr_produk
		response.RespTime = Library.TimeStamp()
	} else {
		response.ErrNumber = 1
		response.Status = "ERROR"
		response.Message = "Header Salah"
		// response.Data = arr_user
		response.RespTime = Library.TimeStamp()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}