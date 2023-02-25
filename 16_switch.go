package main
import "fmt"

// mirip switch case
func main(){
  name := "test"

  switch name{
    case "test":
      fmt.Println("Helo test")
      fmt.Println("Helo test")
    case "juned":
      fmt.Println("Helo juned")
    default:
      fmt.Println("Kenalan yok")
      fmt.Println("Helo test")
  }

  // short statement (local variable) di switch
  //switch length := len(name); length > 5{
  //  case true:
  //    fmt.Println("nama terlalu panjang")
  //  case false:
  //    fmt.Println("nama benar")
  //}

  // switch tanpa kondisi, jd seperti kek if expression
  // kondisi langsung msk ke case
  length := len(name)
  switch {
    case length > 10:
      fmt.Println("nama kepanjangan")
    case length > 5:
      fmt.Println("nama agak kepanjangan")
    default:
      fmt.Println("nama benar")
  }
}