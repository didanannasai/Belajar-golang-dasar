package main
import "fmt"

/* interface kosong atau any, artinya memiliki kontrak ke semua tipe data,
  jadi dalam kasus di bawah maka return bisa berupa semua tipe data */
func ups() any {
	//return 20 --> bisa
	// return true --> bisa
	return "halo"  // --> bisa
}

func main(){
	fmt.Println(ups())
}