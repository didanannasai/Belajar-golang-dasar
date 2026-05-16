package main
import "fmt"

// artinya bisa menjadi bagian dari type HasName jika memiliki method GetName()
type HasName interface {
	GetName() string
}

func sayHello(value HasName){
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	name string
}

func (person Person) GetName() string{
	return "saya " + person.name
}

type Hewan struct {
	name string
}

func (hewan Hewan) GetName() string{
	return hewan.name
}

func main(){
	person := Person{"budi"}
	animal := Hewan{"kucing"}
	sayHello(person)
	sayHello(animal)
}