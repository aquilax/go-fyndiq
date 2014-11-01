package fyndiq

import "testing"

func TestNewFyndiqApi(t *testing.T) {
	fapi := NewFyndiqApi("user", "token");
	if fapi == nil {
		t.Errorf("Expected NewFyndiqApi return pointer")
	}
}
