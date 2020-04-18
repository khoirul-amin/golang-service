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

func GetProduk(w http.ResponseWriter, r *http.Request) {
	var produk structs.Produk
	var arr_produk []structs.Produk
	var response structs.ResponseProduk
	db := config.Connect()
	defer db.Close()

	result := Library.CekAuth(w, r)
	if result {
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
		response.Message = "Daftar Produk"
		response.Data = arr_produk
		response.RespTime = Library.TimeStamp()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		Library.ErrorResponse(w, "Invalid Header", "Header yang anda masukkan tidak sesuai", 4)
	}
}

func GetBarangByProduk(w http.ResponseWriter, r *http.Request) {
	var barang structs.Barang
	var arr_barang []structs.Barang
	var response structs.ResponseBarang
	db := config.Connect()
	defer db.Close()

	idProduk := r.FormValue("id_produk")

	result := Library.CekAuth(w, r)
	if result {
		if idProduk == "" {
			Library.ErrorResponse(w, "Incompleted Parameter", "Lengkapi data terlebih dahulu", 2)
		} else {
			rows, err := db.Query("SELECT id,nama_produk,nama_barang,harga_jual,satuan,stock FROM `v_barang` WHERE jenis_barang = ?",
				idProduk,
			)
			if err != nil {
				log.Print(err)
			}

			for rows.Next() {
				if err := rows.Scan(&barang.Id, &barang.NamaProduk, &barang.NamaBarang, &barang.HargaJual, &barang.Satuan, &barang.Stok); err != nil {
					log.Fatal(err.Error())
				} else {
					arr_barang = append(arr_barang, barang)
				}
			}
			response.ErrNumber = 0
			response.Status = "SUCCESS"
			response.Message = "Daftar Barang"
			response.Data = arr_barang
			response.RespTime = Library.TimeStamp()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	} else {
		Library.ErrorResponse(w, "Invalid Header", "Header yang anda masukkan tidak sesuai", 4)
	}
}
