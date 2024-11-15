package main

import (
	"fmt"
	"time"
)

func main() {
	var scores [3]int
	scores[1] = 90
	// fmt.Println(scores[1], scores[0]) //, scores[3] // invalid argument: index 3 out bounds
	// fmt.Printf("%#v\n", scores)
	// fmt.Println(scores)

	// for i := 0; i <= len(scores); i++ {
	for i := 0; i < len(scores); i++ { // unsafe
		fmt.Printf("%d ", scores[i])
	}

	// var dates [3]time.Timex
	// dates[1] = time.Unix(1947920000, 0)
	// fmt.Println(dates[1])

	// var dates [3]time.Time = [3]time.Time{time.Unix(0, 0), time.Unix(0, 1), time.Unix(194720000, 0)}
	// fmt.Println(dates[0], dates[1], dates[2])

	// dates := [3]time.Time{time.Unix(0, 0), time.Unix(0, 1), time.Unix(194720000, 0)}
	// fmt.Println(dates[0], dates[1], dates[2])

	// dates := [3]time.Time{
	// 	time.Unix(0, 0),
	// 	time.Unix(0, 1),
	// 	time.Unix(194720000, 0), // need comma
	// }
	// fmt.Println(dates[0], dates[1], dates[2])

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(0, 1),
		time.Unix(194720000, 0)}
	// fmt.Println(dates[0], dates[1], dates[2])
	// fmt.Printf("%#v\n", dates)
	// fmt.Println(dates)

	fmt.Println()
	// for i, date := range dates {
	// 	fmt.Println(i, date)
	// }
	for _, date := range dates { // like python style, SAFE!
		fmt.Println(date)
	}
}
