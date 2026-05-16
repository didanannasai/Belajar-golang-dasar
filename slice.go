package main
import "fmt"

func main(){
	/* artinya membuat slice dengan tipe data integer dan jumlah data di dalamnya sementara ada 3(bisa ditambah lagi),
	   lalu capacitynya ada 5, capacity adalah kapasitas dalam satu slice jika tidak muat maka akan dibuat slice baru */
	var number1 = make([]int, 3, 5)
	// bentuk lain :
	var number2 []int
	// cara menambahkan data, 55 dan 12 akan masuk ke slice
	number2 = append(number2, 55, 12)
	fmt.Println(number1)
	fmt.Println(number2)

	var days = [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	//[4:6], 4 adalah index pertama yang akan dimasukkan ke slice hingga sebelum index 6
	var slice1 = days[4:6]
	fmt.Println(slice1)
	//[4:], 4 adalah index pertama yang akan dimasukkan ke dalam slice hingga index terakhir array
	var slice2 = days[4:]
	fmt.Println(slice2)
	//[:3], index 0 hingga index sebelum 3 akan dimasukkan ke dalam slice
	var slice3 = days[:3]
	fmt.Println(slice3)
	//[:], index pertama hingga index terakhir akan dimasukkan ke dalam array
	var slice4 = days[:]
	fmt.Println(slice4)

	// len() akan mendapatkan nilai panjang slicenya
	fmt.Println(len(slice1))
	var daySlice1 = days[4:7]
	daySlice1[0] = "jumat baru"
	daySlice1[1] = "sabtu baru"
	fmt.Println(days)
	var daySlice2 = append(daySlice1, "hari baru")
	daySlice2[0] = "perubahan"
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice = days[:]
	fmt.Println("===")
	fmt.Println(newSlice)
	var newSlice2 = make([]string, len(newSlice), cap(newSlice))
	fmt.Println(newSlice2)
	copy(newSlice2, newSlice)
	fmt.Println(newSlice2)
}