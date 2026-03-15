package typeconv

import "testing"

func TestPtr(t *testing.T) {
	p := Ptr(42)
	if *p != 42 {
		t.Errorf("Ptr(42) = %d, want 42", *p)
	}

	s := Ptr("hello")
	if *s != "hello" {
		t.Errorf("Ptr(\"hello\") = %q, want \"hello\"", *s)
	}
}

func TestDeref_NonNil(t *testing.T) {
	v := 42
	got := Deref(&v, 0)
	if got != 42 {
		t.Errorf("Deref(&42, 0) = %d, want 42", got)
	}
}

func TestDeref_Nil(t *testing.T) {
	got := Deref((*int)(nil), 99)
	if got != 99 {
		t.Errorf("Deref(nil, 99) = %d, want 99", got)
	}
}

func TestDerefOrZero_NonNil(t *testing.T) {
	v := "hello"
	got := DerefOrZero(&v)
	if got != "hello" {
		t.Errorf("DerefOrZero(&\"hello\") = %q, want \"hello\"", got)
	}
}

func TestDerefOrZero_Nil(t *testing.T) {
	got := DerefOrZero((*int)(nil))
	if got != 0 {
		t.Errorf("DerefOrZero(nil) = %d, want 0", got)
	}

	gotStr := DerefOrZero((*string)(nil))
	if gotStr != "" {
		t.Errorf("DerefOrZero(nil) = %q, want \"\"", gotStr)
	}
}
