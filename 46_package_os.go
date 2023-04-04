/**
package2 bw an golang
Args ini untuk ambil args dr hasil run di terminal kt
format : go run file.go args
*/

package main
import (
  "fmt"
  "os"
)

func main() {
  args := os.Args
  fmt.Println("argument: ")
  fmt.Println(args)

  fmt.Println(args[1])
  fmt.Println(args[2])
  fmt.Println(args[3])

  name, err := os.Hostname()
  if err == nil {
    fmt.Println("hostname:", name)
  } else {
    fmt.Println("error:", err.Error())
  }

  username := os.Getenv("APP_USERNAME")
  password := os.Getenv("APP_PASSWORD")

  fmt.Println(username)
  fmt.Println(password)
}