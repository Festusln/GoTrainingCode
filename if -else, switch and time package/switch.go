package main

import (
	"fmt"
	"time"
)

func dayName() {
	today := time.Now().Weekday()
	if today == time.Friday {
		fmt.Println("Thank God it's Friday")
	}
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today is saturday")
	case today + 1:
		fmt.Println("Tomorrow is saturday")
	case today + 2:
		fmt.Println("Saturday is in two days time")
	default:
		fmt.Println("Saturday is more than two days away")
	}
}
