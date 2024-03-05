package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/iulianclita/cron-parser/cron"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("expected exactly one command line argument")
		os.Exit(1)
	}

	cronData := strings.Split(os.Args[1], " ")
	if len(cronData) != 6 {
		fmt.Println("command line should receive exactly 6 arguments: minute, hour, day of month, month, day of week and command")
		os.Exit(1)
	}

	minuteData := strings.TrimSpace(cronData[0])
	hourData := strings.TrimSpace(cronData[1])
	dayOfMonthData := strings.TrimSpace(cronData[2])
	monthData := strings.TrimSpace(cronData[3])
	dayOfMonthWeek := strings.TrimSpace(cronData[4])
	command := strings.TrimSpace(cronData[5])

	fmt.Println(minuteData)
	fmt.Println(hourData)
	fmt.Println(dayOfMonthData)
	fmt.Println(monthData)
	fmt.Println(dayOfMonthWeek)
	fmt.Println(command)

	minutes, err := cron.ExtractMinutes(minuteData)
	if err != nil {
		fmt.Printf("failed to extract minutes from input: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(minutes)
}
