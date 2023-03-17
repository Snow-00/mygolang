package main
import "fmt"

/** 
nilai nil itu cuman ada di tipe data interface, func, map, slice,
pointer dan channel, slain tipe data ini pas deklarasi var kosong
kenanya default value (int -> 0), bkn nil
*/

func newMap(name string) map[string]string {
  if name == "" {
    return nil
  } else {
    return map[string]string {
      "name": name,
    }
  }
}
func main() {
  // bs jg engga ush dikasi = ...
  //var person map[string]string = nil

  person := newMap("eko")
  fmt.Println(person)
}