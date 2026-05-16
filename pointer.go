package main
import "fmt"

type address struct {
	city string
	province string
	country string
}

func main() {
	address1 := address{"semarang", "center java", "Indonesia"}
	// tambahkan tanda "&" artinya yang dipanaggil adalah alamatnya, jadi address2 = alamat dari address1
	var address2 *address = &address1

	address2.city = "solo"
	fmt.Println(address1)
	fmt.Println(address2)
}