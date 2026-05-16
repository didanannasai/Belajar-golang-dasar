package main
import "fmt"

func main(){
	// break
	fmt.Println("Break :")
	for i := 1; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("perulangan ke :", i)
	}

	// continue
	fmt.Print("\n")
	fmt.Println("Continue :")
	for i := 1; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println("perulangan ke :", i)
	}
}