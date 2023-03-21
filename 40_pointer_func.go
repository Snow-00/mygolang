/**
saat membuat func, ketika masukin var ke func, secara default itu
pass by value, jd modif di func itu engga bakal ubah data aslinya

mirip kek global var_ di python konsepnya

fungsi dr penggunaan pointer di func ini, ketika struct nya besar
dan kt sering mskin ke func itu dapat mengurangi penggunaan memory
krn engga diduplikat datanya, hanya mengacu ke ref yg sama
*/

package main
import "fmt"

type Address struct {
  City, Province, Country string
}

/** 
ini data duplikatnya doang yg brubah isinya, alamat aslinya engga
krn yg kt mskin var nya doang, hrsnya pointernya yg kt mskin
*/
//func changeCountryToIndonesia(address Address) {
//  address.Country = "indonesia"
//}

func changeCountryToIndonesia(address *Address) {
  address.Country = "indonesia"
}

func main() {
  alamat := Address{
    City:     "subang",
    Province: "jabar",
    Country:  "",
  }
  
  //changeCountryToIndonesia(alamat)
  alamatPointer := &alamat
  changeCountryToIndonesia(alamatPointer)  // bs gini / lgsg mskin &alamat
  fmt.Println(alamat)
}