package main
import "fmt"

type Address struct {
	city, province, country string
}

func changeAddressToIndonesia(address *Address){
	address.country = "Indonesia"
}

func main() {
	address := Address{"Surabaya", "Jawa Timur", ""}
	changeAddressToIndonesia(&address)

	fmt.Println(address)
}