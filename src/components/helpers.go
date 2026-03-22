package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// GetHTTPStatusCode returns the HTTP status code of the given response.
func GetHTTPStatusCode(response *http.Response) int {
	return response.StatusCode
}

// GetHTTPError returns the HTTP error message of the given response.
func GetHTTPError(response *http.Response) string {
	return response.Status
}

// GetHTTPErrorDetails returns the HTTP error details of the given response.
func GetHTTPErrorDetails(response *http.Response) (string, error) {
	if response == nil {
		return "", fmt.Errorf("nil response")
	}

	if response.Body == nil {
		return "", fmt.Errorf("nil response body")
	}

	defer response.Body.Close()

	var errorDetails struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	err := json.NewDecoder(response.Body).Decode(&errorDetails)
	if err!= nil {
		return "", fmt.Errorf("failed to decode error details: %w", err)
	}

	return fmt.Sprintf("HTTP error %d: %s", errorDetails.Code, errorDetails.Message), nil
}

// GetHTTPBody returns the HTTP response body of the given response.
func GetHTTPBody(response *http.Response) ([]byte, error) {
	if response == nil {
		return nil, fmt.Errorf("nil response")
	}

	if response.Body == nil {
		return nil, fmt.Errorf("nil response body")
	}

	defer response.Body.Close()

	return http.ReadAll(response.Body)
}

// GetHTTPBodyString returns the HTTP response body of the given response as a string.
func GetHTTPBodyString(response *http.Response) (string, error) {
	body, err := GetHTTPBody(response)
	if err!= nil {
		return "", err
	}

	return string(body), nil
}

// GetHTTPBodyJSON returns the HTTP response body of the given response as a JSON object.
func GetHTTPBodyJSON(response *http.Response) (map[string]interface{}, error) {
	body, err := GetHTTPBody(response)
	if err!= nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err!= nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return data, nil
}

// GetHTTPBodyStringArray returns the HTTP response body of the given response as a string array.
func GetHTTPBodyStringArray(response *http.Response) ([]string, error) {
	body, err := GetHTTPBodyString(response)
	if err!= nil {
		return nil, err
	}

	return strings.Split(body, "\n"), nil
}

// GetHTTPBodyJSONArray returns the HTTP response body of the given response as a JSON array.
func GetHTTPBodyJSONArray(response *http.Response) ([]map[string]interface{}, error) {
	data, err := GetHTTPBodyJSON(response)
	if err!= nil {
		return nil, err
	}

	var array []map[string]interface{}
	for _, v := range data {
		array = append(array, v)
	}

	return array, nil
}

// GetHTTPBodyIntArray returns the HTTP response body of the given response as an integer array.
func GetHTTPBodyIntArray(response *http.Response) ([]int, error) {
	data, err := GetHTTPBodyStringArray(response)
	if err!= nil {
		return nil, err
	}

	var array []int
	for _, v := range data {
		var value int
		_, err := fmt.Sscanf(v, "%d", &value)
		if err!= nil {
			return nil, fmt.Errorf("failed to parse int: %w", err)
		}

		array = append(array, value)
	}

	return array, nil
}

// GetHTTPBodyFloat64Array returns the HTTP response body of the given response as a float64 array.
func GetHTTPBodyFloat64Array(response *http.Response) ([]float64, error) {
	data, err := GetHTTPBodyStringArray(response)
	if err!= nil {
		return nil, err
	}

	var array []float64
	for _, v := range data {
		var value float64
		_, err := fmt.Sscanf(v, "%f", &value)
		if err!= nil {
			return nil, fmt.Errorf("failed to parse float64: %w", err)
		}

		array = append(array, value)
	}

	return array, nil
}

// GetHTTPBodyDurationArray returns the HTTP response body of the given response as a duration array.
func GetHTTPBodyDurationArray(response *http.Response) ([]time.Duration, error) {
	data, err := GetHTTPBodyStringArray(response)
	if err!= nil {
		return nil, err
	}

	var array []time.Duration
	for _, v := range data {
		var value time.Duration
		_, err := fmt.Sscanf(v, "%d", &value)
		if err!= nil {
			return nil, fmt.Errorf("failed to parse duration: %w", err)
		}

		array = append(array, value)
	}

	return array, nil
}