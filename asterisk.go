package main
import "fmt"

type address struct {
	city string
	province string
	country string
}

func main() {
	address1 := address{"semarang", "center java", "Indonesia"}
	address2 := &address1
	address2.city = "solo"
	fmt.Println(address1)
	fmt.Println(address2)

	// tanda asterisk digunakan saat kita ingin mengubah semua variabel yang mengacu pada data tersebut
	/* karna variabel address2 pass by reference ke variabel address1
	   maka saat variabel address2 diberikan tanda asterisk nilai dari variabel address1 juga ikut berubah */
	*address2 = address{"Bantul", "DIY YOGYAKARTA", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}