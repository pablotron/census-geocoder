# census-geocoder

[Go][] wrapper for [Census geocoder][].

## Installation

```
go get https://pablotron.org/census-geocoder
```

## Example

Minimal tool which geocodes command-line arguments and prints the
normalized address from the results to standard output:

```go
package main

import (
  "fmt"
  "log"
  "os"
  "pablotron.org/census-geocoder/geocoder"
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
```

## Documentation

See <https://pkg.go.dev/pablotron.org/census-geocoder/geocoder>

[go]: https://go.dev/
  "Go programming language."
[census geocoder]: https://geocoding.geo.census.gov/geocoder/Geocoding_Services_API.html
  "Census geocoding API."
