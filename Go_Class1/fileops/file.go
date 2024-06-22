package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetStringFloatFromFile(fileName string) (float64, error) {
	data, error := os.ReadFile(fileName)
	if error != nil {
		return 1000, errors.New("failed to find file")
	}

	valueString := string(data)
	value, error := strconv.ParseFloat(valueString, 64)

	if error != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueString := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueString), 0644)
}