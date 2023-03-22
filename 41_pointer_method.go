package main
import "fmt"

 /**
 saat membuat method di dlm struct biasa, data juga diduplikat kek func
 makanya bs makan memory bny kalo engga pake pointer
 disarankan gunakan pointer utk struct fun / struct method
 */

type Man struct {
  Name string
}

//func (man Man) Married() {
func (man *Man) Married() {
  man.Name = "mr. " + man.Name
  fmt.Println("di method", man.Name)
}

func main() {
  eko := Man{"eko"}
  eko.Married()
  fmt.Println(eko.Name)
}