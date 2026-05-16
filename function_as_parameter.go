package main
import "fmt"

/*
   filter func(string) string : filter artinya nama funtion di parameter tersebut, func artinya
   parameter tersebut adalah function, (string) artinya terdapat parameter di function filter bertipe data
   string, string di luar kurung artinya terdapat return bertipe data string
*/
func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string)string{
	if name == "Anjing" {
		return "..."
	}else if name == "anjing" {
		return "..."
	}else {
		return name
	}
}

func main(){
	sayHelloWithFilter("udin", spamFilter)

	// mengkombinasikan dengan function as value
	filteredName := spamFilter
	sayHelloWithFilter("anjing", filteredName)
}