package main

import (
	"fmt"
	"time"
)

func main() {

	// Current local time
	fmt.Println(time.Now())

	// Specific time
	specificTime := time.Date(2025, time.December, 31, 12, 12, 32, 5, time.UTC)
	fmt.Println("ST:", specificTime)

	// Parse time
	parsedTime, _ := time.Parse("2006-01-02", "2025-12-01")       // Mon Jan 2 15:04:05 MST 2006
	parsedTime1, _ := time.Parse("06-01-02", "25-12-01")          // Mon Jan 2 15:04:05 MST 2006
	parsedTime2, _ := time.Parse("06-1-2", "25-12-1")             // Mon Jan 2 15:04:05 MST 2006
	parsedTime3, _ := time.Parse("06-1-2 15-04", "25-12-1 21-59") // Mon Jan 2 15:04:05 MST 2006
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)
	fmt.Println(parsedTime2)
	fmt.Println(parsedTime3)

	// Formatting time
	t := time.Now()
	fmt.Println("FT:", t.Format("Sunday 06-01-02 25-02-22"))

	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday())

	fmt.Println("RT:", t.Round(time.Hour))

	// loc, _ := time.LoadLocation("Asia/Kolkata")
	// t = time.Date(2025, time.December, 8, 14, 16, 40, 00, time.UTC)

	// // Convert this to the specific time zone
	// tLocal := t.In(loc)

	// // Perform rounding
	// roundedTime := t.Round(time.Hour)
	// roundedTimeLocal := roundedTime.In(loc)

	// fmt.Println("Original Time (UTC):", t)
	// fmt.Println("Original Time (Local):", tLocal)
	// fmt.Println("Rounded Time (UTC):", roundedTime)
	// fmt.Println("Rounded Time (Local):", roundedTimeLocal)

	fmt.Println("Truncated Time:", t.Truncate(time.Hour))

	loc, _ := time.LoadLocation("America/New_York")

	// Convert time to location
	tInNY := time.Now().In(loc)
	fmt.Println("New York Time:", tInNY)

	t1 := time.Date(2024, time.December, 4, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, time.December, 4, 18, 0, 0, 0, time.UTC)
	duration := t2.Sub(t1)
	fmt.Println("Duration:", duration)

	// Compare times
	fmt.Println("t2 is after t1?", t2.After(t1))

}
