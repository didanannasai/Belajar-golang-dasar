package main
import "fmt"

func getGoodBye(name string)string{
	return "Bye " + name
}

func main(){
	/* sebuah function bisa dimasukkan ke dalam variabel sebagai value
	   dengan cara tidak menambahkan kurung buka dan tutup setelah nama functionnya */
	var goodBye = getGoodBye
	fmt.Println(goodBye("siregar"))
}