package main

import (
	// tambahkan tanda "_" jika hanya ingin mengeksekusi function ini tanpa mengeksekusi function lainnya
	_"golang-dasar/internal"
	"golang-dasar/database"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}