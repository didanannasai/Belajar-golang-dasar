package main
import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHello(name string){
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main(){
	var data Customer
	data.Name = "dono"
	data.Address = "Surabaya"
	data.Age = 21
	fmt.Println(data)

	// sebelum functionnya tambahkan object dari struct yang telah dibuat
	data.sayHello("budi")
}