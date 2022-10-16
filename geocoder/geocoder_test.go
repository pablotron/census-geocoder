package geocoder

import (
  _ "embed"
  "encoding/json"
  "reflect"
  "testing"
)

// func TestNewGeocoder(t *testing.T) {
//   t.Run("valid_url", func(t *testing.T) {
//     // create geocoder
//     if _, err := NewGeocoder(DefaultUrl); err != nil {
//       t.Fatal(err)
//     }
//   })
// 
//   t.Run("invalid_url", func(t *testing.T) {
//     badUrl := "lkfsa\\dad"
// 
//     // create geocoder
//     gc, err := NewGeocoder(badUrl)
//     if err != nil {
//       t.Fatalf("got %v, exp error", gc)
//     }
//   })
// }

//go:embed testdata/data/benchmarks.json
var mockBenchmarksJson []byte

func TestGeocoderBenchmarks(t *testing.T) {
  // create mock server
  ms, url, err := newMockServer()
  if err != nil {
    t.Fatal(err)
  }
  defer ms.Close()

  // decode expected results
  var exp []Benchmark
  if err := json.Unmarshal(mockBenchmarksJson, &exp); err != nil {
    t.Fatal(err)
  }

  // create geocoder
  gc := NewGeocoder(url)

  // create benchmarks
  got, err := gc.Benchmarks()
  if err != nil {
    t.Fatal(err)
  }

  // compare against expected value
  if !reflect.DeepEqual(got, exp) {
    t.Fatalf("got %v, exp %v", got, exp)
  }
}

func TestBenchmarks(t *testing.T) {
  if testing.Short() {
    t.Skip("skipping in short mode")
  }

  // get benchmarks, check for error
  if _, err := Benchmarks(); err != nil {
    t.Fatal(err)
  }
}

//go:embed testdata/data/vintages.json
var mockVintagesJson []byte

// test benchmark ID
var testBenchmarkId = "Public_AR_Current"

// test address
var testAddress = "4600 Silver Hill Rd, Washington, DC 20233"

func TestGeocoderVintages(t *testing.T) {
  // create mock server
  ms, url, err := newMockServer()
  if err != nil {
    t.Fatal(err)
  }
  defer ms.Close()

  // decode expected results
  var exp []Vintage
  if err = json.Unmarshal(mockVintagesJson, &exp); err != nil {
    t.Fatal(err)
  }

  // create geocoder
  gc := NewGeocoder(url)

  // get vintages, check for error
  got, err := gc.Vintages(testBenchmarkId)
  if err != nil {
    t.Fatal(err)
  }

  // compare against expected value
  if !reflect.DeepEqual(got, exp) {
    t.Fatalf("got %v, exp %v", got, exp)
  }
}

func TestVintages(t *testing.T) {
  if testing.Short() {
    t.Skip("skipping in short mode")
  }

  // get results, check for error
  if _, err := Vintages(testBenchmarkId); err != nil {
    t.Fatal(err)
  }
}

//go:embed testdata/data/address-matches.json
var mockAddressMatchesJson []byte

func TestGeocoderLocationsFromBenchmark(t *testing.T) {
  // create mock server
  ms, url, err := newMockServer()
  if err != nil {
    t.Fatal(err)
  }
  defer ms.Close()

  // decode expected results
  var exp []AddressMatch
  if err = json.Unmarshal(mockAddressMatchesJson, &exp); err != nil {
    t.Fatal(err)
  }

  // create geocoder
  gc := NewGeocoder(url)

  // get locations, check for error
  got, err := gc.LocationsFromBenchmark(testAddress, testBenchmarkId)
  if err != nil {
    t.Fatal(err)
  }

  // compare against expected value
  if !reflect.DeepEqual(got, exp) {
    t.Fatalf("got %v, exp %v", got, exp)
  }
}

func TestLocationsFromBenchmark(t *testing.T) {
  if testing.Short() {
    t.Skip("skipping in short mode")
  }

  // decode expected results
  var exp []AddressMatch
  if err := json.Unmarshal(mockAddressMatchesJson, &exp); err != nil {
    t.Fatal(err)
  }

  // get locations, check for error
  _, err := LocationsFromBenchmark(testAddress, testBenchmarkId)
  if err != nil {
    t.Fatal(err)
  }
}

func TestGeocoderLocations(t *testing.T) {
  // create mock server
  ms, url, err := newMockServer()
  if err != nil {
    t.Fatal(err)
  }
  defer ms.Close()

  // decode expected results
  var exp []AddressMatch
  if err = json.Unmarshal(mockAddressMatchesJson, &exp); err != nil {
    t.Fatal(err)
  }

  // create geocoder
  gc := NewGeocoder(url)

  // get locations, check for error
  got, err := gc.Locations(testAddress)
  if err != nil {
    t.Fatal(err)
  }

  // compare against expected value
  if !reflect.DeepEqual(got, exp) {
    t.Fatalf("got %v, exp %v", got, exp)
  }
}

func TestLocations(t *testing.T) {
  if testing.Short() {
    t.Skip("skipping in short mode")
  }

  // get locations, check for error
  _, err := Locations(testAddress)
  if err != nil {
    t.Fatal(err)
  }
}
