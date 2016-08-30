package main

import (
	"log"
	"net"
	"time"

	"database/sql"
	"fmt"

	"github.com/nictuku/dht"
	"github.com/nictuku/roachy/geo"

	_ "github.com/lib/pq"
)

// Not creating a primary key because I don't want unique entries. This means
// cockroachDB creates its own rowid hidden column.
// CREATE TABLE get_peers_log (infohash STRING, node STRING, address STRING, time TIMESTAMPTZ)

func newDBLogger() (*logger, error) {
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/roachy?sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("error connection to the database: %s", err)
	}
	return &logger{db}, nil

}

type logger struct {
	db *sql.DB
}

func (l *logger) GetPeers(addr net.UDPAddr, nodeID string, infoHash dht.InfoHash) {
	node := fmt.Sprintf("%x", nodeID)
	ip := addr.IP.String()
	lat, long := geolookup.LatLong(ip)
	city, country := geolookup.CityCountry(ip)
	if _, err := l.db.Exec(
		"INSERT INTO get_peers_log (infohash, node, address, time, city, country, latitude, longitude) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		infoHash.String(), node, ip, time.Now(), city, country, lat, long); err != nil {
		log.Printf("Error inserting into database: %v", err)
		return
	}
	log.Printf("insert addr %v, infohash %v", ip, infoHash)
	log.Printf("city: %v, country: %v", city, country)
	log.Printf("latitude: %v, longitude: %v", lat, long)
}
