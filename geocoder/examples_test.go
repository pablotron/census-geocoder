package geocoder

import (
  "fmt"
  "log"
)

func ExampleGeocoder_Benchmarks() {
  // create geocoder using default geocoder URL
  g := NewGeocoder(DefaultUrl)

  // get benchmarks
  benchmarks, err := g.Benchmarks()
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

func ExampleGeocoder_Vintages() {
  // create geocoder using default geocoder URL
  g := NewGeocoder(DefaultUrl)

  // get vintages
  vintages, err := g.Vintages("2020")
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

func ExampleGeocoder_Locations() {
  // create geocoder using default geocoder URL
  g := NewGeocoder(DefaultUrl)

  // get address matches
  locs, err := g.Locations("3444 gallows rd annandale va 22003")
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

func ExampleGeocoder_Geographies() {
  // create geocoder using default geocoder URL
  g := NewGeocoder(DefaultUrl)

  // get address matches with additional geographical information
  locs, err := g.Geographies("3444 gallows rd annandale va 22003", "2020", "2020")
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
