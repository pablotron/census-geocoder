package geocoder

import (
  "fmt"
  "math"
)

// Batch geocoder output row.
type BatchOutputRow struct {
  // unique row ID
  Id string `json:"id"`

  // source address
  InputAddress string `json:"input_address"`

  // was this input row matched?
  Match bool `json:"is_match"`

  // is address an exact match?
  Exact bool `json:"is_exact"`

  // normalized matched address
  MatchAddress string `json:"match_address"`

  // lat/long
  Coordinates Coordinates `json:"coordinates"`

  TigerLine TigerLine `json:"tigerLine"`
}

// Create batch output row from CSV row.
func NewBatchOutputRow(row []string) (BatchOutputRow, error) {
  if len(row) < 3 {
    return BatchOutputRow{}, fmt.Errorf("invalid batch output row: %#v", row)
  }

  match := (row[2] == "Match")
  exact := (row[2] == "Match") && (len(row) > 3) && (row[3] == "Exact")
  matchAddress := ""
  var matchCoords Coordinates
  var matchLine TigerLine
  if match {
    matchAddress = row[4]

    if tmpCoords, err := NewCoordinates(row[5]); err != nil {
      return BatchOutputRow{}, err
    } else {
      matchCoords = tmpCoords
    }

    matchLine = TigerLine { row[6], row[7] }
  }

  return BatchOutputRow {
    Id: row[0],
    InputAddress: row[1],
    Match: match,
    Exact: exact,
    MatchAddress: matchAddress,
    Coordinates: matchCoords,
    TigerLine: matchLine,
  }, nil
}

// Compare two batch output rows and return true if they are equal.
//
// Note: the coordinates values compare as equal if they are less than
// 0.001 degrees apart.
func compareBatchOutputRow(a, b BatchOutputRow) bool {
  return a.Id == b.Id &&
         a.InputAddress == b.InputAddress &&
         a.Match == b.Match &&
         a.Exact == b.Exact &&
         a.MatchAddress == b.MatchAddress &&
         math.Abs(a.Coordinates.X - b.Coordinates.X) < 0.001 &&
         math.Abs(a.Coordinates.Y - b.Coordinates.Y) < 0.001 &&
         a.TigerLine.Id == b.TigerLine.Id &&
         a.TigerLine.Side == b.TigerLine.Side
}
