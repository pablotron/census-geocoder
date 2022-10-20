// Census geocoder.
package geocoder

import (
  "bytes"
  "encoding/json"
  "errors"
  "io"
  "mime/multipart"
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
  // base API URL
  Url *net_url.URL

  // shared HTTP client
  Client http.Client
}

// Parsed default geocoder URL.
var DefaultUrl = &net_url.URL{
  Scheme: "https",
  Host: "geocoding.geo.census.gov",
  Path: "/geocoder/",
}

// Default geocoder.
var DefaultGeocoder = Geocoder { Url: DefaultUrl }

// Create new geocoder from parsed URL.
func NewGeocoder(url *net_url.URL) Geocoder {
  return Geocoder { Url: url }
}

// Build request, send to API endpoint, and parse response.
func (g Geocoder) get(path string, args map[string]string, cb func(*json.Decoder) error) error {
  // build url
  url := g.Url.JoinPath(path)

  // build query parameters
  q := net_url.Values{}
  for k, v := range(args) {
    q.Add(k, v)
  }
  url.RawQuery = q.Encode()

  // fetch response
  resp, err := g.Client.Get(url.String())
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

// Encode batch input rows and field values as multipart body and write
// it to the given writer.
func createBatchBody(w io.Writer, rows []BatchInputRow, fields map[string]string) (string, error) {
  // create multipart writer
  mw := multipart.NewWriter(w)

  // populate form fields
  for k, v := range(fields) {
    if err := mw.WriteField(k, v); err != nil {
      return "", err
    }
  }

  // attach address file
  f, err := mw.CreateFormFile("addressFile", "input.csv")
  if err != nil {
    return "", err
  }

  // write input rows to multipart writer as CSV
  biw := NewBatchInputWriter(f)
  if err := biw.WriteAll(rows); err != nil {
    return "", err
  }

  // get content type
  contentType := mw.FormDataContentType()

  // close multipart writer
  return contentType, mw.Close()
}

// Upload input addresses to batch geocoder.
func (g Geocoder) batchUpload(rows []BatchInputRow, returnType string, fields map[string]string) ([]BatchOutputRow, error) {
  // populate buffer with multipart-encoded request body
  var buf bytes.Buffer
  contentType, err := createBatchBody(&buf, rows, fields)
  if err != nil {
    return []BatchOutputRow{}, err
  }

  // build url
  url := g.Url.JoinPath(returnType, "addressbatch")

  // create request
  req, err := http.NewRequest("POST", url.String(), &buf)
  if err != nil {
    return []BatchOutputRow{}, err
  }

  // set request headers
  req.Header.Add("Content-Type", contentType)

  // send request
  resp, err := g.Client.Do(req)
  if err != nil {
    return []BatchOutputRow{}, err
  }
  defer resp.Body.Close()

  // read rows from response
  return NewBatchOutputReader(resp.Body).ReadAll()
}

// Batch geocode street addresses with given benchmark then return
// matches.
func (g Geocoder) BatchLocationsFromBenchmark(rows []BatchInputRow, benchmark string) ([]BatchOutputRow, error) {
  return g.batchUpload(rows, "locations", map[string]string {
    "benchmark": benchmark,
  })
}

// Batch geocode street addresses then return matches.
func (g Geocoder) BatchLocations(rows []BatchInputRow) ([]BatchOutputRow, error) {
  return g.BatchLocationsFromBenchmark(rows, DefaultBenchmark)
}

// Batch geocode street addresses with given benchmark and vintage then
// return matches with additional geography fields.
//
// The additional BatchOutputRow fields populated by this method
// compared to `BatchLocationsFromBenchmark()` are as follows:
//
// - State
// - County
// - Tract
// - Block
func (g Geocoder) BatchGeographies(rows []BatchInputRow, benchmark, vintage string) ([]BatchOutputRow, error) {
  return g.batchUpload(rows, "geographies", map[string]string {
    "benchmark": benchmark,
    "vintage": vintage,
  })
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

// Batch geocode street addresses with given benchmark using default
// geocoder then return matches.
func BatchLocationsFromBenchmark(rows []BatchInputRow, benchmark string) ([]BatchOutputRow, error) {
  return DefaultGeocoder.BatchLocationsFromBenchmark(rows, benchmark)
}

// Batch geocode street addresses using default geocoder then return
// matches.
func BatchLocations(rows []BatchInputRow) ([]BatchOutputRow, error) {
  return DefaultGeocoder.BatchLocations(rows)
}

// Batch geocode street addresses with given benchmark and vintage using
// default geocoder then return matches with additional geography fields.
//
// The additional BatchOutputRow fields populated by this function
// compared to `BatchLocationsFromBenchmark()` are as follows:
//
// - State
// - County
// - Tract
// - Block
func BatchGeographies(rows []BatchInputRow, benchmark, vintage string) ([]BatchOutputRow, error) {
  return DefaultGeocoder.BatchGeographies(rows, benchmark, vintage)
}
