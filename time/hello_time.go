package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now().UTC()
	fmt.Printf("Now: %v\n", now)
	today := time.Date(now.Year(), 1, now.YearDay(), 0, 0, 0, 0, time.UTC)
	fmt.Printf("Today: %v\n", today)
	today2 := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	fmt.Printf("Today2: %v\n", today2)
	fmt.Printf("Day of the year: %v", today.YearDay())

	header := uint32(1 << 31)
	header |= uint32(today.YearDay())

	fmt.Printf("header (the first element in the ETT array): %v", header)
}
