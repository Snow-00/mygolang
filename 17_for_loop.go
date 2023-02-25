package main
import "fmt"

// di golang hanya ada for loop sejauh ini
func main(){
  //counter := 1

  //for counter <= 10{
  //  fmt.Println("Perulangan ke", counter)
  //  counter++
  //}

  for counter := 1; counter <= 10; counter++ {
    fmt.Println("Perulangan ke", counter)
  }

  // for untuk array / slice secara manual
  slice := []string{"eko", "juju", "budi"}

  //for i := 0; i < len(slice); i++ {
  //  fmt.Println(slice[i])
  //}

  // for range untuk mengambil semua isi dari array / slice / map
  //for i, value := range slice {
  for _, value := range slice {     // jika i tdk ingin dipakai di code
    //fmt.Println("index", i, "=", value)
    fmt.Println(value)
  }

  // for range untuk map
  person := make(map[string]string)
  person["name"] = "dedod"
  person["title"] = "programmer"

  for key, value := range person {
    fmt.Println(key, "=", value)
  }
}