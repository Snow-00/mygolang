/**
di golang huruf besar kecil menentukan access package nya
jd kalo nama func / var diawali huruf besar bs diakses package lain, 
kalo kecil engga bs
ini merujuk ke code 43 yg sebelumnya ga bs run terus code nya
*/

package main
import (
  "mygolang/helper"
  "fmt"
)

func main() {
  helper.SayHello("eko")
  //helper.sayGoodbye("eko")  -> error
  fmt.Println(helper.Application)
  //fmt.Println(helper.version) -> error
}