package main

import (
	"testing"
)

func TestStateAbbr(t *testing.T) {
	if s := stateAbbr("11"); s != "DC" {
		t.Errorf("State lookup failed")
	}
	if a := stateAbbr("01"); a != "AL" {
		t.Errorf("State lookup failed")
	}
}
