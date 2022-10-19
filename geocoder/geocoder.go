// Census geocoder.
//
// TODO:
// - structured address input (street, city, zip)
// - coordinate input (x, y)
// - batch uploads
// - custom geography layers
package geocoder

import (
  "encoding/json"
  "errors"
  "net/http"
  net_url "net/url"
)

// Benchmark from Benchmarks()
type Benchmark struct {
  Id string `json:"id"`
  Name string `json:"benchmarkName"`
  Description string `json:"benchmarkDescription"`
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
  Id string `json:"tigerLineId"`
  Side string `json:"side"`
}

// Address match result from Locations().
type AddressMatch struct {
  TigerLine TigerLine `json:"tigerLine"`
  Coordinates Coordinates `json:"coordinates"`

  AddressComponents struct {
    Zip string `json:"zip"`
    StreetName string `json:"streetName"`
    PreType string `json:"preType"`
    City string `json:"city"`
    PreDirection string `json:"preDirection"`
    SuffixDirection string `json:"suffixDirection"`
    FromAddress string `json:"fromAddress"`
    State string `json:"state"`
    SuffixType string `json:"suffixType"`
    ToAddress string `json:"toAddress"`
    SuffixQualifier string `json:"suffixQualifier"`
    PreQualifier string `json:"preQualifier"`
  } `json:"addressComponents"`

  MatchedAddress string `json:"matchedAddress"`

  Geographies map[string][]map[string]any `json:"geographies"`
}

// Census geocoder.
type Geocoder struct {
  url *net_url.URL
}

// Parsed default geocoder URL.
var DefaultUrl = &net_url.URL{
  Scheme: "https",
  Host: "geocoding.geo.census.gov",
  Path: "/geocoder/",
}

// Default geocoder.
var DefaultGeocoder = Geocoder { DefaultUrl }

// Create new geocoder from parsed URL.
func NewGeocoder(url *net_url.URL) Geocoder {
  return Geocoder { url }
}

// Build request, send to API endpoint, and parse response.
func (g Geocoder) get(path string, args map[string]string, cb func(*json.Decoder) error) error {
  // build url
  url := g.url.JoinPath(path)

  // build query parameters
  q := net_url.Values{}
  for k, v := range(args) {
    q.Add(k, v)
  }
  url.RawQuery = q.Encode()

  // fetch response
  resp, err := http.Get(url.String())
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  // create decoder from response body, call handler
  return cb(json.NewDecoder(resp.Body))
}


// Get benchmarks from geocoder.
//
// Example: https://geocoding.geo.census.gov/geocoder/benchmarks
func (g Geocoder) Benchmarks() ([]Benchmark, error) {
  var r struct {
    Benchmarks []Benchmark `json:"benchmarks"`
		Errors []string `json:"errors"`
  }

  // send request, parse response
  args := map[string]string{}
  err := g.get("benchmarks", args, func(d *json.Decoder) error {
    return d.Decode(&r)
  })

  // check for errors
  if err != nil {
    return []Benchmark{}, err
  }

  // check for errors in decoded response
  if len(r.Errors) > 0 {
    return []Benchmark{}, errors.New(r.Errors[0])
  }

  // return result
  return r.Benchmarks, nil
}

// Get vintages matching benchmark ID from default geocoder.
//
// Example: https://geocoding.geo.census.gov/geocoder/vintages?benchmark=4
func (g Geocoder) Vintages(benchmarkId string) ([]Vintage, error) {
  var r struct {
    Vintages []Vintage `json:"vintages"`
		Errors []string `json:"errors"`
  }

  // send request, parse response
  err := g.get("vintages", map[string]string {
    "benchmark": benchmarkId,
  }, func(d *json.Decoder) error {
    return d.Decode(&r)
  })

  // check for request errors
  if err != nil {
    return []Vintage{}, err
  }

  // check for errors in decoded response
  if len(r.Errors) > 0 {
    return []Vintage{}, errors.New(r.Errors[0])
  }

  // return result
  return r.Vintages, nil
}

// Geocode street address with given benchmark ID return address
// matches.
//
// Example: https://geocoding.geo.census.gov/geocoder/locations/onelineaddress?address=4600+Silver+Hill+Rd%2C+Washington%2C+DC+20233&benchmark=2020&format=json
func (g Geocoder) LocationsFromBenchmark(address, benchmarkId string) ([]AddressMatch, error) {
  var r struct {
    Result struct {
      AddressMatches []AddressMatch `json:"addressMatches"`
    } `json:"result"`

		Errors []string `json:"errors"`
  }

  // send request, decode response
  err := g.get("locations/onelineaddress", map[string]string {
    "address": address,
    "benchmark": benchmarkId,
    "format": "json",
  }, func(d *json.Decoder) error {
    return d.Decode(&r)
  })

  // check for request errors
  if err != nil {
    return []AddressMatch{}, err
  }

  // check for errors in decoded response
  if len(r.Errors) > 0 {
    return []AddressMatch{}, errors.New(r.Errors[0])
  }

  // return result
  return r.Result.AddressMatches, nil
}

// Default benchmark ID.
var DefaultBenchmark = "Public_AR_Current"

// Geocode street address and return address matches.
func (g Geocoder) Locations(address string) ([]AddressMatch, error) {
  return g.LocationsFromBenchmark(address, DefaultBenchmark)
}

// Geocode street address using  given benchmark and given vintage, then
// return address matches with geography layers.
//
// Example: https://geocoding.geo.census.gov/geocoder/geographies/address?street=4600+Silver+Hill+Rd&city=Washington&state=DC&benchmark=Public_AR_Census2020&vintage=Census2020_Census2020&layers=10&format=json
func (g Geocoder) Geographies(address, benchmark, vintage string) ([]AddressMatch, error) {
  var r struct {
    Result struct {
      AddressMatches []AddressMatch `json:"addressMatches"`
    } `json:"result"`

		Errors []string `json:"errors"`
  }

  // send request, decode response
  err := g.get("geographies/onelineaddress", map[string]string {
    "address": address,
    "benchmark": benchmark,
    "vintage": vintage,
    "format": "json",
  }, func(d *json.Decoder) error {
    return d.Decode(&r)
  })

  // check for request errors
  if err != nil {
    return []AddressMatch{}, err
  }

  // check for errors in decoded response
  if len(r.Errors) > 0 {
    return []AddressMatch{}, errors.New(r.Errors[0])
  }

  // return result
  return r.Result.AddressMatches, nil
}

// Get benchmarks from default geocoder.
func Benchmarks() ([]Benchmark, error) {
  return DefaultGeocoder.Benchmarks()
}

// Get vintages matching benchmark ID from default geocoder.
func Vintages(benchmarkId string) ([]Vintage, error) {
  return DefaultGeocoder.Vintages(benchmarkId)
}

// Geocode street address with given benchmark ID using default geocoder
// and return address matches.
func LocationsFromBenchmark(address, benchmarkId string) ([]AddressMatch, error) {
  return DefaultGeocoder.LocationsFromBenchmark(address, benchmarkId)
}

// Geocode street address using default geocoder and return address
// matches.
func Locations(address string) ([]AddressMatch, error) {
  return DefaultGeocoder.Locations(address)
}

// Geocode street address using default geocoder, given benchmark, and
// given vintage, then return address matches with geography layers.
func Geographies(address, benchmark, vintage string) ([]AddressMatch, error) {
  return DefaultGeocoder.Geographies(address, benchmark, vintage)
}
