package common

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

func ParseJson(path string) (string, error) {
	jsonFile, err := os.Open(path)

	if err != nil {
		return "error", err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return "error", err
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return string(byteValue), nil
}

// func ConvertToJSON(input string) (string, error) {
// 	// Remove the curly braces
// 	input = strings.Trim(input, "{}")

// 	// Split the string by comma to get individual key-value pairs
// 	parts := strings.Split(input, ",")

// 	// Map to hold the key-value pairs
// 	data := make(map[string]interface{})

// 	// Regular expression to match keys and values
// 	re := regexp.MustCompile(`\s*([^:]+)\s*:\s*(.+)\s*`)

// 	// Iterate over the parts and extract key-value pairs
// 	for _, part := range parts {
// 		// Find the key-value pairs using regex
// 		matches := re.FindStringSubmatch(part)
// 		if len(matches) != 3 {
// 			return "", fmt.Errorf("invalid input format")
// 		}

// 		key := strings.TrimSpace(matches[1])
// 		value := strings.TrimSpace(matches[2])

// 		// Check if the value is a number
// 		if numValue, err := json.Number(value).Int64(); err == nil {
// 			data[key] = numValue
// 		} else {
// 			// If it's not a number, treat it as a string
// 			data[key] = value
// 		}
// 	}

// 	// Convert the map to JSON
// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		return "", fmt.Errorf("error marshalling JSON: %v", err)
// 	}

// 	return string(jsonData), nil
// }

func CheckIfValid(path string) (string, error) {
	if strings.HasPrefix(path, "{") && strings.HasSuffix(path, "}") {
		r, err := ConvertToJSON(path)
		if err != nil {
			return "error", err
		}

		return r, nil
	}

	r, err := ParseJson(path)

	if err != nil {
		return "error", err
	}

	return r, nil
}
