package cron

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var reNumber = regexp.MustCompile(`^\d+$`)

func ExtractValuesInInterval(kind string, data string, minValue, maxValue int) ([]int, error) {
	var values []int
	switch {
	case data == "*":
		for i := range maxValue {
			values = append(values, i)
		}
	case strings.HasPrefix(data, "*/"):
		valueData := data[2:]
		extractedValue, err := getValueInInterval(valueData, minValue, maxValue)
		if err != nil {
			return nil, fmt.Errorf("failed to get %s value: %w", kind, err)
		}
		// start with 0
		values = append(values, minValue)
		var validValue int
		for validValue < maxValue-extractedValue {
			validValue += extractedValue
			values = append(values, validValue)
		}
	case strings.Contains(data, ","):
		for _, valueData := range strings.Split(data, ",") {
			valueData = strings.TrimSpace(valueData)
			extractedValue, err := getValueInInterval(valueData, minValue, maxValue)
			if err != nil {
				return nil, fmt.Errorf("failed to get %s value: %w", kind, err)
			}
			values = append(values, extractedValue)
		}
	case strings.Contains(data, "-"):
		valueData := strings.Split(data, "-")
		if len(valueData) != 2 {
			return nil, fmt.Errorf("%s interval should contain only 2 values: start and end", kind)
		}

		startValueData := strings.TrimSpace(valueData[0])
		startExtractedValue, err := getValueInInterval(startValueData, minValue, maxValue)
		if err != nil {
			return nil, fmt.Errorf("failed to get start %s value: %w", kind, err)
		}
		endValueData := strings.TrimSpace(valueData[1])
		endExtractedValue, err := getValueInInterval(endValueData, minValue, maxValue)
		if err != nil {
			return nil, fmt.Errorf("failed to get end %s value: %w", kind, err)
		}

		if startExtractedValue > endExtractedValue {
			return nil, fmt.Errorf("start %s (%d) cannot be smaller than end minute (%d)", kind, startExtractedValue, endExtractedValue)
		}

		for i := startExtractedValue; i <= endExtractedValue; i++ {
			values = append(values, i)
		}
	case reNumber.MatchString(data):
		value, err := getValueInInterval(data, minValue, maxValue)
		if err != nil {
			return nil, fmt.Errorf("failed to get %s value: %w", kind, err)
		}
		values = append(values, value)
	default:
		return nil, fmt.Errorf("input (%s) is not valid", data)
	}

	return values, nil
}

func getValueInInterval(data string, minValue, maxValue int) (int, error) {
	// check if data is a valid number
	value, err := strconv.Atoi(data)
	if err != nil {
		return 0, fmt.Errorf("data (%s) is not a number", data)
	}
	// check if value is between min and max
	if value < minValue || value > maxValue {
		return 0, fmt.Errorf("value should be between %d-%d. current given value is %d", minValue, maxValue, value)
	}

	return value, nil
}
