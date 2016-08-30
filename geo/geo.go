package geolookup

import (
	"net"
	"os"
	"path"

	geo "github.com/oschwald/geoip2-golang"
)

const (
	loc = "en"
)

var db *geo.Reader

func init() {

	mmdbFile := path.Join(os.Getenv("HOME"), "city.mmdb")
	var err error
	db, err = geo.Open(mmdbFile)
	if err != nil {
		panic(err)
	}
	// defer db.Close()
}

func LatLong(ipAddress string) (latitude, longitude float64) {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return 0, 0
	}
	record, err := db.City(ip)
	if err != nil {
		return 0, 0
	}
	return record.Location.Latitude, record.Location.Longitude
}

func CityCountry(ipAddress string) (city string, country string) {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return "", ""
	}
	record, err := db.City(ip)
	if err != nil {
		return "", ""
	}
	return record.City.Names[loc], record.Country.Names[loc]
}
