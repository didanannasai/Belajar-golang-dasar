package main
import "fmt"

type Man struct {
	name string
}

func (man *Man) Married() {
	man.name = "Mr. " + man.name
}

func main() {
	tono := Man{"tono"}
	tono.Married()

	fmt.Println(tono.name)
}