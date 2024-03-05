package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var reNumber = regexp.MustCompile(`^\d+$`)

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

	minutes, err := extractMinutes(minuteData)
	if err != nil {
		fmt.Printf("failed to extract minutes from input: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(minutes)
}

func extractMinutes(data string) ([]int, error) {
	var minutes []int
	switch {
	case data == "*":
		for i := range 60 {
			minutes = append(minutes, i)
		}
	case strings.HasPrefix(data, "*/"):
		minuteData := data[2:]
		minuteValue, err := getMinuteValue(minuteData)
		if err != nil {
			return nil, fmt.Errorf("failed to get minute value: %w", err)
		}
		// start with 0
		minutes = append(minutes, 0)
		var validMinute int
		for validMinute < 60-minuteValue {
			validMinute += minuteValue
			minutes = append(minutes, validMinute)
		}
	case strings.Contains(data, ","):
		for _, minuteData := range strings.Split(data, ",") {
			minuteData = strings.TrimSpace(minuteData)
			minuteValue, err := getMinuteValue(minuteData)
			if err != nil {
				return nil, fmt.Errorf("failed to get minute value: %w", err)
			}
			minutes = append(minutes, minuteValue)
		}
	case strings.Contains(data, "-"):
		minuteData := strings.Split(data, "-")
		if len(minuteData) != 2 {
			return nil, fmt.Errorf("minute interval should contain only 2 values: start and end")
		}

		startMinuteData := strings.TrimSpace(minuteData[0])
		startMinuteValue, err := getMinuteValue(startMinuteData)
		if err != nil {
			return nil, fmt.Errorf("failed to get start minute value: %w", err)
		}
		endMinuteData := strings.TrimSpace(minuteData[1])
		endMinuteValue, err := getMinuteValue(endMinuteData)
		if err != nil {
			return nil, fmt.Errorf("failed to get end minute value: %w", err)
		}

		if startMinuteValue > endMinuteValue {
			return nil, fmt.Errorf("start minute (%d) cannot be smaller than end minute (%d)", startMinuteValue, endMinuteValue)
		}

		for i := startMinuteValue; i <= endMinuteValue; i++ {
			minutes = append(minutes, i)
		}
	case reNumber.MatchString(data):
		value, err := getMinuteValue(data)
		if err != nil {
			return nil, fmt.Errorf("failed to get minute value: %w", err)
		}
		minutes = append(minutes, value)
	default:
		return nil, fmt.Errorf("input (%s) is not valid", data)
	}

	return minutes, nil
}

func getMinuteValue(data string) (int, error) {
	// check if minute data is a valid number
	value, err := strconv.Atoi(data)
	if err != nil {
		return 0, fmt.Errorf("minute data (%s) is not a number", data)
	}
	// check if minute value is between 0 and 59
	if value < 0 || value > 59 {
		return 0, fmt.Errorf("minute value should be between 0-59. current given value is %d", value)
	}

	return value, nil
}
