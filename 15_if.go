package main
import "fmt"

func main(){
  //name := "jojo"
  name := "joki"

  // penulisan else hrs kek gini, atau nanti error
  if name == "jojo"{
    fmt.Println("Hello jojo")
  } else if name == "purwo"{
    fmt.Println("hello purwo")
  } else if name == "joki"{
    fmt.Println("hello joki")
  } else {    
    fmt.Println("hi, kenalan yuk")
  }

  // if golang bisa membuat short statement (engga ada di yg lain)
  if length := len(name); length > 5{
    fmt.Println("terlalu panjang")
  } else{
    fmt.Println("benar")
  }
  
  // error krn length = local var
  fmt.Println(length)
}