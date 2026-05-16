package main
import "fmt"

func factorial(i int)int {
	if i == 1 {
		return 1
	}else {
		return i * factorial(i - 1)
	}
}

func main(){
	fmt.Println(factorial(10))
}