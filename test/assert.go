package test

import "testing"

func AssertEqual(t *testing.T, x, y interface{}) {
	if x != y {
		t.Error("Expected ", y, ", got ", x)
	}
}
