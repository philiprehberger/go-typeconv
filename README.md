# go-typeconv

[![CI](https://github.com/philiprehberger/go-typeconv/actions/workflows/ci.yml/badge.svg)](https://github.com/philiprehberger/go-typeconv/actions/workflows/ci.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/philiprehberger/go-typeconv.svg)](https://pkg.go.dev/github.com/philiprehberger/go-typeconv) [![License](https://img.shields.io/github/license/philiprehberger/go-typeconv)](LICENSE)

Safe type conversion utilities for Go. Handle `interface{}` values from JSON, configs, and APIs

## Installation

```bash
go get github.com/philiprehberger/go-typeconv
```

## Usage

### Type Conversions

```go
import "github.com/philiprehberger/go-typeconv"

// Working with JSON-decoded data
data := map[string]any{
    "port":    8080.0,
    "debug":   "true",
    "timeout": "30s",
}

port, err := typeconv.ToInt(data["port"])       // 8080, nil
debug, err := typeconv.ToBool(data["debug"])    // true, nil
timeout, err := typeconv.ToDuration(data["timeout"]) // 30s, nil
```

### Pointer Helpers

```go
// Create pointers to literals
name := typeconv.Ptr("Alice")   // *string

// Safe dereference with fallback
val := typeconv.Deref(name, "unknown")   // "Alice"
val = typeconv.Deref((*string)(nil), "unknown") // "unknown"

// Safe dereference with zero value
val = typeconv.DerefOrZero((*string)(nil)) // ""
```

### Slice Conversions

```go
// Convert mixed slices from JSON to typed slices
ids, err := typeconv.ToIntSlice([]any{1, "2", 3.0})     // [1, 2, 3], nil
prices, err := typeconv.ToFloat64Slice([]string{"1.5", "2.5"}) // [1.5, 2.5], nil

// Also accepts []int, []float64, []int64, []string directly
vals, err := typeconv.ToIntSlice([]float64{10.0, 20.0})  // [10, 20], nil
```

### Time Parsing

```go
// Parse from RFC3339, date, or datetime strings
t1, err := typeconv.ToTime("2024-06-15T10:30:00Z")    // RFC3339
t2, err := typeconv.ToTime("2024-06-15")               // date only
t3, err := typeconv.ToTime("2024-06-15 10:30:00")      // datetime

// From Unix timestamp (seconds)
t4, err := typeconv.ToTime(int64(1718444400))

// Pass through time.Time
t5, err := typeconv.ToTime(time.Now())
```

### Struct to Map

```go
type User struct {
    Name  string
    Age   int
}

m, err := typeconv.ToMap(User{Name: "Alice", Age: 30})
// map[string]any{"Name": "Alice", "Age": 30}

// Also accepts map[string]any (pass through) and struct pointers
```

### Must Variants

```go
// When you know the conversion will succeed
port := typeconv.MustInt("8080") // 8080
t := typeconv.MustTime("2024-06-15T10:30:00Z")
ids := typeconv.MustIntSlice([]any{1, 2, 3})
// Panics on failure — use only when input is trusted
```

## API

| Function | Description |
|----------|-------------|
| `ToInt(v any) (int, error)` | Convert to int from int/float/string/bool |
| `ToInt64(v any) (int64, error)` | Convert to int64 |
| `ToFloat64(v any) (float64, error)` | Convert to float64 from numeric/string |
| `ToString(v any) (string, error)` | Convert to string from string/[]byte/Stringer/numeric |
| `ToBool(v any) (bool, error)` | Convert to bool from bool/string/int |
| `ToDuration(v any) (time.Duration, error)` | Convert to Duration from Duration/string/int64 |
| `ToStringSlice(v any) ([]string, error)` | Convert to []string from []string/[]any |
| `ToIntSlice(v any) ([]int, error)` | Convert to []int from []int/[]any/[]string/[]float64/[]int64 |
| `ToFloat64Slice(v any) ([]float64, error)` | Convert to []float64 from []float64/[]any/[]string/[]int/[]int64 |
| `ToTime(v any) (time.Time, error)` | Convert to time.Time from string/int64/time.Time |
| `ToMap(v any) (map[string]any, error)` | Convert struct or map[string]any to map[string]any |
| `MustInt(v any) int` | ToInt or panic |
| `MustInt64(v any) int64` | ToInt64 or panic |
| `MustFloat64(v any) float64` | ToFloat64 or panic |
| `MustString(v any) string` | ToString or panic |
| `MustBool(v any) bool` | ToBool or panic |
| `MustDuration(v any) time.Duration` | ToDuration or panic |
| `MustTime(v any) time.Time` | ToTime or panic |
| `MustIntSlice(v any) []int` | ToIntSlice or panic |
| `Ptr[T any](v T) *T` | Return pointer to value |
| `Deref[T any](p *T, fallback T) T` | Dereference or fallback |
| `DerefOrZero[T any](p *T) T` | Dereference or zero value |

## Development

```bash
go test ./...
go vet ./...
```

## License

MIT
