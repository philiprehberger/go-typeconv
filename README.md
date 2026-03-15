# go-typeconv

Safe type conversion utilities for Go. Handle `interface{}` values from JSON, configs, and APIs.

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

### Must Variants

```go
// When you know the conversion will succeed
port := typeconv.MustInt("8080") // 8080
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
| `MustInt(v any) int` | ToInt or panic |
| `MustInt64(v any) int64` | ToInt64 or panic |
| `MustFloat64(v any) float64` | ToFloat64 or panic |
| `MustString(v any) string` | ToString or panic |
| `MustBool(v any) bool` | ToBool or panic |
| `MustDuration(v any) time.Duration` | ToDuration or panic |
| `Ptr[T any](v T) *T` | Return pointer to value |
| `Deref[T any](p *T, fallback T) T` | Dereference or fallback |
| `DerefOrZero[T any](p *T) T` | Dereference or zero value |

## License

MIT
