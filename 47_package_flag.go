/**
package flag berisikan fungsionalitas utk memparsing command line argument
dipake buat bikin apps berbasis terminal
*/

package main
import (
  "fmt"
  "flag"
)

func main() {
  host := flag.String("host", "localhost", "put ur host")
  var user *string = flag.String("user", "root", "put ur database urser")
  password := flag.String("password", "test", "put ur database pass")
  number := flag.Int("number", 100, "put ur num")
  
  // ini penting, kalo ga ttp nampilin defaultnya
  flag.Parse()

  fmt.Println("host:", *host)
  fmt.Println("user:", *user)
  fmt.Println("pass:", *password)
  fmt.Println("number:", *number)
}