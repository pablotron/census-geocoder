package geocoder

import (
  "log"
  "net/http"
  "net/http/httptest"
  net_url "net/url"
  "os"
)

// Add mock handler to mux at given URL path.
func addHandler(mux *http.ServeMux, urlPath, dataPath string) error {
  // read mock data
  data, err := os.ReadFile(dataPath)
  if err != nil {
    return err
  }

  // add endpoint handler
  mux.HandleFunc(urlPath, func(w http.ResponseWriter, r *http.Request) {
    // write response, check for error
    if _, err := w.Write(data); err != nil {
      // log error
      log.Print(err)
    }
  })

  // return success
  return nil
}

// mock server endpoints and data
var mockEndpoints = []struct {
  urlPath string // url path
  dataPath string // mock data path
} {
  { "/benchmarks", "testdata/responses/benchmarks.json" },
  { "/vintages", "testdata/responses/vintages.json" },
  { "/locations/onelineaddress", "testdata/responses/locations.json" },
  { "/geographies/onelineaddress", "testdata/responses/geographies.json" },
  { "/locations/addressbatch", "testdata/responses/batch-locations-2020.csv" },
  { "/geographies/addressbatch", "testdata/responses/batch-geographies-2020-2020.csv" },
}

// Start new mock server, then return server and server URL.
func newMockServer() (*httptest.Server, *net_url.URL, error) {
  // create mux
  mux := http.NewServeMux()

  // add handlers
  for _, row := range(mockEndpoints) {
    if err := addHandler(mux, row.urlPath, row.dataPath); err != nil {
      return nil, nil, err
    }
  }

  // create mock server
  server := httptest.NewServer(mux)

  // parse server URL
  url, err := net_url.Parse(server.URL)
  if err != nil {
    return nil, nil, err
  }

  // return server and url
  return server, url, nil
}
