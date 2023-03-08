package main
import "fmt"

// beda vararg dan array, kalo array perlu ada deklarasi array,
// vararg bisa langsung kirim lebih dr 1 data, tinggal kasi koma
// jd untungnya pake vararg, kalo kt butuh param berupa array/slice,
// kt engga perlu buat array dl baru dimskin ke func, 
// tp bs langsung mskin angka ke param func nya dan jg engga wajib,
// kasi param ke func nya, bisa diisi kosong
// variadic func ditandai dengan ... di deklarasi param
// vararg cmn blh ada 1 di 1 func dan diletakan di akhir 
// (msh bs ada var param yg lain asal engga vararg)

func sumAll(numbers ...int) int {
  total := 0

  for _, value := range numbers {
    total += value
  }
  return total
}

func main() {
  total := sumAll(10, 90, 30, 20)     // param bs diisi kosong
  fmt.Println(total)
  
  // kalo dah terlanjur buat slice, bs lgsg dimskin jg ke func nya
  slice := []int{20, 20, 30, 40, 50}
  total = sumAll(slice...)
  fmt.Println(total)
}