package geocoder

import (
  "encoding/csv"
  "io"
)

// Batch geocode output CSV reader
type BatchOutputReader struct {
  // CSV reader
  r *csv.Reader
}

// Create batch CSV reader.
func NewBatchOutputReader(r io.Reader) BatchOutputReader {
  cr := csv.NewReader(r)
  cr.FieldsPerRecord = -1
  return BatchOutputReader { cr }
}

// Parse CSV rows as BatchOutputRow items.
func (me BatchOutputReader) ReadAll() ([]BatchOutputRow, error) {
  // read CSV rows
  rows, err := me.r.ReadAll()
  if err != nil {
    return []BatchOutputRow{}, err
  }

  // populate result
  r := make([]BatchOutputRow, len(rows))
  for i := range(rows) {
    if outRow, err := NewBatchOutputRow(rows[i]); err != nil {
      return []BatchOutputRow{}, err
    } else {
      r[i] = outRow
    }
  }

  // return result
  return r, nil
}
