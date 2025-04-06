package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadFromJSON is a generic function that loads data from a JSON file
// into any structure that matches the JSON schema.
func LoadFromJSON[T any](filePath string) (T, error) {
	var result T

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return result, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Create a new decoder and decode into the generic type
	if err := json.NewDecoder(file).Decode(&result); err != nil {
		return result, fmt.Errorf("error parsing JSON: %w", err)
	}

	return result, nil
}

// SaveToJSON is a generic function that saves any data structure
// to a JSON file, overwriting its content completely.
func SaveToJSON[T any](filePath string, data T) error {
	// Create or truncate the file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating/truncating file: %w", err)
	}
	defer file.Close()

	// Create an encoder and write the JSON
	if err := json.NewEncoder(file).Encode(data); err != nil {
		return fmt.Errorf("error encoding JSON: %w", err)
	}

	return nil
}
