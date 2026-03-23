# Changelog

## 0.2.0

- Add `ToIntSlice` and `ToFloat64Slice` for typed slice conversions
- Add `ToTime` for parsing strings (RFC3339, date, datetime), Unix timestamps, and time.Time
- Add `ToMap` for converting structs and maps to map[string]any
- Add `MustTime` and `MustIntSlice` panicking variants

## 0.1.3

- Consolidate README badges onto single line, fix CHANGELOG format

## 0.1.2

- Add Development section to README

## 0.1.0

- Initial release
- Safe type conversions: int, float, string, bool, duration
- Pointer helpers: Ptr, Deref
- Must-variants for panic-on-error
