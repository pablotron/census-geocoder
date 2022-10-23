package geocoder

import (
  _ "embed"
  "encoding/json"
  "os"
  "testing"
)

//go:embed testdata/data/benchmarks.json
var mockBenchmarksJson []byte

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

func TestVintages(t *testing.T) {
  if testing.Short() {
    t.Skip("skipping in short mode")
  }

  // get results, check for error
  if _, err := Vintages(testBenchmarkId); err != nil {
    t.Fatal(err)
  }
}

//go:embed testdata/data/locations.json
var mockLocationsJson []byte

func TestLocationsFromBenchmark(t *testing.T) {
  if testing.Short() {
    t.Skip("skipping in short mode")
  }

  // decode expected results
  var exp []AddressMatch
  if err := json.Unmarshal(mockLocationsJson, &exp); err != nil {
    t.Fatal(err)
  }

  // get locations, check for error
  _, err := LocationsFromBenchmark(testAddress, testBenchmarkId)
  if err != nil {
    t.Fatal(err)
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

//go:embed testdata/data/geographies.json
var mockGeographiesJson []byte

func TestGeographies(t *testing.T) {
  testAddress := "4600 silver hill rd, 20233"
  testBenchmark := "Public_AR_Census2020"
  testVintage := "Census2010_Census2020"

  if testing.Short() {
    t.Skip("skipping in short mode")
  }

  // get geographies, check for error
  _, err := Geographies(testAddress, testBenchmark, testVintage)
  if err != nil {
    t.Fatal(err)
  }
}

func getBatchInputRows(t *testing.T) []BatchInputRow {
  // open input file
  f, err := os.Open("testdata/data/batch-input.csv")
  if err != nil {
    t.Fatal(err)
  }
  defer f.Close()

  // read rows
  rows, err := NewBatchInputReader(f).ReadAll()
  if err != nil {
    t.Fatal(err)
  }

  // return rows
  return rows
}

func getBatchOutputRows(t *testing.T, path string) []BatchOutputRow {
  // open input file
  f, err := os.Open(path)
  if err != nil {
    t.Fatal(err)
  }
  defer f.Close()

  // read rows
  rows, err := NewBatchOutputReader(f).ReadAll()
  if err != nil {
    t.Fatal(err)
  }

  // return rows
  return rows
}

func TestBatchLocationsFromBenchmark(t *testing.T) {
  if testing.Short() {
    t.Skip("skipping in short mode")
  }

  // get input and expected output
  rows := getBatchInputRows(t)

  // get expected output, build map
  expRows := getBatchOutputRows(t, "testdata/data/batch-output-locations-2020.csv")
  exp := make(map[string]BatchOutputRow)
  for _, row := range(expRows) {
    exp[row.Id] = row
  }

  // send rows, check for error
  gotRows, err := BatchLocationsFromBenchmark(rows, testBenchmarkId)
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

func TestBatchLocations(t *testing.T) {
  if testing.Short() {
    t.Skip("skipping in short mode")
  }

  // get input rows
  rows := getBatchInputRows(t)

  // get expected output, build map
  expRows := getBatchOutputRows(t, "testdata/data/batch-output-locations-2020.csv")
  exp := make(map[string]BatchOutputRow)
  for _, row := range(expRows) {
    exp[row.Id] = row
  }

  // send rows, check for error
  gotRows, err := BatchLocations(rows)
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

func TestBatchGeographies(t *testing.T) {
  if testing.Short() {
    t.Skip("skipping in short mode")
  }

  // get input and expected output
  rows := getBatchInputRows(t)

  // get expected output, build map
  expRows := getBatchOutputRows(t, "testdata/data/batch-output-geographies-2020-2020.csv")
  exp := make(map[string]BatchOutputRow)
  for _, row := range(expRows) {
    exp[row.Id] = row
  }

  // send rows, check for error
  gotRows, err := BatchGeographies(rows, "2020", "2020")
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
