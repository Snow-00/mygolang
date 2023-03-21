/**
secara default semua var di go dipassing by value, bkn reference
artinya jika kita mengirim var ke dlm func, method, var lain, itu 
sebenarnya mengirimkan duplikasi value nya

pointer adl kemampuan membuat reference ke lokasi data di memory
yg sama tanpa duplikasi (pass by reference)
*/

package main
import "fmt"

type Address struct {
  City, Province, Country string
}

func main() {
  // ini contoh pass by value
  //address1 := Address{"subang", "jabar", "indonesia"}
  //address2 := address1
  
  // ini contoh pass by reference
  address1 := Address{"subang", "jabar", "indonesia"}
  address2 := &address1
  var address3 *Address = &address1

  address2.City = "bandung"
  
  // meng-assign var address2 ke pointer baru (reference baru)
  //address2 = &Address{"malang", "jatim", "indonesia"}
  
  // memindahkan smua var yg terkonek ke pointer 1 ke pointer 2
  *address2 = Address{"malang", "jatim", "indonesia"}

  fmt.Println(address1)
  fmt.Println(address2)
  fmt.Println(address3)
  
  // membuat pointer baru yg isinya kosong
  address4:= new(Address)
  fmt.Println(address4)
  
  address4.City = "jakarta"
  fmt.Println(address4)
}