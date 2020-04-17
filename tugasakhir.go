package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

var baseURL = "http://localhost:8000"

type Barang struct {
	Id          int
	NamaBarang  string
	HargaBarang int
}

func GetBarang() {
	var client = &http.Client{}
	var barang []Barang

	request, err := http.NewRequest("POST", baseURL, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&barang)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range barang {
		fmt.Printf("ID: %d\t Name: %s\t Harga: %d\n", each.Id, each.NamaBarang, each.HargaBarang)
	}
}

func GetBarangId(ID string) {
	var err error
	var client = &http.Client{}
	var barang []Barang
	var param = url.Values{}
	param.Set("id", ID)
	var payload = bytes.NewBufferString(param.Encode())
	request, err := http.NewRequest("POST", baseURL+"/cari", payload)
	if err != nil {
		fmt.Println(err.Error())

	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())

	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&barang)
	if err != nil {
		fmt.Println(err.Error())

	}

	for _, each := range barang {
		fmt.Printf("ID: %d\t Name: %s\t Harga: %d\n", each.Id, each.NamaBarang, each.HargaBarang)
	}
}

func GetTambahBarang(NamaBarang string, HargaBarang string) {
	var client = &http.Client{}
	var param = url.Values{}
	param.Set("nama_barang", NamaBarang)
	param.Set("harga_barang", HargaBarang)
	var payload = bytes.NewBufferString(param.Encode())
	request, err := http.NewRequest("POST", baseURL+"/tambah", payload)
	if err != nil {
		fmt.Println(err.Error())

	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())

	}
	defer response.Body.Close()

	fmt.Println("Data Barhasil di tambahkan")
}

func GetDeleteBarang(id string) {
	var client = &http.Client{}
	var param = url.Values{}
	param.Set("id", id)
	var payload = bytes.NewBufferString(param.Encode())
	request, err := http.NewRequest("POST", baseURL+"/delete", payload)
	if err != nil {
		fmt.Println(err.Error())

	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())

	}
	defer response.Body.Close()

	fmt.Println("Data Barhasil di hapus")
}

func GetUpdateBarang(id string, NamaBarang string, HargaBarang string) {
	var client = &http.Client{}
	var param = url.Values{}
	param.Set("id", id)
	param.Set("nama_barang", NamaBarang)
	param.Set("harga_barang", HargaBarang)
	var payload = bytes.NewBufferString(param.Encode())
	request, err := http.NewRequest("POST", baseURL+"/update", payload)
	if err != nil {
		fmt.Println(err.Error())

	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())

	}
	defer response.Body.Close()

	fmt.Println("Data Barhasil di update")
}

func main() {
	GetBarang()
	GetUpdateBarang("2", "Meja", "20")
	GetDeleteBarang("2")
	GetTambahBarang("Meja", "20")
	GetBarangId("2")

}
