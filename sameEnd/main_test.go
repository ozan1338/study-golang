package main

import "testing"

var test = []struct {
	name string
	inputStr string
	expectedValue string
}{
	{"testing-data-1","abcd cbdd",""},
	{"testing-data-2","weiwei","wei"},
	{"testing-data-3","gongxi gongxi","gongxi"},
}

func TestSameEnd (t *testing.T) {
	for _, tt := range test {
		got := sameEnd(tt.inputStr)
		
		if got != tt.expectedValue {
			t.Errorf("expected %s but got %s", tt.expectedValue, got)
		}
	}
}

