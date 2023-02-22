package main
import "fmt"

// map mirip dictionary
func main(){
  person := map[string]string{
    "name": "dani",
    "address": "jakarta",
  }
  person["title"] = "program"

  fmt.Println(person)
  fmt.Println(person["name"])
  fmt.Println(person["address"])

  // membuat map baru tanpa harus diisi datanya dulu
  book := make(map[string]string)
  book["title"] = "belajar"
  book["author"] = "joko"
  book["ups"] = "salah"
  
  fmt.Println(book)
  fmt.Println(len(book))

  // untuk hapus data di map
  delete(book, "ups")

  fmt.Println(book)
  fmt.Println(len(book))
}