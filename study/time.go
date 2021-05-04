package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	year := now.Year()
	mon := now.Month()
	//day := now.Day()
	hour := now.Hour()
	fmt.Println(year, " ", mon, " ", hour)
}
