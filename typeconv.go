// Package typeconv provides safe type conversion utilities for Go.
//
// It handles interface{} values commonly encountered when working with JSON,
// configuration files, and external APIs, converting them to concrete Go types
// with clear error messages for unsupported conversions.
package typeconv

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// ToInt converts a value to int. Supported source types:
//   - int, int8, int16, int32, int64
//   - uint, uint8, uint16, uint32, uint64
//   - float32, float64 (only if the value has no fractional part)
//   - string (parsed via strconv.Atoi)
//   - bool (true=1, false=0)
func ToInt(v any) (int, error) {
	if v == nil {
		return 0, fmt.Errorf("typeconv: cannot convert nil to int")
	}

	switch val := v.(type) {
	case int:
		return val, nil
	case int8:
		return int(val), nil
	case int16:
		return int(val), nil
	case int32:
		return int(val), nil
	case int64:
		return int(val), nil
	case uint:
		return int(val), nil
	case uint8:
		return int(val), nil
	case uint16:
		return int(val), nil
	case uint32:
		return int(val), nil
	case uint64:
		return int(val), nil
	case float32:
		if val != float32(int(val)) {
			return 0, fmt.Errorf("typeconv: cannot convert float %v to int without losing precision", val)
		}
		return int(val), nil
	case float64:
		if val != math.Trunc(val) {
			return 0, fmt.Errorf("typeconv: cannot convert float %v to int without losing precision", val)
		}
		return int(val), nil
	case string:
		return strconv.Atoi(val)
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("typeconv: cannot convert %T to int", v)
	}
}

// ToInt64 converts a value to int64. Supported source types:
//   - int, int8, int16, int32, int64
//   - uint, uint8, uint16, uint32, uint64
//   - float32, float64 (only if the value has no fractional part)
//   - string (parsed via strconv.ParseInt)
//   - bool (true=1, false=0)
func ToInt64(v any) (int64, error) {
	if v == nil {
		return 0, fmt.Errorf("typeconv: cannot convert nil to int64")
	}

	switch val := v.(type) {
	case int:
		return int64(val), nil
	case int8:
		return int64(val), nil
	case int16:
		return int64(val), nil
	case int32:
		return int64(val), nil
	case int64:
		return val, nil
	case uint:
		return int64(val), nil
	case uint8:
		return int64(val), nil
	case uint16:
		return int64(val), nil
	case uint32:
		return int64(val), nil
	case uint64:
		return int64(val), nil
	case float32:
		if val != float32(int64(val)) {
			return 0, fmt.Errorf("typeconv: cannot convert float %v to int64 without losing precision", val)
		}
		return int64(val), nil
	case float64:
		if val != math.Trunc(val) {
			return 0, fmt.Errorf("typeconv: cannot convert float %v to int64 without losing precision", val)
		}
		return int64(val), nil
	case string:
		return strconv.ParseInt(val, 10, 64)
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("typeconv: cannot convert %T to int64", v)
	}
}

// ToFloat64 converts a value to float64. Supported source types:
//   - int, int8, int16, int32, int64
//   - uint, uint8, uint16, uint32, uint64
//   - float32, float64
//   - string (parsed via strconv.ParseFloat)
func ToFloat64(v any) (float64, error) {
	if v == nil {
		return 0, fmt.Errorf("typeconv: cannot convert nil to float64")
	}

	switch val := v.(type) {
	case int:
		return float64(val), nil
	case int8:
		return float64(val), nil
	case int16:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case uint:
		return float64(val), nil
	case uint8:
		return float64(val), nil
	case uint16:
		return float64(val), nil
	case uint32:
		return float64(val), nil
	case uint64:
		return float64(val), nil
	case float32:
		return float64(val), nil
	case float64:
		return val, nil
	case string:
		return strconv.ParseFloat(val, 64)
	default:
		return 0, fmt.Errorf("typeconv: cannot convert %T to float64", v)
	}
}

