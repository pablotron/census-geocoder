package geocoder

// Row of input batch CSV.
type BatchInputRow struct {
  // Unique row ID (required)
  Id string

  // street address (required)
  Address string

  // city
  City string

  // state
  State string

  // zip code
  Zip string
}
