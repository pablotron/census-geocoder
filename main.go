package main

import (
  "fmt"
  "log"
  "pmdn.org/census-geocoder/geocoder"
  "os"
)

func main() {
  for _, arg := range(os.Args[1:]) {
    // get addresses
    addresses, err := geocoder.Locations(arg)
    if err != nil {
      log.Fatal(err)
    }

    // print addresses
    for _, address := range(addresses) {
      fmt.Println(address)
    }
  }
}
