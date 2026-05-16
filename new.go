package main
import "fmt"

type address struct {
	city, province, country string
}

func main(){
	/* new() adalah function yang ada di golang digunakan untuk menggantikan
	  tanda "&" jika ingin mengembalikan alamat memorinya */
	var alamat1 *address = new(address)
	var alamat2 *address = alamat1

	alamat2.country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}