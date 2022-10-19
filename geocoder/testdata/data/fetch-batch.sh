#!/bin/sh
# ref: https://geocoding.geo.census.gov/geocoder/Geocoding_Services_API.html
exec curl -v --form addressFile=@batch-input.csv --form benchmark=2020  https://geocoding.geo.census.gov/geocoder/locations/addressbatch --output output.csv
