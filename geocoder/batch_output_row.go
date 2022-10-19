package geocoder

import (
  "fmt"
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
  Coordinates struct {
    X float64 `json:"x"` // longitude
    Y float64 `json:"y"` // latitude
  } `json:"coordinates"`

  TigerLine struct {
    Id string `json:"tigerLineId"`
    Side string `json:"side"`
  } `json:"tigerLine"`
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
