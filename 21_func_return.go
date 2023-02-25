package main
import "fmt"

func getHello(name string) string {
  result := "hello"
  if name == "" {
    return "hello bro"
  } else {
    return result + " " + name
  }
}

func main() {
  result := getHello("budik")
  fmt.Println(result)

  result = getHello("")
  fmt.Println(result)
}