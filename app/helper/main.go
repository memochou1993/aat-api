package helper

import (
	"log"
	"time"
)

// Measure measures the execution time.
func Measure(start time.Time, name string) {
	log.Printf("%s took %s", name, time.Since(start))
}
