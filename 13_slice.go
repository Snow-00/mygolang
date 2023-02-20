// slice bagian dari array ex: array[1:10]
// slice terhubung dengan array sehingga jika data di array brubah
// maka isi slice juga brubah
// slice terdiri dari 3 bagian = pointer, length, capacity

package main
import "fmt"

func main(){
  // jika tidak tau jumlah data di array bisa menggunakan ... dahulu
  var months = [...]string{
    "Januari",
    "Februari",
    "Maret",
    "april",
    "mei",
    "juni",
    "juli",
    "agustus",
    "september",
    "oktober",
    "november",
    "desember",
  }
  
  // dlm membuat slice, jika menambahkan data > capacity, mk akan terbentuk
  // array baru sehingga perubahan data di array br tdk akan mempengaruhi
  // array asli
  var slice1 = months[4:7]
  fmt.Println(slice1)
  fmt.Println(len(slice1))
  fmt.Println(cap(slice1))

  months[5] = "test"
  fmt.Println(slice1)
  
  slice1[0] = "ubah"
  fmt.Println(months)
  
  // capacity = dihitung dr data awal slice ke data akhir array
  //slice2 := months[10:]
  slice2 := months[2:4]
  fmt.Println(slice2)
  //fmt.Println(cap(slice2))

  slice3 := append(slice2, "tambah")
  fmt.Println(slice3)

  slice3[1] = "bukan desember"
  fmt.Println(slice3)
  fmt.Println(slice2)
  fmt.Println(months)

  newSlice := make([]string, 2, 5)

  newSlice[0] = "test"
  newSlice[1] = "coba"
  fmt.Println(newSlice)
  fmt.Println(len(newSlice))
  fmt.Println(cap(newSlice))
}