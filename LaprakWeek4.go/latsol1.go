package main

import "fmt"

func main() {
	var total_belanja_awal, diskon, besar_diskon, harga_akhir int
	fmt.Scan(&total_belanja_awal)
	fmt.Scan(&diskon)
	besar_diskon = total_belanja_awal * diskon / 100
	harga_akhir = total_belanja_awal - besar_diskon
	fmt.Println(harga_akhir)
}
