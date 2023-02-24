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
}