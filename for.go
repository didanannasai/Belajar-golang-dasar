package main
import "fmt"

func main(){
	counter := 1

	// for while
	for counter <= 10{
		fmt.Println("perulangan ke :", counter)
		counter++
	}

	// for statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke :", counter)
	}

	// for range
	var nama = [5]string {
		"udin",
		"jamal",
		"joko",
		"dono",
		"lukman",
	}
	  // cara manual
	for i := 0; i < len(nama); i++ {
		fmt.Println(nama[i])
	}
	  // cara for range
	for index, value := range nama {
		fmt.Println("index ke", index, "=", value)
	}
}