package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("failed to read file")
	}

	valueStr := string(data)
	value, err := strconv.ParseFloat(valueStr, 64)

	if err != nil {
		return defaultValue, errors.New("failed to parse string")
	}

	return value, nil
}

func WriteFloatToFile(fileName string, value float64) {
	valueStr := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueStr), 0644)
}
