package main

import (
	"fmt"
	"os"
	"strconv"
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
	dayOfWeekData := strings.TrimSpace(cronData[4])
	command := strings.TrimSpace(cronData[5])

	extractMinutes := func(data string) ([]int, error) {
		return cron.ExtractValuesInInterval("minute", data, 0, 59)
	}
	minutes, err := extractMinutes(minuteData)
	if err != nil {
		fmt.Printf("failed to extract minutes from input: %v\n", err)
		os.Exit(1)
	}

	extractHours := func(data string) ([]int, error) {
		return cron.ExtractValuesInInterval("hour", data, 0, 23)
	}
	hours, err := extractHours(hourData)
	if err != nil {
		fmt.Printf("failed to extract hours from input: %v\n", err)
		os.Exit(1)
	}

	extractDaysOfMonth := func(data string) ([]int, error) {
		// TODO: This may need extra checkup and correlated with month values since not all months have 31 days
		return cron.ExtractValuesInInterval("day of month", data, 1, 31)
	}
	daysOfMonth, err := extractDaysOfMonth(dayOfMonthData)
	if err != nil {
		fmt.Printf("failed to extract days of month from input: %v\n", err)
		os.Exit(1)
	}

	extractMonths := func(data string) ([]int, error) {
		// TODO: This may need extra checkup and correlated with day of month values since not all months have 31 days
		return cron.ExtractValuesInInterval("month", data, 1, 12)
	}
	months, err := extractMonths(monthData)
	if err != nil {
		fmt.Printf("failed to extract months from input: %v\n", err)
		os.Exit(1)
	}

	extractDaysOfWeek := func(data string) ([]int, error) {
		return cron.ExtractValuesInInterval("day of week", data, 1, 7)
	}
	daysOfWeek, err := extractDaysOfWeek(dayOfWeekData)
	if err != nil {
		fmt.Printf("failed to extract days of week from input: %v\n", err)
		os.Exit(1)
	}

	displayTable(minutes, hours, daysOfMonth, months, daysOfWeek, command)
}

func displayTable(minutes, hours, daysOfMonth, months, daysOfWeek []int, command string) {
	minuteText := makeText(minutes)
	hourText := makeText(hours)
	dayOfMonthText := makeText(daysOfMonth)
	monthText := makeText(months)
	dayOfWeekText := makeText(daysOfWeek)

	minuteLine := "minute" + strings.Repeat(" ", 14-len("minute")) + minuteText
	hourLine := "hour" + strings.Repeat(" ", 14-len("hour")) + hourText
	dayOfMonthLine := "day of month" + strings.Repeat(" ", 14-len("day of month")) + dayOfMonthText
	monthLine := "month" + strings.Repeat(" ", 14-len("month")) + monthText
	dayOfWeekLine := "day of week" + strings.Repeat(" ", 14-len("day of week")) + dayOfWeekText
	commandLine := "command" + strings.Repeat(" ", 14-len("command")) + command

	fmt.Printf("%s\n%s\n%s\n%s\n%s\n%s\n",
		minuteLine, hourLine, dayOfMonthLine, monthLine, dayOfWeekLine, commandLine,
	)
}

func makeText(values []int) string {
	strValues := make([]string, len(values))
	for i, v := range values {
		strValues[i] = strconv.Itoa(v)
	}

	return strings.Join(strValues, " ")
}
