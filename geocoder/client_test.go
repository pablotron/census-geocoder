package geocoder

import (
  _ "embed"
  "encoding/json"
  "reflect"
  "testing"
)

func TestClientBenchmarks(t *testing.T) {
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

  // create client
  c := NewClient(url)

  // create benchmarks
  got, err := c.Benchmarks()
  if err != nil {
    t.Fatal(err)
  }

  // compare against expected value
  if !reflect.DeepEqual(got, exp) {
    t.Fatalf("got %v, exp %v", got, exp)
  }
}

func TestClientVintages(t *testing.T) {
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

  // create client
  c := NewClient(url)

  // get vintages, check for error
  got, err := c.Vintages(testBenchmarkId)
  if err != nil {
    t.Fatal(err)
  }

  // compare against expected value
  if !reflect.DeepEqual(got, exp) {
    t.Fatalf("got %v, exp %v", got, exp)
  }
}

func TestClientLocationsFromBenchmark(t *testing.T) {
  // create mock server
  ms, url, err := newMockServer()
  if err != nil {
    t.Fatal(err)
  }
  defer ms.Close()

  // decode expected results
  var exp []AddressMatch
  if err = json.Unmarshal(mockLocationsJson, &exp); err != nil {
    t.Fatal(err)
  }

  // create client
  c := NewClient(url)

  // get locations, check for error
  got, err := c.LocationsFromBenchmark(testAddress, testBenchmarkId)
  if err != nil {
    t.Fatal(err)
  }

  // compare against expected value
  if !reflect.DeepEqual(got, exp) {
    t.Fatalf("got %v, exp %v", got, exp)
  }
}

func TestClientLocations(t *testing.T) {
  // create mock server
  ms, url, err := newMockServer()
  if err != nil {
    t.Fatal(err)
  }
  defer ms.Close()

  // decode expected results
  var exp []AddressMatch
  if err = json.Unmarshal(mockLocationsJson, &exp); err != nil {
    t.Fatal(err)
  }

  // create client
  c := NewClient(url)

  // get locations, check for error
  got, err := c.Locations(testAddress)
  if err != nil {
    t.Fatal(err)
  }

  // compare against expected value
  if !reflect.DeepEqual(got, exp) {
    t.Fatalf("got %v, exp %v", got, exp)
  }
}

func TestClientGeographies(t *testing.T) {
  testAddress := "4600 silver hill rd, 20233"
  testBenchmark := "Public_AR_Census2020"
  testVintage := "Census2010_Census2020"

  // create mock server
  ms, url, err := newMockServer()
  if err != nil {
    t.Fatal(err)
  }
  defer ms.Close()

  // decode expected results
  var exp []AddressMatch
  if err = json.Unmarshal(mockGeographiesJson, &exp); err != nil {
    t.Fatal(err)
  }

  // create client
  c := NewClient(url)

  // get geographies, check for error
  got, err := c.Geographies(testAddress, testBenchmark, testVintage)
  if err != nil {
    t.Fatal(err)
  }

  // compare against expected value
  if !reflect.DeepEqual(got, exp) {
    t.Fatalf("got %v, exp %v", got, exp)
  }
}

func TestClientBatchLocationsFromBenchmark(t *testing.T) {
  // get input and expected output
  rows := getBatchInputRows(t)
  exp := getBatchOutputRows(t, "testdata/data/batch-output-locations-2020.csv")

  // create mock server
  ms, url, err := newMockServer()
  if err != nil {
    t.Fatal(err)
  }
  defer ms.Close()

  // create client
  c := NewClient(url)

  // send rows, check for error
  got, err := c.BatchLocationsFromBenchmark(rows, testBenchmarkId)
  if err != nil {
    t.Fatal(err)
  }

  // compare against expected value
  if !reflect.DeepEqual(got, exp) {
    t.Fatalf("got %v, exp %v", got, exp)
  }
}

func TestClientBatchLocations(t *testing.T) {
  // get input and expected output
  rows := getBatchInputRows(t)

  // get expected output, build map
  expRows := getBatchOutputRows(t, "testdata/data/batch-output-locations-2020.csv")
  exp := make(map[string]BatchOutputRow)
  for _, row := range(expRows) {
    exp[row.Id] = row
  }

  // create mock server
  ms, url, err := newMockServer()
  if err != nil {
    t.Fatal(err)
  }
  defer ms.Close()

  // create client
  c := NewClient(url)

  // send rows, check for error
  gotRows, err := c.BatchLocations(rows)
  if err != nil {
    t.Fatal(err)
  }

  // compare against expected rows
  for _, row := range(gotRows) {
    if !compareBatchOutputRow(row, exp[row.Id]) {
      t.Fatalf("%s: got %v, exp %v", row.Id, row, exp[row.Id])
    }
  }
}

func TestClientBatchGeographies(t *testing.T) {
  // get input and expected output
  rows := getBatchInputRows(t)
  exp := getBatchOutputRows(t, "testdata/data/batch-output-geographies-2020-2020.csv")

  // create mock server
  ms, url, err := newMockServer()
  if err != nil {
    t.Fatal(err)
  }
  defer ms.Close()

  // create client
  c := NewClient(url)

  // send rows, check for error
  got, err := c.BatchGeographies(rows, "2020", "2020")
  if err != nil {
    t.Fatal(err)
  }

  // compare against expected value
  if !reflect.DeepEqual(got, exp) {
    t.Fatalf("got %v, exp %v", got, exp)
  }
}
