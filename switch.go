package main
import "fmt"

func main(){
	name := "dono"

	switch name{
	case "joko":
		fmt.Println("hello joko")
	case "budi":
		fmt.Println("hello budi")
	case "dono":
		fmt.Println("hello dono")
	default:
		fmt.Println("hello, boleh kenalan?")
	}
}