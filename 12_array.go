package main
import "fmt"

func main(){
  // manual array
  var names [3]string
  names[0] = "Joko"
  names[1] = "Linda"
  names[2] = "budi"

  fmt.Println(names[0])
  fmt.Println(names[1])
  fmt.Println(names[2])
  fmt.Println(names)

  // direct array
  var values = [3]int{
    90,
    80,
    14,  // must hv comma at the end of data
  }
  fmt.Println(values)

  // func array
  fmt.Println(len(names))

  var test [10]string
  fmt.Println(len(test))  
}