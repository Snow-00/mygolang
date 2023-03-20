package main
import "fmt"

/**
type assertion adl kemampuan utk merubah tipe data menjd tipe data yg kt ingin
fitur ini sering digunakan ketika bertemu data interface kosong
hrs hati2 ketika pake ini, krn jika salah ubah yg misal hrsnya bool
tp jd string nanti bakal panic
*/

func random() interface{} {
  return "100"
}

func main() {
  var result interface{} = random()
  var resultString string = result.(string)  // ini command type assertion
  fmt.Println(resultString)

  // agar tidak salah ubah, lakukan seperti ini
  switch value := result.(type) {
    case string:
      fmt.Println("value", value, "is string")
    case int:
      fmt.Println("value", value, "is int")
    default:
      fmt.Println("unknown type")
  }
}