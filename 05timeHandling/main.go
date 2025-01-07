package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Understanding Time package")

	presentTime := time.Now()
	fmt.Println("present time:", presentTime)

	dateOnly := presentTime.Format(time.DateOnly)
	fmt.Println("present date:", dateOnly)

	timeOnly := presentTime.Format(time.TimeOnly)
	fmt.Println("present time:", timeOnly)

	datetime := presentTime.Format(time.DateTime)
	fmt.Println("present Date-Time:", datetime)

	weekday := presentTime.Format("Monday")
	fmt.Println("weekday:", weekday)

	wd := time.Now().Weekday()
	fmt.Println("now day (wd):", wd)

	newDate := time.Date(2025, time.October, 24, 22, 15, 5, 0, time.Local)
	fmt.Println("newDate:", newDate.Format(time.DateTime+" Monday"), "Happy Birthday!!")
}
