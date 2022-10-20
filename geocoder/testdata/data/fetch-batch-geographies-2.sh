#!/bin/sh

# use curl to fetch batch geographies
# ref: https://geocoding.geo.census.gov/geocoder/Geocoding_Services_API.html

exec curl -v --form addressFile=@batch-input.csv --form benchmark=4 --form vintage=4 https://geocoding.geo.census.gov/geocoder/geographies/addressbatch --output output-geographies-2.csv