// ToString converts a value to string. Supported source types:
//   - string (pass through)
//   - []byte
//   - fmt.Stringer (calls String())
//   - numeric types and bool (via fmt.Sprint)
func ToString(v any) (string, error) {
	if v == nil {
		return "", fmt.Errorf("typeconv: cannot convert nil to string")
	}

	switch val := v.(type) {
	case string:
		return val, nil
	case []byte:
		return string(val), nil
	case fmt.Stringer:
		return val.String(), nil
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64, bool:
		return fmt.Sprint(val), nil
	default:
		return "", fmt.Errorf("typeconv: cannot convert %T to string", v)
	}
}

// ToBool converts a value to bool. Supported source types:
//   - bool (pass through)
//   - string ("true", "false", "1", "0", "yes", "no" — case insensitive)
//   - int, int8, int16, int32, int64 (0=false, nonzero=true)
//   - uint, uint8, uint16, uint32, uint64 (0=false, nonzero=true)
func ToBool(v any) (bool, error) {
	if v == nil {
		return false, fmt.Errorf("typeconv: cannot convert nil to bool")
	}

	switch val := v.(type) {
	case bool:
		return val, nil
	case string:
		switch strings.ToLower(strings.TrimSpace(val)) {
		case "true", "1", "yes":
			return true, nil
		case "false", "0", "no":
			return false, nil
		default:
			return false, fmt.Errorf("typeconv: cannot convert string %q to bool", val)
		}
	case int:
		return val != 0, nil
	case int8:
		return val != 0, nil
	case int16:
		return val != 0, nil
	case int32:
		return val != 0, nil
	case int64:
		return val != 0, nil
	case uint:
		return val != 0, nil
	case uint8:
		return val != 0, nil
	case uint16:
		return val != 0, nil
	case uint32:
		return val != 0, nil
	case uint64:
		return val != 0, nil
	default:
		return false, fmt.Errorf("typeconv: cannot convert %T to bool", v)
	}
}

// ToDuration converts a value to time.Duration. Supported source types:
//   - time.Duration (pass through)
//   - string (parsed via time.ParseDuration, e.g. "30s", "5m", "1h30m")
//   - int64 (interpreted as nanoseconds)
func ToDuration(v any) (time.Duration, error) {
	if v == nil {
		return 0, fmt.Errorf("typeconv: cannot convert nil to time.Duration")
	}

	switch val := v.(type) {
	case time.Duration:
		return val, nil
	case string:
		return time.ParseDuration(val)
	case int64:
		return time.Duration(val), nil
	default:
		return 0, fmt.Errorf("typeconv: cannot convert %T to time.Duration", v)
	}
}

// ToStringSlice converts a value to []string. Supported source types:
//   - []string (pass through)
//   - []any (each element is converted via ToString)
func ToStringSlice(v any) ([]string, error) {
	if v == nil {
		return nil, fmt.Errorf("typeconv: cannot convert nil to []string")
	}

	switch val := v.(type) {
	case []string:
		return val, nil
	case []any:
		result := make([]string, len(val))
		for i, elem := range val {
			s, err := ToString(elem)
			if err != nil {
				return nil, fmt.Errorf("typeconv: cannot convert element %d in slice: %w", i, err)
			}
			result[i] = s
		}
		return result, nil
	default:
		return nil, fmt.Errorf("typeconv: cannot convert %T to []string", v)
	}
}

// MustInt converts a value to int, panicking if the conversion fails.
func MustInt(v any) int {
	result, err := ToInt(v)
	if err != nil {
		panic(err)
	}
	return result
}

// MustInt64 converts a value to int64, panicking if the conversion fails.
func MustInt64(v any) int64 {
	result, err := ToInt64(v)
	if err != nil {
		panic(err)
	}
	return result
}

// MustFloat64 converts a value to float64, panicking if the conversion fails.
func MustFloat64(v any) float64 {
	result, err := ToFloat64(v)
	if err != nil {
		panic(err)
	}
	return result
}

// MustString converts a value to string, panicking if the conversion fails.
func MustString(v any) string {
	result, err := ToString(v)
	if err != nil {
		panic(err)
	}
	return result
}

// MustBool converts a value to bool, panicking if the conversion fails.
func MustBool(v any) bool {
	result, err := ToBool(v)
	if err != nil {
		panic(err)
	}
	return result
}

// MustDuration converts a value to time.Duration, panicking if the conversion fails.
func MustDuration(v any) time.Duration {
	result, err := ToDuration(v)
	if err != nil {
		panic(err)
	}
	return result
}
