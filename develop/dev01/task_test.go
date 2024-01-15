package main

import (
	"testing"
)

func TestGetTime(t *testing.T) {
	time := ntpTime{}
	got, err := time.getTime()
	if err != nil {
		t.Errorf("getTime() error = %v", err)
	}
	if got.IsZero() {
		t.Errorf("getTime() = %v, want non-zero time", got)
	}
}
