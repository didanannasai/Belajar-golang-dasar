package main
import "fmt"

func main(){
	// closure adalah kemampuan function yang bisa mengakses variabel di sekitarnya selama satu scope
	counter := 0

	increment := func(){
		/* variabel counter bisa diakses dan dirubah dalam function ini,
		   karna variabel counter berada dalam 1 scope */
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
}