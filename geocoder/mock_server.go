package geocoder

import (
  "net/http"
  "net/http/httptest"
  net_url "net/url"
  "os"
)

func newMockServer() (*httptest.Server, *net_url.URL, error) {
  mux := http.NewServeMux()

  mockBenchmarks, err := os.ReadFile("testdata/responses/benchmarks.json")
  if err != nil {
    return nil, nil, err
  }

  mux.HandleFunc("/benchmarks", func(w http.ResponseWriter, r *http.Request) {
    w.Write(mockBenchmarks)
  })

  mockVintages, err := os.ReadFile("testdata/responses/vintages.json")
  if err != nil {
    return nil, nil, err
  }

  mux.HandleFunc("/vintages", func(w http.ResponseWriter, r *http.Request) {
    w.Write(mockVintages)
  })

  mockLocations, err := os.ReadFile("testdata/responses/locations.json")
  if err != nil {
    return nil, nil, err
  }

  mux.HandleFunc("/locations/onelineaddress", func(w http.ResponseWriter, r *http.Request) {
    w.Write(mockLocations)
  })

  mockGeographies, err := os.ReadFile("testdata/responses/geographies.json")
  if err != nil {
    return nil, nil, err
  }

  mux.HandleFunc("/geographies/onelineaddress", func(w http.ResponseWriter, r *http.Request) {
    w.Write(mockGeographies)
  })

  mockBatchLocations, err := os.ReadFile("testdata/responses/batch-locations-2020.csv")
  if err != nil {
    return nil, nil, err
  }

  mux.HandleFunc("/locations/addressbatch", func(w http.ResponseWriter, r *http.Request) {
    w.Write(mockBatchLocations)
  })

  mockBatchGeographies, err := os.ReadFile("testdata/responses/batch-geographies-2020-2020.csv")
  if err != nil {
    return nil, nil, err
  }

  mux.HandleFunc("/geographies/addressbatch", func(w http.ResponseWriter, r *http.Request) {
    w.Write(mockBatchGeographies)
  })

  // create mock server
  server := httptest.NewServer(mux)

  // parse server URL
  url, err := net_url.Parse(server.URL)
  if err != nil {
    return nil, nil, err
  }

  return server, url, nil
}
