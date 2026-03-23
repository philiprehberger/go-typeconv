// Package typeconv provides safe type conversion utilities for Go.
//
// It handles interface{} values commonly encountered when working with JSON,
// configuration files, and external APIs, converting them to concrete Go types
// with clear error messages for unsupported conversions.
package typeconv

import (
	"fmt"
	"math"
	"reflect"
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

// ToIntSlice converts a value to []int. Supported source types:
//   - []int (pass through)
//   - []any (each element is converted via ToInt)
//   - []string (each element is parsed via ToInt)
//   - []float64 (each element is converted via ToInt)
//   - []int64 (each element is cast to int)
func ToIntSlice(v any) ([]int, error) {
	if v == nil {
		return nil, fmt.Errorf("typeconv: cannot convert nil to []int")
	}

	switch val := v.(type) {
	case []int:
		return val, nil
	case []any:
		result := make([]int, len(val))
		for i, elem := range val {
			n, err := ToInt(elem)
			if err != nil {
				return nil, fmt.Errorf("typeconv: cannot convert element %d in slice: %w", i, err)
			}
			result[i] = n
		}
		return result, nil
	case []string:
		result := make([]int, len(val))
		for i, elem := range val {
			n, err := ToInt(elem)
			if err != nil {
				return nil, fmt.Errorf("typeconv: cannot convert element %d in slice: %w", i, err)
			}
			result[i] = n
		}
		return result, nil
	case []float64:
		result := make([]int, len(val))
		for i, elem := range val {
			n, err := ToInt(elem)
			if err != nil {
				return nil, fmt.Errorf("typeconv: cannot convert element %d in slice: %w", i, err)
			}
			result[i] = n
		}
		return result, nil
	case []int64:
		result := make([]int, len(val))
		for i, elem := range val {
			result[i] = int(elem)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("typeconv: cannot convert %T to []int", v)
	}
}

// ToFloat64Slice converts a value to []float64. Supported source types:
//   - []float64 (pass through)
//   - []any (each element is converted via ToFloat64)
//   - []string (each element is parsed via ToFloat64)
//   - []int (each element is cast to float64)
//   - []int64 (each element is cast to float64)
func ToFloat64Slice(v any) ([]float64, error) {
	if v == nil {
		return nil, fmt.Errorf("typeconv: cannot convert nil to []float64")
	}

	switch val := v.(type) {
	case []float64:
		return val, nil
	case []any:
		result := make([]float64, len(val))
		for i, elem := range val {
			f, err := ToFloat64(elem)
			if err != nil {
				return nil, fmt.Errorf("typeconv: cannot convert element %d in slice: %w", i, err)
			}
			result[i] = f
		}
		return result, nil
	case []string:
		result := make([]float64, len(val))
		for i, elem := range val {
			f, err := ToFloat64(elem)
			if err != nil {
				return nil, fmt.Errorf("typeconv: cannot convert element %d in slice: %w", i, err)
			}
			result[i] = f
		}
		return result, nil
	case []int:
		result := make([]float64, len(val))
		for i, elem := range val {
			result[i] = float64(elem)
		}
		return result, nil
	case []int64:
		result := make([]float64, len(val))
		for i, elem := range val {
			result[i] = float64(elem)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("typeconv: cannot convert %T to []float64", v)
	}
}

// ToTime converts a value to time.Time. Supported source types:
//   - time.Time (pass through)
//   - string (parsed as RFC3339, "2006-01-02", or "2006-01-02 15:04:05")
//   - int64 (interpreted as Unix timestamp in seconds)
func ToTime(v any) (time.Time, error) {
	if v == nil {
		return time.Time{}, fmt.Errorf("typeconv: cannot convert nil to time.Time")
	}

	switch val := v.(type) {
	case time.Time:
		return val, nil
	case string:
		if t, err := time.Parse(time.RFC3339, val); err == nil {
			return t, nil
		}
		if t, err := time.Parse("2006-01-02 15:04:05", val); err == nil {
			return t, nil
		}
		if t, err := time.Parse("2006-01-02", val); err == nil {
			return t, nil
		}
		return time.Time{}, fmt.Errorf("typeconv: cannot parse string %q as time (supported: RFC3339, 2006-01-02, 2006-01-02 15:04:05)", val)
	case int64:
		return time.Unix(val, 0), nil
	default:
		return time.Time{}, fmt.Errorf("typeconv: cannot convert %T to time.Time", v)
	}
}

// ToMap converts a value to map[string]any. Supported source types:
//   - map[string]any (pass through)
//   - struct (exported fields become keys, using field names)
func ToMap(v any) (map[string]any, error) {
	if v == nil {
		return nil, fmt.Errorf("typeconv: cannot convert nil to map[string]any")
	}

	switch val := v.(type) {
	case map[string]any:
		return val, nil
	default:
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Ptr {
			if rv.IsNil() {
				return nil, fmt.Errorf("typeconv: cannot convert nil pointer to map[string]any")
			}
			rv = rv.Elem()
		}
		if rv.Kind() != reflect.Struct {
			return nil, fmt.Errorf("typeconv: cannot convert %T to map[string]any", v)
		}
		rt := rv.Type()
		result := make(map[string]any, rt.NumField())
		for i := 0; i < rt.NumField(); i++ {
			field := rt.Field(i)
			if !field.IsExported() {
				continue
			}
			result[field.Name] = rv.Field(i).Interface()
		}
		return result, nil
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

// MustTime converts a value to time.Time, panicking if the conversion fails.
func MustTime(v any) time.Time {
	result, err := ToTime(v)
	if err != nil {
		panic(err)
	}
	return result
}

// MustIntSlice converts a value to []int, panicking if the conversion fails.
func MustIntSlice(v any) []int {
	result, err := ToIntSlice(v)
	if err != nil {
		panic(err)
	}
	return result
}
