package main

import (
	"log"
	"net"
	"time"

	"database/sql"
	"fmt"

	"github.com/nictuku/dht"

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
	if _, err := l.db.Exec(
		"INSERT INTO get_peers_log (infohash, node, address, time) VALUES ($1, $2, $3, $4))", infoHash.String(), nodeID, addr.String(), time.Now()); err != nil {
		log.Printf("Error inserting into database: %v", err)
	}
	log.Printf("insert addr %v, infohash %x", addr.String(), infoHash)
}
