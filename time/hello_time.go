package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now().UTC()
	fmt.Printf("Now: %v\n", now)
	fmt.Printf("Now in UnixNano: %v\n", now.UnixNano())
	today := time.Date(now.Year(), 1, now.YearDay(), 0, 0, 0, 0, time.UTC)
	fmt.Printf("Today: %v\n", today)
	today2 := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	fmt.Printf("Today2: %v\n", today2)
	fmt.Printf("Day of the year: %v", today.YearDay())

	header := uint32(1 << 31)
	header |= uint32(today.YearDay())

	fmt.Printf("header (the first element in the ETT array): %v", header)

	var t time.Time
	fmt.Println("\nZero time:", t)
	fmt.Println("\nZero time UnixNano:", t.UnixNano())
	fmt.Println("\nzero time IsZero:", t.IsZero())

	t2 := time.Unix(0, 0)
	fmt.Println("\nUnix zero time:", t2)
	fmt.Println("\nUnix zero time UnixNano:", t2.UnixNano())
	fmt.Println("\nUnix zero time IsZero:", t2.IsZero())
}
