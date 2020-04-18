package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi/Library"
	"restapi/config"
	"restapi/structs"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func GetOrder(w http.ResponseWriter, r *http.Request) {
	var barang structs.Barang
	var order structs.Order
	var user structs.Users
	var arr_order structs.Order
	var response structs.ResponseOrder

	db := config.Connect()
	defer db.Close()

	result := Library.CekAuth(w, r)
	if result {
		id := r.FormValue("id_barang")
		id_user := r.FormValue("id_user")
		jumlah_barang := r.FormValue("jumlah_barang")
		tgl_beli := Library.TimeStamp()

		if id == "" && id_user == "" && jumlah_barang == "" {
			Library.ErrorResponse(w, "Incompleted Parameter", "Lengkapi data terlebih dahulu", 2)
		} else {
			//Select Users
			rows, _ := db.Query("Select id,first_name,saldo from users where id = ?",
				id_user,
			)

			for rows.Next() {
				rows.Scan(&user.Id, &user.FirstName, &user.Saldo)
			}

			//Select Barang
			getBarang, _ := db.Query("Select id,nama_barang,harga_jual,stock from barang where id = ?",
				id,
			)

			for getBarang.Next() {
				getBarang.Scan(&barang.Id, &barang.NamaBarang, &barang.HargaJual, &barang.Stok)
			}

			jml_barang, _ := strconv.Atoi(jumlah_barang)
			hargaJual := &barang.HargaJual
			harga := *hargaJual * *&jml_barang
			status := "Sukses"
			generateId := Library.Inv(id_user, id)

			totalSaldo := user.Saldo - harga
			if totalSaldo < 0 {
				Library.ErrorResponse(w, "Saldo Habis", "Saldo Anda tidak cukup untuk melakukan transaksi", 6)
			} else {
				//Update Saldo
				_, _ = db.Exec("UPDATE users SET saldo = ? WHERE id = ?",
					totalSaldo,
					id_user,
				)
				//Order Barang
				_, err := db.Exec("INSERT INTO pembelian (id,barang,pembeli,jumlah,harga,tgl_beli,status) VALUES (?,?,?,?,?,?,?)",
					generateId,
					id,
					id_user,
					jumlah_barang,
					harga,
					tgl_beli,
					status,
				)
				if err != nil {
					log.Print(err)
				}
				stok := *&barang.Stok - jml_barang

				//update data barang
				_, errResp := db.Exec("UPDATE barang SET stock = ? WHERE id = ?",
					stok, id,
				)
				if errResp != nil {
					log.Print(err)
				}

				order.Id = generateId
				order.NamaBarang = *&barang.NamaBarang
				order.NamaPembeli = *&user.FirstName
				order.JumlahBarang = jumlah_barang
				order.Harga = harga
				order.TglBeli = tgl_beli
				order.Status = status

				arr_order = order

				response.ErrNumber = 0
				response.Status = "SUCCESS"
				response.Message = "Order Success"
				response.Data = arr_order
				response.RespTime = Library.TimeStamp()
				// log.Print(response)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			}
		}
	} else {
		Library.ErrorResponse(w, "Invalid Header", "Header yang anda masukkan tidak sesuai", 4)
	}
}

func RiwayatTransaksi(w http.ResponseWriter, r *http.Request) {
	var order structs.Order
	var arr_order []structs.Order
	var response structs.ResponseRiwayatOrder

	db := config.Connect()
	defer db.Close()

	result := Library.CekAuth(w, r)
	if result {
		id_user := r.FormValue("id_user")
		if id_user == "" {
			Library.ErrorResponse(w, "Incompleted Parameter", "Lengkapi data terlebih dahulu", 2)
		} else {
			rows, err := db.Query("Select id,nama_barang,first_name,jumlah,harga,tgl_beli,status from v_pembelian where id_user = ?",
				id_user,
			)
			for rows.Next() {
				if err := rows.Scan(&order.Id, &order.NamaBarang, &order.NamaPembeli, &order.JumlahBarang, &order.Harga, &order.TglBeli, &order.Status); err != nil {
					log.Fatal(err.Error())

				} else {
					arr_order = append(arr_order, order)
				}
			}
			if err != nil {
				log.Print(err)
			}

			response.ErrNumber = 0
			response.Status = "SUCCESS"
			response.Message = "Riwayat Transaksi"
			response.Data = arr_order
			response.RespTime = Library.TimeStamp()
			// log.Print(response)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	} else {
		Library.ErrorResponse(w, "Invalid Header", "Header yang anda masukkan tidak sesuai", 4)
	}
}
