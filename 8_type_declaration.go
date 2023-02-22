package main
import "fmt"

// type is for giving aliases or name for the var type, nothing more
func main(){
  type NoKTP string
  type Married bool

  var no_ktp NoKTP = "123561234345"
  var married_stat Married = true
  fmt.Println(no_ktp)
  fmt.Println(married_stat)
}