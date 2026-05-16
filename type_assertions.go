package main
import "fmt"

func random() any {
	return "oke"
}

func main(){
	result := random()
	// bagian yang di dalam kurung harus sesuai dengan return dari fungsinya agar tidak error
	resultString := result.(string)
	fmt.Println(resultString)

	// bisa memberikan keyword type di dalam kurung yang akan digunakan untuk menangkap tipe data apapun setelah itu bisa dimasukkan ke dalam variabel
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}