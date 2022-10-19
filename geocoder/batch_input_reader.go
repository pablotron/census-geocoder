package geocoder

import (
  "encoding/csv"
  "io"
)

// Batch geocode CSV reader
type BatchInputReader struct {
  // CSV reader
  r *csv.Reader
}

// Create batch CSV reader.
func NewBatchInputReader(r io.Reader) BatchInputReader {
  cr := csv.NewReader(r)
  cr.FieldsPerRecord = -1
  return BatchInputReader { cr }
}

// Parse all rows in CSV as BatchInputRow items.
//
// Note: The first row of the CSV file is *not* skipped, so if it
// contains column headers it should be removed.
func (me BatchInputReader) ReadAll() ([]BatchInputRow, error) {
  // read CSV rows
  rows, err := me.r.ReadAll()
  if err != nil {
    return []BatchInputRow{}, err
  }

  // populate result
  r := make([]BatchInputRow, len(rows))
  for i, row := range(rows) {
    r[i] = BatchInputRow { row[0], row[1], row[2], row[3], row[4] }
  }

  // return result
  return r, nil
}
