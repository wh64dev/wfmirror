package util

import (
	"encoding/json"
	"os"
)

func ParseJSON[T any](filename string) (*T, error) {
	file, _ := os.Open(filename)
	defer file.Close()

	decoder := json.NewDecoder(file)
	var data T

	err := decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
