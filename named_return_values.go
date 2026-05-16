package main
import "fmt"

func getFullName()(firstName, middleName, lastName string){
	firstName = "xi"
	middleName = "jin"
	lastName = "ping"
	
	return firstName, middleName, lastName
}

func main(){
	fmt.Println(getFullName())
}