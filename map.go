package main
import "fmt"

func main(){
	// cara mendeklarasikan sebuah map, string di dalam kurung adalah tipe data key, string di luar kurung tipe data value
	var data = make(map[string]string)

	// cara memberikan nilai ke dalam map
	data["nama"] = "udin"
	data["alamat"] = "jakarta"

	var data2 = map[string]string{
		"nama" : "budi",
		"alamat" : "bandung",
		"wrong" : "ups",
	}
	fmt.Println(data)
	fmt.Println(data2)

	// cara menghapus data di map
	delete(data2, "wrong")
	fmt.Println(data2)
}