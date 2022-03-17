package main

//go test -coverprofile=coverage.out && go tool cover -html=coverage.out

import "testing"

var tests = []struct {
	name string
	dividend float32
	divisor float32
	expected float32
	isErr bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0, 0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
}

func TestDivision (t *testing.T) {
	for _, tt := range tests{
		got, err := Divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("ex[ected %f but got %f", tt.expected, got)
		}
	}
}

// func TestDivide (t *testing.T) {
// 	_, err := divide(10.0, 1.0)

// 	if err != nil {
// 		t.Error("Got an error when we should not have")
// 	}
// }

// func TestBadDivide (t *testing.T) {
// 	_, err := divide(10.0, 0)

// 	if err == nil {
// 		t.Error("Dod not get an error when we should not have")
// 	}
//}