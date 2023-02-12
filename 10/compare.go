package main
import "fmt"

func main(){
  var name1 = "eko"
  var name2 = "budi"

  var result = name1 == name2
  fmt. Println(result)
  
  result = name1 > name2     // true bcos "e" is after "b"
  fmt. Println(result)

  var value1 = 100
  var value2 = 200

  fmt.Println(value1 > value2)
  fmt.Println(value1 < value2)
  fmt.Println(value1 == value2)
  fmt.Println(value1 != value2)
}