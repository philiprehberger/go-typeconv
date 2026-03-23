package typeconv

import (
	"fmt"
	"testing"
	"time"
)

func TestToInt_FromInt(t *testing.T) {
	tests := []struct {
		input any
		want  int
	}{
		{int(42), 42},
		{int8(8), 8},
		{int16(16), 16},
		{int32(32), 32},
		{int64(64), 64},
	}
	for _, tt := range tests {
		got, err := ToInt(tt.input)
		if err != nil {
			t.Errorf("ToInt(%v) returned error: %v", tt.input, err)
		}
		if got != tt.want {
			t.Errorf("ToInt(%v) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestToInt_FromFloat(t *testing.T) {
	got, err := ToInt(float64(8080.0))
	if err != nil {
		t.Fatalf("ToInt(8080.0) returned error: %v", err)
	}
	if got != 8080 {
		t.Errorf("ToInt(8080.0) = %d, want 8080", got)
	}
}

func TestToInt_FromFloat_Fractional(t *testing.T) {
	_, err := ToInt(3.14)
	if err == nil {
		t.Fatal("ToInt(3.14) should return error for fractional float")
	}
}

func TestToInt_FromString(t *testing.T) {
	got, err := ToInt("42")
	if err != nil {
		t.Fatalf("ToInt(\"42\") returned error: %v", err)
	}
	if got != 42 {
		t.Errorf("ToInt(\"42\") = %d, want 42", got)
	}
}

func TestToInt_FromBool(t *testing.T) {
	got, err := ToInt(true)
	if err != nil {
		t.Fatalf("ToInt(true) returned error: %v", err)
	}
	if got != 1 {
		t.Errorf("ToInt(true) = %d, want 1", got)
	}

	got, err = ToInt(false)
	if err != nil {
		t.Fatalf("ToInt(false) returned error: %v", err)
	}
	if got != 0 {
		t.Errorf("ToInt(false) = %d, want 0", got)
	}
}

func TestToInt_Invalid(t *testing.T) {
	_, err := ToInt(struct{}{})
	if err == nil {
		t.Fatal("ToInt(struct{}{}) should return error")
	}
}

func TestToInt_Nil(t *testing.T) {
	_, err := ToInt(nil)
	if err == nil {
		t.Fatal("ToInt(nil) should return error")
	}
}

func TestToInt64(t *testing.T) {
	tests := []struct {
		input any
		want  int64
	}{
		{int(42), 42},
		{int64(999999999999), 999999999999},
		{float64(100.0), 100},
		{"123", 123},
		{true, 1},
		{false, 0},
	}
	for _, tt := range tests {
		got, err := ToInt64(tt.input)
		if err != nil {
			t.Errorf("ToInt64(%v) returned error: %v", tt.input, err)
		}
		if got != tt.want {
			t.Errorf("ToInt64(%v) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestToFloat64_FromInt(t *testing.T) {
	got, err := ToFloat64(42)
	if err != nil {
		t.Fatalf("ToFloat64(42) returned error: %v", err)
	}
	if got != 42.0 {
		t.Errorf("ToFloat64(42) = %f, want 42.0", got)
	}
}

func TestToFloat64_FromString(t *testing.T) {
	got, err := ToFloat64("3.14")
	if err != nil {
		t.Fatalf("ToFloat64(\"3.14\") returned error: %v", err)
	}
	if got != 3.14 {
		t.Errorf("ToFloat64(\"3.14\") = %f, want 3.14", got)
	}
}

func TestToFloat64_Nil(t *testing.T) {
	_, err := ToFloat64(nil)
	if err == nil {
		t.Fatal("ToFloat64(nil) should return error")
	}
}

type testStringer struct {
	val string
}

func (s testStringer) String() string {
	return s.val
}

func TestToString_FromString(t *testing.T) {
	got, err := ToString("hello")
	if err != nil {
		t.Fatalf("ToString(\"hello\") returned error: %v", err)
	}
	if got != "hello" {
		t.Errorf("ToString(\"hello\") = %q, want \"hello\"", got)
	}
}

func TestToString_FromInt(t *testing.T) {
	got, err := ToString(42)
	if err != nil {
		t.Fatalf("ToString(42) returned error: %v", err)
	}
	if got != "42" {
		t.Errorf("ToString(42) = %q, want \"42\"", got)
	}
}

func TestToString_FromStringer(t *testing.T) {
	s := testStringer{val: "custom"}
	got, err := ToString(s)
	if err != nil {
		t.Fatalf("ToString(stringer) returned error: %v", err)
	}
	if got != "custom" {
		t.Errorf("ToString(stringer) = %q, want \"custom\"", got)
	}
}

func TestToString_Nil(t *testing.T) {
	_, err := ToString(nil)
	if err == nil {
		t.Fatal("ToString(nil) should return error")
	}
}

func TestToBool_FromBool(t *testing.T) {
	got, err := ToBool(true)
	if err != nil {
		t.Fatalf("ToBool(true) returned error: %v", err)
	}
	if !got {
		t.Error("ToBool(true) = false, want true")
	}
}

func TestToBool_FromString(t *testing.T) {
	trueStrings := []string{"true", "True", "TRUE", "yes", "Yes", "1"}
	for _, s := range trueStrings {
		got, err := ToBool(s)
		if err != nil {
			t.Errorf("ToBool(%q) returned error: %v", s, err)
		}
		if !got {
			t.Errorf("ToBool(%q) = false, want true", s)
		}
	}

	falseStrings := []string{"false", "False", "FALSE", "no", "No", "0"}
	for _, s := range falseStrings {
		got, err := ToBool(s)
		if err != nil {
			t.Errorf("ToBool(%q) returned error: %v", s, err)
		}
		if got {
			t.Errorf("ToBool(%q) = true, want false", s)
		}
	}
}

func TestToBool_FromInt(t *testing.T) {
	got, err := ToBool(0)
	if err != nil {
		t.Fatalf("ToBool(0) returned error: %v", err)
	}
	if got {
		t.Error("ToBool(0) = true, want false")
	}

	got, err = ToBool(1)
	if err != nil {
		t.Fatalf("ToBool(1) returned error: %v", err)
	}
	if !got {
		t.Error("ToBool(1) = false, want true")
	}

	got, err = ToBool(42)
	if err != nil {
		t.Fatalf("ToBool(42) returned error: %v", err)
	}
	if !got {
		t.Error("ToBool(42) = false, want true")
	}
}

func TestToBool_Invalid(t *testing.T) {
	_, err := ToBool("maybe")
	if err == nil {
		t.Fatal("ToBool(\"maybe\") should return error")
	}
}

func TestToDuration_FromString(t *testing.T) {
	tests := []struct {
		input string
		want  time.Duration
	}{
		{"30s", 30 * time.Second},
		{"5m", 5 * time.Minute},
	}
	for _, tt := range tests {
		got, err := ToDuration(tt.input)
		if err != nil {
			t.Errorf("ToDuration(%q) returned error: %v", tt.input, err)
		}
		if got != tt.want {
			t.Errorf("ToDuration(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestToDuration_FromDuration(t *testing.T) {
	want := 10 * time.Second
	got, err := ToDuration(want)
	if err != nil {
		t.Fatalf("ToDuration(duration) returned error: %v", err)
	}
	if got != want {
		t.Errorf("ToDuration(duration) = %v, want %v", got, want)
	}
}

func TestToDuration_FromInt64(t *testing.T) {
	ns := int64(5 * time.Second)
	got, err := ToDuration(ns)
	if err != nil {
		t.Fatalf("ToDuration(int64) returned error: %v", err)
	}
	if got != 5*time.Second {
		t.Errorf("ToDuration(%d) = %v, want 5s", ns, got)
	}
}

func TestToDuration_Nil(t *testing.T) {
	_, err := ToDuration(nil)
	if err == nil {
		t.Fatal("ToDuration(nil) should return error")
	}
}

func TestToStringSlice_FromStringSlice(t *testing.T) {
	input := []string{"a", "b", "c"}
	got, err := ToStringSlice(input)
	if err != nil {
		t.Fatalf("ToStringSlice([]string) returned error: %v", err)
	}
	if len(got) != 3 || got[0] != "a" || got[1] != "b" || got[2] != "c" {
		t.Errorf("ToStringSlice([]string) = %v, want [a b c]", got)
	}
}

func TestToStringSlice_FromAnySlice(t *testing.T) {
	input := []any{"a", "b"}
	got, err := ToStringSlice(input)
	if err != nil {
		t.Fatalf("ToStringSlice([]any) returned error: %v", err)
	}
	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Errorf("ToStringSlice([]any) = %v, want [a b]", got)
	}
}

func TestToStringSlice_Nil(t *testing.T) {
	_, err := ToStringSlice(nil)
	if err == nil {
		t.Fatal("ToStringSlice(nil) should return error")
	}
}

func TestMustInt_Success(t *testing.T) {
	got := MustInt(42)
	if got != 42 {
		t.Errorf("MustInt(42) = %d, want 42", got)
	}
}

func TestMustInt_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("MustInt(struct{}{}) should panic")
		}
	}()
	MustInt(struct{}{})
}

func TestMustString_Success(t *testing.T) {
	got := MustString("hello")
	if got != "hello" {
		t.Errorf("MustString(\"hello\") = %q, want \"hello\"", got)
	}
}

func TestMustBool_Success(t *testing.T) {
	got := MustBool(true)
	if !got {
		t.Error("MustBool(true) = false, want true")
	}
}

func TestMustFloat64_Success(t *testing.T) {
	got := MustFloat64(3.14)
	if got != 3.14 {
		t.Errorf("MustFloat64(3.14) = %f, want 3.14", got)
	}
}

func TestMustInt64_Success(t *testing.T) {
	got := MustInt64(int64(100))
	if got != 100 {
		t.Errorf("MustInt64(100) = %d, want 100", got)
	}
}

func TestMustDuration_Success(t *testing.T) {
	got := MustDuration("5s")
	if got != 5*time.Second {
		t.Errorf("MustDuration(\"5s\") = %v, want 5s", got)
	}
}

// --- ToIntSlice tests ---

func TestToIntSlice_FromIntSlice(t *testing.T) {
	input := []int{1, 2, 3}
	got, err := ToIntSlice(input)
	if err != nil {
		t.Fatalf("ToIntSlice([]int) returned error: %v", err)
	}
	if len(got) != 3 || got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Errorf("ToIntSlice([]int) = %v, want [1 2 3]", got)
	}
}

func TestToIntSlice_FromAnySlice(t *testing.T) {
	input := []any{1, "2", float64(3.0)}
	got, err := ToIntSlice(input)
	if err != nil {
		t.Fatalf("ToIntSlice([]any) returned error: %v", err)
	}
	if len(got) != 3 || got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Errorf("ToIntSlice([]any) = %v, want [1 2 3]", got)
	}
}

func TestToIntSlice_FromStringSlice(t *testing.T) {
	input := []string{"10", "20", "30"}
	got, err := ToIntSlice(input)
	if err != nil {
		t.Fatalf("ToIntSlice([]string) returned error: %v", err)
	}
	if len(got) != 3 || got[0] != 10 || got[1] != 20 || got[2] != 30 {
		t.Errorf("ToIntSlice([]string) = %v, want [10 20 30]", got)
	}
}

func TestToIntSlice_FromFloat64Slice(t *testing.T) {
	input := []float64{1.0, 2.0, 3.0}
	got, err := ToIntSlice(input)
	if err != nil {
		t.Fatalf("ToIntSlice([]float64) returned error: %v", err)
	}
	if len(got) != 3 || got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Errorf("ToIntSlice([]float64) = %v, want [1 2 3]", got)
	}
}

func TestToIntSlice_FromFloat64Slice_Fractional(t *testing.T) {
	input := []float64{1.0, 2.5}
	_, err := ToIntSlice(input)
	if err == nil {
		t.Fatal("ToIntSlice with fractional float64 should return error")
	}
}

func TestToIntSlice_FromInt64Slice(t *testing.T) {
	input := []int64{100, 200}
	got, err := ToIntSlice(input)
	if err != nil {
		t.Fatalf("ToIntSlice([]int64) returned error: %v", err)
	}
	if len(got) != 2 || got[0] != 100 || got[1] != 200 {
		t.Errorf("ToIntSlice([]int64) = %v, want [100 200]", got)
	}
}

func TestToIntSlice_EmptySlice(t *testing.T) {
	got, err := ToIntSlice([]any{})
	if err != nil {
		t.Fatalf("ToIntSlice([]any{}) returned error: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("ToIntSlice([]any{}) = %v, want []", got)
	}
}

func TestToIntSlice_Nil(t *testing.T) {
	_, err := ToIntSlice(nil)
	if err == nil {
		t.Fatal("ToIntSlice(nil) should return error")
	}
}

func TestToIntSlice_Invalid(t *testing.T) {
	_, err := ToIntSlice("not a slice")
	if err == nil {
		t.Fatal("ToIntSlice(string) should return error")
	}
}

func TestToIntSlice_InvalidElement(t *testing.T) {
	_, err := ToIntSlice([]any{1, "abc"})
	if err == nil {
		t.Fatal("ToIntSlice with unconvertible element should return error")
	}
}

// --- ToFloat64Slice tests ---

func TestToFloat64Slice_FromFloat64Slice(t *testing.T) {
	input := []float64{1.1, 2.2, 3.3}
	got, err := ToFloat64Slice(input)
	if err != nil {
		t.Fatalf("ToFloat64Slice([]float64) returned error: %v", err)
	}
	if len(got) != 3 || got[0] != 1.1 || got[1] != 2.2 || got[2] != 3.3 {
		t.Errorf("ToFloat64Slice([]float64) = %v, want [1.1 2.2 3.3]", got)
	}
}

func TestToFloat64Slice_FromAnySlice(t *testing.T) {
	input := []any{1, "2.5", float64(3.0)}
	got, err := ToFloat64Slice(input)
	if err != nil {
		t.Fatalf("ToFloat64Slice([]any) returned error: %v", err)
	}
	if len(got) != 3 || got[0] != 1.0 || got[1] != 2.5 || got[2] != 3.0 {
		t.Errorf("ToFloat64Slice([]any) = %v, want [1 2.5 3]", got)
	}
}

func TestToFloat64Slice_FromStringSlice(t *testing.T) {
	input := []string{"1.5", "2.5"}
	got, err := ToFloat64Slice(input)
	if err != nil {
		t.Fatalf("ToFloat64Slice([]string) returned error: %v", err)
	}
	if len(got) != 2 || got[0] != 1.5 || got[1] != 2.5 {
		t.Errorf("ToFloat64Slice([]string) = %v, want [1.5 2.5]", got)
	}
}

func TestToFloat64Slice_FromIntSlice(t *testing.T) {
	input := []int{1, 2, 3}
	got, err := ToFloat64Slice(input)
	if err != nil {
		t.Fatalf("ToFloat64Slice([]int) returned error: %v", err)
	}
	if len(got) != 3 || got[0] != 1.0 || got[1] != 2.0 || got[2] != 3.0 {
		t.Errorf("ToFloat64Slice([]int) = %v, want [1 2 3]", got)
	}
}

func TestToFloat64Slice_FromInt64Slice(t *testing.T) {
	input := []int64{10, 20}
	got, err := ToFloat64Slice(input)
	if err != nil {
		t.Fatalf("ToFloat64Slice([]int64) returned error: %v", err)
	}
	if len(got) != 2 || got[0] != 10.0 || got[1] != 20.0 {
		t.Errorf("ToFloat64Slice([]int64) = %v, want [10 20]", got)
	}
}

func TestToFloat64Slice_EmptySlice(t *testing.T) {
	got, err := ToFloat64Slice([]any{})
	if err != nil {
		t.Fatalf("ToFloat64Slice([]any{}) returned error: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("ToFloat64Slice([]any{}) = %v, want []", got)
	}
}

func TestToFloat64Slice_Nil(t *testing.T) {
	_, err := ToFloat64Slice(nil)
	if err == nil {
		t.Fatal("ToFloat64Slice(nil) should return error")
	}
}

func TestToFloat64Slice_Invalid(t *testing.T) {
	_, err := ToFloat64Slice(42)
	if err == nil {
		t.Fatal("ToFloat64Slice(int) should return error")
	}
}

func TestToFloat64Slice_InvalidElement(t *testing.T) {
	_, err := ToFloat64Slice([]any{1.0, "notnum"})
	if err == nil {
		t.Fatal("ToFloat64Slice with unconvertible element should return error")
	}
}

// --- ToTime tests ---

func TestToTime_FromTimeTime(t *testing.T) {
	now := time.Now()
	got, err := ToTime(now)
	if err != nil {
		t.Fatalf("ToTime(time.Time) returned error: %v", err)
	}
	if !got.Equal(now) {
		t.Errorf("ToTime(time.Time) = %v, want %v", got, now)
	}
}

func TestToTime_FromStringRFC3339(t *testing.T) {
	input := "2024-06-15T10:30:00Z"
	got, err := ToTime(input)
	if err != nil {
		t.Fatalf("ToTime(%q) returned error: %v", input, err)
	}
	want := time.Date(2024, 6, 15, 10, 30, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("ToTime(%q) = %v, want %v", input, got, want)
	}
}

func TestToTime_FromStringDateOnly(t *testing.T) {
	input := "2024-06-15"
	got, err := ToTime(input)
	if err != nil {
		t.Fatalf("ToTime(%q) returned error: %v", input, err)
	}
	want := time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("ToTime(%q) = %v, want %v", input, got, want)
	}
}

func TestToTime_FromStringDateTime(t *testing.T) {
	input := "2024-06-15 10:30:00"
	got, err := ToTime(input)
	if err != nil {
		t.Fatalf("ToTime(%q) returned error: %v", input, err)
	}
	want := time.Date(2024, 6, 15, 10, 30, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("ToTime(%q) = %v, want %v", input, got, want)
	}
}

func TestToTime_FromInt64(t *testing.T) {
	ts := int64(1718444400) // 2024-06-15 11:00:00 UTC
	got, err := ToTime(ts)
	if err != nil {
		t.Fatalf("ToTime(int64) returned error: %v", err)
	}
	want := time.Unix(ts, 0)
	if !got.Equal(want) {
		t.Errorf("ToTime(int64) = %v, want %v", got, want)
	}
}

func TestToTime_FromInt64_Zero(t *testing.T) {
	got, err := ToTime(int64(0))
	if err != nil {
		t.Fatalf("ToTime(0) returned error: %v", err)
	}
	want := time.Unix(0, 0)
	if !got.Equal(want) {
		t.Errorf("ToTime(0) = %v, want %v", got, want)
	}
}

func TestToTime_Nil(t *testing.T) {
	_, err := ToTime(nil)
	if err == nil {
		t.Fatal("ToTime(nil) should return error")
	}
}

func TestToTime_InvalidString(t *testing.T) {
	_, err := ToTime("not-a-date")
	if err == nil {
		t.Fatal("ToTime(\"not-a-date\") should return error")
	}
}

func TestToTime_InvalidType(t *testing.T) {
	_, err := ToTime(3.14)
	if err == nil {
		t.Fatal("ToTime(float64) should return error")
	}
}

// --- ToMap tests ---

func TestToMap_FromMapStringAny(t *testing.T) {
	input := map[string]any{"a": 1, "b": "two"}
	got, err := ToMap(input)
	if err != nil {
		t.Fatalf("ToMap(map[string]any) returned error: %v", err)
	}
	if got["a"] != 1 || got["b"] != "two" {
		t.Errorf("ToMap(map[string]any) = %v, want %v", got, input)
	}
}

type testStruct struct {
	Name    string
	Age     int
	hidden  string
}

func TestToMap_FromStruct(t *testing.T) {
	input := testStruct{Name: "Alice", Age: 30, hidden: "secret"}
	got, err := ToMap(input)
	if err != nil {
		t.Fatalf("ToMap(struct) returned error: %v", err)
	}
	if got["Name"] != "Alice" {
		t.Errorf("ToMap(struct)[\"Name\"] = %v, want \"Alice\"", got["Name"])
	}
	if got["Age"] != 30 {
		t.Errorf("ToMap(struct)[\"Age\"] = %v, want 30", got["Age"])
	}
	if _, ok := got["hidden"]; ok {
		t.Error("ToMap(struct) should not include unexported fields")
	}
	if len(got) != 2 {
		t.Errorf("ToMap(struct) has %d keys, want 2", len(got))
	}
}

func TestToMap_FromStructPointer(t *testing.T) {
	input := &testStruct{Name: "Bob", Age: 25}
	got, err := ToMap(input)
	if err != nil {
		t.Fatalf("ToMap(*struct) returned error: %v", err)
	}
	if got["Name"] != "Bob" || got["Age"] != 25 {
		t.Errorf("ToMap(*struct) = %v, unexpected", got)
	}
}

func TestToMap_EmptyStruct(t *testing.T) {
	input := struct{}{}
	got, err := ToMap(input)
	if err != nil {
		t.Fatalf("ToMap(struct{}) returned error: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("ToMap(struct{}) has %d keys, want 0", len(got))
	}
}

func TestToMap_Nil(t *testing.T) {
	_, err := ToMap(nil)
	if err == nil {
		t.Fatal("ToMap(nil) should return error")
	}
}

func TestToMap_NilPointer(t *testing.T) {
	_, err := ToMap((*testStruct)(nil))
	if err == nil {
		t.Fatal("ToMap(nil pointer) should return error")
	}
}

func TestToMap_Invalid(t *testing.T) {
	_, err := ToMap(42)
	if err == nil {
		t.Fatal("ToMap(int) should return error")
	}
}

// --- MustTime tests ---

func TestMustTime_Success(t *testing.T) {
	got := MustTime("2024-06-15T10:30:00Z")
	want := time.Date(2024, 6, 15, 10, 30, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("MustTime() = %v, want %v", got, want)
	}
}

func TestMustTime_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("MustTime(\"invalid\") should panic")
		}
	}()
	MustTime("invalid")
}

// --- MustIntSlice tests ---

func TestMustIntSlice_Success(t *testing.T) {
	got := MustIntSlice([]any{1, 2, 3})
	if len(got) != 3 || got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Errorf("MustIntSlice() = %v, want [1 2 3]", got)
	}
}

func TestMustIntSlice_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("MustIntSlice(nil) should panic")
		}
	}()
	MustIntSlice(nil)
}

// Ensure fmt import is used
var _ = fmt.Sprint
