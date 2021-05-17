package fdb

import (
	"log"
	"time"
)

const (
	printMetricsInterval = 10 * time.Second
)

var (
	databasesCreated      uint32
	databasesDestroyed    uint32
	transactionsCreated   uint32
	transactionsDestroyed uint32
	futuresCreated        uint32
	futuresDestroyed      uint32
)

func init() {
	go printMetrics()
}

func printMetrics() {
	ticker := time.NewTicker(printMetricsInterval)

	for range ticker.C {
		log.Printf(
			"FoundationDB Metrics:\n"+
				"  databasesCreated: %d, databasesDestroyed: %d\n"+
				"  transactionsCreated: %d, transactionsDestroyed: %d\n"+
				"  futuresCreated: %d, futuresDestroyed: %d",
			databasesCreated, databasesDestroyed,
			transactionsCreated, transactionsDestroyed,
			futuresCreated, futuresDestroyed)
	}
}
