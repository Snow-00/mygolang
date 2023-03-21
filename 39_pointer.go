/**
secara default semua var di go dipassing by value, bkn reference
artinya jika kita mengirim var ke dlm func, method, var lain, itu 
sebenarnya mengirimkan duplikasi value nya
*/

package main
import "fmt"

type Address struct {
  City, Province, Country string
}

func main() {
  address1 := Address{"subang", "jabar", "indonesia"}
  address2 := address1
  
  address2.City = "bandung"
  
  fmt.Println(address1)
  fmt.Println(address2)
}