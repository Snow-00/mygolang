package main
import "fmt"

/**
interface kosong adl kontrak yg isinya kosong
jd semua tipe data golang mengikuti kontrak interface kosong
enaknya interface kosong, 1 var bisa nrima semua tipe data
*/

func ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}
func main() {
	// ini salah, krn interface kosong tdk bs dimskin ke suatu tipe data 
	//var data int = ups(1)

	var data interface{} = ups(3)
	fmt.Println(data)
}