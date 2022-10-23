package geocoder

import (
  "fmt"
  "log"
)

func ExampleClient_Benchmarks() {
  // create client using default URL
  c := NewClient(DefaultUrl)

  // get benchmarks
  benchmarks, err := c.Benchmarks()
  if err != nil {
    log.Fatal(err)
  }

  // print benchmark names to standard output
  for _, b := range(benchmarks) {
    fmt.Println(b.Name)
  }

  // Unordered output:
  // Public_AR_Current
  // Public_AR_ACS2022
  // Public_AR_Census2020
}

func ExampleClient_Vintages() {
  // create client using default URL
  c := NewClient(DefaultUrl)

  // get vintages
  vintages, err := c.Vintages("2020")
  if err != nil {
    log.Fatal(err)
  }

  // print vintage names to standard output
  for _, v := range(vintages) {
    fmt.Println(v.Name)
  }

  // Unordered output:
  // Census2020_Census2020
  // Census2010_Census2020
}

func ExampleClient_Locations() {
  // create client using default URL
  c := NewClient(DefaultUrl)

  // get address matches
  locs, err := c.Locations("3444 gallows rd annandale va 22003")
  if err != nil {
    log.Fatal(err)
  }

  // print matched addresses to standard output
  for _, v := range(locs) {
    fmt.Println(v.MatchedAddress)
  }

  // Unordered output:
  // 3444 GALLOWS RD, ANNANDALE, VA, 22003
}

func ExampleClient_Geographies() {
  // create client using default URL
  c := NewClient(DefaultUrl)

  // get address matches with additional geographical information
  locs, err := c.Geographies("3444 gallows rd annandale va 22003", "2020", "2020")
  if err != nil {
    log.Fatal(err)
  }

  // print matched addresses to standard output
  for _, v := range(locs) {
    cbsaName := v.Geographies["Combined Statistical Areas"][0]["NAME"]
    fmt.Printf("%s - %s\n", v.MatchedAddress, cbsaName)
  }

  // Unordered output:
  // 3444 GALLOWS RD, ANNANDALE, VA, 22003 - Washington-Baltimore-Arlington, DC-MD-VA-WV-PA CSA
}

func ExampleBenchmarks() {
  // get benchmarks
  benchmarks, err := Benchmarks()
  if err != nil {
    log.Fatal(err)
  }

  // print benchmark names to standard output
  for _, b := range(benchmarks) {
    fmt.Println(b.Name)
  }

  // Unordered output:
  // Public_AR_Current
  // Public_AR_ACS2022
  // Public_AR_Census2020
}

func ExampleVintages() {
  // get vintages
  vintages, err := Vintages("2020")
  if err != nil {
    log.Fatal(err)
  }

  // print vintage names to standard output
  for _, v := range(vintages) {
    fmt.Println(v.Name)
  }

  // Unordered output:
  // Census2020_Census2020
  // Census2010_Census2020
}

func ExampleLocations() {
  // get address matches
  locs, err := Locations("3444 gallows rd annandale va 22003")
  if err != nil {
    log.Fatal(err)
  }

  // print matched addresses to standard output
  for _, v := range(locs) {
    fmt.Println(v.MatchedAddress)
  }

  // Unordered output:
  // 3444 GALLOWS RD, ANNANDALE, VA, 22003
}

func ExampleGeographies() {
  // get address matches with additional geographical information
  locs, err := Geographies("3444 gallows rd annandale va 22003", "2020", "2020")
  if err != nil {
    log.Fatal(err)
  }

  // print matched addresses to standard output
  for _, v := range(locs) {
    cbsaName := v.Geographies["Combined Statistical Areas"][0]["NAME"]
    fmt.Printf("%s - %s\n", v.MatchedAddress, cbsaName)
  }

  // Unordered output:
  // 3444 GALLOWS RD, ANNANDALE, VA, 22003 - Washington-Baltimore-Arlington, DC-MD-VA-WV-PA CSA
}
