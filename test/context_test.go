package test

import "testing"

func TestNewContext(t *testing.T) {
	ctx := NewContext("foo")
	if ctx.URL != "foo" {
		t.Fatal()
	}
}
