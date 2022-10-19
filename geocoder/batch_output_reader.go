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
  return BatchOutputReader { csv.NewReader(r) }
}

// Parse CSV rows as BatchOutputRow items.
func (me BatchOutputReader) ReadAll() ([]BatchOutputRow, error) {
  // read CSV rows
  rows, err := me.r.ReadAll()
  if err != nil {
    return []BatchOutputRow{}, err
  }

  // populate result
  r := make([]BatchOutputRow, 0, len(rows))
  for i, row := range(rows) {
    if outRow, err := NewBatchOutputRow(row); err != nil {
      return []BatchOutputRow{}, err
    } else {
      r[i] = outRow
    }
  }

  // return result
  return r, nil
}
