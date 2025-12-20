package model

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"strings"
)

// Int32Array is a custom type for handling PostgreSQL integer arrays
type Int32Array []int32

// Scan implements the Scanner interface for database deserialization
func (a *Int32Array) Scan(value interface{}) error {
	if value == nil {
		*a = nil
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return nil
	}

	// Remove curly braces
	str = strings.Trim(str, "{}")
	
	// Handle empty array
	if str == "" {
		*a = []int32{}
		return nil
	}

	// Split by comma
	parts := strings.Split(str, ",")
	result := make([]int32, len(parts))
	
	for i, part := range parts {
		// Trim whitespace
		part = strings.TrimSpace(part)
		// Convert to int32
		val, err := strconv.ParseInt(part, 10, 32)
		if err != nil {
			return err
		}
		result[i] = int32(val)
	}
	
	*a = result
	return nil
}

// Value implements the Valuer interface for database serialization
func (a Int32Array) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	
	strs := make([]string, len(a))
	for i, v := range a {
		strs[i] = strconv.FormatInt(int64(v), 10)
	}
	
	return "{" + strings.Join(strs, ",") + "}", nil
}

// MarshalJSON implements json.Marshaler interface
func (a Int32Array) MarshalJSON() ([]byte, error) {
	return json.Marshal([]int32(a))
}

// UnmarshalJSON implements json.Unmarshaler interface
func (a *Int32Array) UnmarshalJSON(data []byte) error {
	var arr []int32
	if err := json.Unmarshal(data, &arr); err != nil {
		return err
	}
	*a = Int32Array(arr)
	return nil
}