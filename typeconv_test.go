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

// Ensure fmt import is used
var _ = fmt.Sprint
