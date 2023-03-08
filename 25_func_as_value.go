package main
import "fmt"

// mirip kek OOP tp ini bnt nya msh func
// untungnya bs mskin func ke func secara lgsg
func getGoodBye(name string)string {
  return "goodbye " + name
}

func main() {
  sayGoodBye := getGoodBye

  result := sayGoodBye("test")
  fmt.Println(result)
  fmt.Println(getGoodBye("test"))
}