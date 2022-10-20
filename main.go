// minimal address geocoder
package main

import (
  "fmt"
  "log"
  "os"
  "pmdn.org/census-geocoder/geocoder"
)

func main() {
  for _, arg := range(os.Args[1:]) {
    // get address matches
    matches, err := geocoder.Locations(arg)
    if err != nil {
      log.Fatal(err)
    }

    // print matches
    for _, match := range(matches) {
      fmt.Println(match)
    }
  }
}
