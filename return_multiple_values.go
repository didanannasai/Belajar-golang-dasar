package main
import "fmt"

func total(a, b int)(string, int){
	return "hasilnya :", a * b
}

func main(){
	// cara 1 :
	result1, result2 := total(5, 4)
	fmt.Println(result1, result2)
	// cara 2 :
	fmt.Println(total(5, 4))
}