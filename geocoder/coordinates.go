package geocoder

import (
  "fmt"
  "strconv"
  "strings"
)

// lat/long coordinates.
type Coordinates struct {
  X float64 `json:"x"`
  Y float64 `json:"y"`
}

// Create coordinates from input string
func NewCoordinates(s string) (Coordinates, error) {
  xs, ys, found := strings.Cut(s, ",")
  if !found {
    err := fmt.Errorf("invalid coordinates: %s", s)
    return Coordinates{}, err
  }

  // parse X coordinate
  x, err := strconv.ParseFloat(xs, 64)
  if err != nil {
    return Coordinates{}, err
  }

  // parse Y coordinate
  y, err := strconv.ParseFloat(ys, 64)
  if err != nil {
    return Coordinates{}, err
  }

  return Coordinates { x, y }, nil
}
