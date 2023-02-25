package main
import "fmt"

func sayHello() {
  fmt.Println("helo")
}

func main() {
  sayHello()
  sayHello()

  for i := 0; i < 3; i++ {
    sayHello()
  }
}