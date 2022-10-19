package geocoder

import (
  "encoding/csv"
  "io"
)

type BatchInputWriter struct {
  w *csv.Writer
}

func NewBatchInputWriter(w io.Writer) BatchInputWriter {
  return BatchInputWriter { csv.NewWriter(w) }
}

func (me BatchInputWriter) WriteAll(rows []BatchInputRow) error {
  for _, row := range(rows) {
    csvRow := []string { row.Id, row.Address, row.City, row.State, row.Zip }
    if err := me.w.Write(csvRow); err != nil {
      return err
    }
  }

  // flush writes
  me.w.Flush()

  return nil
}
