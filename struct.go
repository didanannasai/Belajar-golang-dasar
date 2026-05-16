package main
import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func main(){
	var data Customer
	data.Name = "dono"
	data.Address = "Surabaya"
	data.Age = 21
	fmt.Println(data)

	// cara lain
	joko := Customer{
		Name: "joko",
		Address: "Sidoarjo",
		Age: 25,
	}
	fmt.Println(joko)

	// cara lain
	budi := Customer{"Budi", "Sidoarjo", 19}
	fmt.Println(budi)
}