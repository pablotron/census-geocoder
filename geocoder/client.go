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

// Census geocoder client.
type Client struct {
  // base API URL
  Url *net_url.URL

  // shared HTTP client
  Client http.Client
}

// Create new geocoder client from URL.
//
// Note: This method is primarily for testing.  You should be able to
// use the top-level functions from the [geocoder] package (e.g.
// [geocoder.Locations], [geocoder.Geographies], etc).
func NewClient(url *net_url.URL) Client {
  return Client { Url: url }
}

// Build request, send to API endpoint, and parse response.
func (c Client) get(path string, args map[string]string, cb func(*json.Decoder) error) error {
  // build url
  url := c.Url.JoinPath(path)

  // build query parameters
  q := net_url.Values{}
  for k, v := range(args) {
    q.Add(k, v)
  }
  url.RawQuery = q.Encode()

  // fetch response
  resp, err := c.Client.Get(url.String())
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  // create decoder from response body, call handler
  return cb(json.NewDecoder(resp.Body))
}

// Get available benchmarks.
func (c Client) Benchmarks() ([]Benchmark, error) {
  var r struct {
    Benchmarks []Benchmark `json:"benchmarks"`
		Errors []string `json:"errors"`
  }

  // send request, parse response
  args := map[string]string{}
  err := c.get("benchmarks", args, func(d *json.Decoder) error {
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

// Get vintages matching benchmark ID.
func (c Client) Vintages(benchmarkId string) ([]Vintage, error) {
  var r struct {
    Vintages []Vintage `json:"vintages"`
		Errors []string `json:"errors"`
  }

  // send request, parse response
  err := c.get("vintages", map[string]string {
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
func (c Client) LocationsFromBenchmark(address, benchmarkId string) ([]AddressMatch, error) {
  var r struct {
    Result struct {
      AddressMatches []AddressMatch `json:"addressMatches"`
    } `json:"result"`

		Errors []string `json:"errors"`
  }

  // send request, decode response
  err := c.get("locations/onelineaddress", map[string]string {
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

// Geocode street address and return address matches.
func (c Client) Locations(address string) ([]AddressMatch, error) {
  return c.LocationsFromBenchmark(address, DefaultBenchmark)
}

// Geocode street address using  given benchmark and given vintage, then
// return address matches with geography layers.
func (c Client) Geographies(address, benchmark, vintage string) ([]AddressMatch, error) {
  var r struct {
    Result struct {
      AddressMatches []AddressMatch `json:"addressMatches"`
    } `json:"result"`

		Errors []string `json:"errors"`
  }

  // send request, decode response
  err := c.get("geographies/onelineaddress", map[string]string {
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
func (c Client) batchUpload(rows []BatchInputRow, returnType string, fields map[string]string) ([]BatchOutputRow, error) {
  // populate buffer with multipart-encoded request body
  var buf bytes.Buffer
  contentType, err := createBatchBody(&buf, rows, fields)
  if err != nil {
    return []BatchOutputRow{}, err
  }

  // build url
  url := c.Url.JoinPath(returnType, "addressbatch")

  // create request
  req, err := http.NewRequest("POST", url.String(), &buf)
  if err != nil {
    return []BatchOutputRow{}, err
  }

  // set request headers
  req.Header.Add("Content-Type", contentType)

  // send request
  resp, err := c.Client.Do(req)
  if err != nil {
    return []BatchOutputRow{}, err
  }
  defer resp.Body.Close()

  // read rows from response
  return NewBatchOutputReader(resp.Body).ReadAll()
}

// Batch geocode street addresses with given benchmark then return
// matches.
func (c Client) BatchLocationsFromBenchmark(rows []BatchInputRow, benchmark string) ([]BatchOutputRow, error) {
  return c.batchUpload(rows, "locations", map[string]string {
    "benchmark": benchmark,
  })
}

// Batch geocode street addresses then return matches.
func (c Client) BatchLocations(rows []BatchInputRow) ([]BatchOutputRow, error) {
  return c.BatchLocationsFromBenchmark(rows, DefaultBenchmark)
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
func (c Client) BatchGeographies(rows []BatchInputRow, benchmark, vintage string) ([]BatchOutputRow, error) {
  return c.batchUpload(rows, "geographies", map[string]string {
    "benchmark": benchmark,
    "vintage": vintage,
  })
}
