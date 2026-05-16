package main
import "fmt"

func main(){
	name := "budi"

	if name == "eko" {
		fmt.Println("hello eko")
	}else if name == "budi" {
		fmt.Println("hello budi")
	}else if name == "joko" {
		fmt.Println("hello joko")
	}else {
		fmt.Println("hello, boleh kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("nama anda kepanjangan")
	}else {
		fmt.Println("nama anda sudah benar")
	}
}