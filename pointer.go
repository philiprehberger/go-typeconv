package typeconv

// Ptr returns a pointer to the given value.
// This is useful for creating pointers to literals or inline values.
func Ptr[T any](v T) *T {
	return &v
}

// Deref dereferences a pointer, returning the pointed-to value.
// If the pointer is nil, it returns the provided fallback value.
func Deref[T any](p *T, fallback T) T {
	if p == nil {
		return fallback
	}
	return *p
}

// DerefOrZero dereferences a pointer, returning the pointed-to value.
// If the pointer is nil, it returns the zero value for the type.
func DerefOrZero[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}
