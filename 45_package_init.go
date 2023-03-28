/**
blank identifier adl ketika kt cmn pingin jalanin func init dr package
aj tanpa make func lain dr package itu, contoh engga mau jalanin GetDatabase
*/

package main
import (
  //"mygolang/database"
  //"fmt"
  
  _ "mygolang/database"   // kalo tanpa _ bakal error
)

func main() {
  //result := database.GetDatabase()
  //fmt.Println(result)
}