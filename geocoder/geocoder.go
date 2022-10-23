// Census geocoder.
package geocoder

import (
  net_url "net/url"
)

// Benchmark from Benchmarks()
type Benchmark struct {
  // Vintage ID
  Id string `json:"id"`

  // Vintage name
  Name string `json:"benchmarkName"`

  // Vintage description
  Description string `json:"benchmarkDescription"`

  // Is this the default vintage?
  Default bool `json:"isDefault"`
}

// Vintage from Vintages().
type Vintage struct {
  Id string `json:"id"`
  Name string `json:"vintageName"`
  Description string `json:"vintageDescription"`
  Default bool `json:"isDefault"`
}

type TigerLine struct {
  // Line ID
  Id string `json:"tigerLineId"`

  // Line side
  Side string `json:"side"`
}

// Address match result from [Locations()] or [Geographies()].
type Match struct {
  // tiger data
  TigerLine TigerLine `json:"tigerLine"`

  // latitude and longitude
  Coordinates Coordinates `json:"coordinates"`

  // matched components
  AddressComponents struct {
    // zip code
    Zip string `json:"zip"`

    // street name
    StreetName string `json:"streetName"`

    // street prefix type
    PreType string `json:"preType"`

    // city
    City string `json:"city"`

    // prefix direction
    PreDirection string `json:"preDirection"`

    // suffix direction
    SuffixDirection string `json:"suffixDirection"`

    // from address
    FromAddress string `json:"fromAddress"`

    // state
    State string `json:"state"`

    // suffix type
    SuffixType string `json:"suffixType"`

    // to address
    ToAddress string `json:"toAddress"`

    // suffix qualifier
    SuffixQualifier string `json:"suffixQualifier"`

    // prefix qualifier
    PreQualifier string `json:"preQualifier"`
  } `json:"addressComponents"`

  // matched address
  MatchedAddress string `json:"matchedAddress"`

  // map of ID to geography components.
  //
  // Note: only populated for calls to `Geographies()`.
  Geographies map[string][]map[string]any `json:"geographies"`
}

// Default Census geocoder URL.
var DefaultUrl = &net_url.URL{
  Scheme: "https",
  Host: "geocoding.geo.census.gov",
  Path: "/geocoder/",
}

// Default client.
var DefaultClient = Client { Url: DefaultUrl }

// Default benchmark ID.
var DefaultBenchmark = "Public_AR_Current"

// Get benchmarks from default client.
func Benchmarks() ([]Benchmark, error) {
  return DefaultClient.Benchmarks()
}

// Get vintages matching benchmark ID from default client.
func Vintages(benchmarkId string) ([]Vintage, error) {
  return DefaultClient.Vintages(benchmarkId)
}

// Geocode street address with given benchmark ID using default client
// and return address matches.
func LocationsFromBenchmark(address, benchmarkId string) ([]Match, error) {
  return DefaultClient.LocationsFromBenchmark(address, benchmarkId)
}

// Geocode street address using default client and return address
// matches.
func Locations(address string) ([]Match, error) {
  return DefaultClient.Locations(address)
}

// Geocode street address using default client, given benchmark, and
// given vintage, then return address matches with geography layers.
func Geographies(address, benchmark, vintage string) ([]Match, error) {
  return DefaultClient.Geographies(address, benchmark, vintage)
}

// Batch geocode street addresses with given benchmark using default
// client then return matches.
func BatchLocationsFromBenchmark(rows []BatchInputRow, benchmark string) ([]BatchOutputRow, error) {
  return DefaultClient.BatchLocationsFromBenchmark(rows, benchmark)
}

// Batch geocode street addresses using default client then return
// matches.
func BatchLocations(rows []BatchInputRow) ([]BatchOutputRow, error) {
  return DefaultClient.BatchLocations(rows)
}

// Batch geocode street addresses with given benchmark and vintage using
// default client, then return matches with additional geography fields.
//
// The additional BatchOutputRow fields populated by this function
// compared to `BatchLocationsFromBenchmark()` are as follows:
//
// - State
// - County
// - Tract
// - Block
func BatchGeographies(rows []BatchInputRow, benchmark, vintage string) ([]BatchOutputRow, error) {
  return DefaultClient.BatchGeographies(rows, benchmark, vintage)
}
