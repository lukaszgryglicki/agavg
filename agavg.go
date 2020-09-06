package main

import (
	"fmt"
)

func getComplex(str string) (c complex128) {
  str = strings.ToLower(str)
  return
}

func main() {
	if len(os.Args) < 3 {
    fmt.Printf("Usage: %s '1.2+3.1i' '-0.5-1.12i'\n")
		return
	}
  x, err := getComplex(os.Args[1])
  if err != nil {
    fmt.Fatalf("%s: %+v\n", os.Args[1], err)
  }
  y, err := getComplex(os.Args[2])
  if err != nil {
    fmt.Fatalf("%s: %+v\n", os.Args[2], err)
  }
 fmt.Printf("%v %v\n", x, y)
}
