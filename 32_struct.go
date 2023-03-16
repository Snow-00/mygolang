package main
import "fmt"

type Customer struct {
  //Name string
  //Address string
  //Age int

  Name, Address string
  Age           int
}

func main() {
  var eko Customer
  eko.Name = "eko"
  eko.Address = "indonesia"
  eko.Age = 29

  fmt.Println(eko)
  fmt.Println(eko.Name)
  fmt.Println(eko.Address)
  fmt.Println(eko.Age)

  joko := Customer {
    Name: "joko",
    Address: "jaksel",
    Age: 35,
  }
  fmt.Println(joko)
  
  /** 
  cara ke3 ini rawan error krn sangat bergantung dg posisi dan jumlah data
  makanya cara ini engga rekomen
  */
  budi := Customer{"budi", "Bekasi", 30}
  fmt.Println(budi)
}