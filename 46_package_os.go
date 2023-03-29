/**
package2 bw an golang
ini untuk ambil args dr hasil run di terminal kt
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
}