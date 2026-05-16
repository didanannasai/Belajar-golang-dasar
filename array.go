package main
import "fmt"

func main(){
	var names [3]string

	names[0] = "udin"
	names[1] = "dono"
	names[2] = "dodo"
	fmt.Println(names)

	// jika ingin memberikan nilai langsung ke setiap indexnya
	var values = [5]int{
		1,
		7,
		55,
		21,
		99,
	}

	fmt.Println(values[1])
	fmt.Println(len(values))
	values[2] = 2
	fmt.Println(values[2])
}