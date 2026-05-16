package main
import "fmt"

func registerUser(name string, blacklist func(string) bool){
	if blacklist(name){
		fmt.Println("You are blocked", name)
	}else {
		fmt.Println("Welcome", name)
	}
}

func main(){
	// bentuk 1
	blacklist := func(name string) bool{
		if name == "joko" {
			return true
		}else if name == "budi" {
			return true
		}else {
			return false
		}
	}
	registerUser("joko", blacklist)

	// bentuk 2
	registerUser("tito", func(name string) bool {
		if name == "joko" {
			return true
		}else if name == "budi" {
			return true
		}else {
			return false
		}
	})
}