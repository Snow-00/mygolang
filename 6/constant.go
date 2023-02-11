package main
import "fmt"

// constant usually meant to be used for multiple src
// so we it won't return an error even if we don't
// use the variable in the same srcfile

func main(){
  // const dont hv cheat txt like var :=
  // const hv multi cheat txt

  const first string = "testing"
  const last = "jud"
  const value = 100
  const(
    first_name = "danny"
    new_value = 193
  )

  fmt.Println(first)
  fmt.Println(value)
}