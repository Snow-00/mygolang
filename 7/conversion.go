package main
import "fmt"

func main() {
  var nilai32 int32 = 12324
  var nilai64 int64 = int64(nilai32)
  
  // if convert from big to small int type value
  // need to consider the size of the real value
  var nilai16 int16 = int16(nilai32)
  var nilai8 int8 = int8(nilai32)

  fmt.Println(nilai32)
  fmt.Println(nilai64)
  fmt.Println(nilai16)
  fmt.Println(nilai8) // int overflow

  var name = "zaki"
  var e = name[0] // this is byte (uint8)
  var estring = string(e)

  fmt.Println(name)
  fmt.Println(e)
  fmt.Println(estring)  
}