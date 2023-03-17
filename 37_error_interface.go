package main
import (
  "errors"
  "fmt"
)

/**
error interface adl kontrak default utk membuat error dg func nya adl
Error()
*/

func pembagi(nilai, pembagi int) (int, error) {
  if pembagi == 0 {
    return 0, errors.New("pembagi engga boleh 0")
  } else {
    result := nilai / pembagi
    return result, nil
  }
}

func main() {
  hasil, err := pembagi(100, 0)
  if err == nil {
    fmt.Println("hasil", hasil)
  } else {
    fmt.Println("error", err.Error())
  }
}