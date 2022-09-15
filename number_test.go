package randomutilsgo

import (
	"fmt"
	"testing"
)

func TestIsInteger(t *testing.T) {
	emails := []struct {
		integer       interface{}
		ExpectedValue bool
	}{
		{1, true},
		{1.0, false},
		{1.52, false},
		{int64(1), true},
		{int32(1), true},
		{int16(1), true},
		{int8(1), true},
		{int32(1), true},
		{float32(1), false},
		{"", false},
		{map[string]interface{}{}, false},
		{'d', true},
	}

	for _, entry := range emails {
		if IsInteger(entry.integer) != entry.ExpectedValue {
			t.Errorf("FAIL: Integer %v expected with validity %t\n", entry.integer, entry.ExpectedValue)
		}
		fmt.Printf("PASS: Integer %v expected with validity %t\n", entry.integer, entry.ExpectedValue)
	}
}

func TestConvertToInt(t *testing.T) {
	testcases := []struct {
		integer       interface{}
		expectedValue int
	}{
		{1, 1},
		{1.0, 0},
		{1.52, 0},
		{int64(100), int(100)},
		{int32(100), int(100)},
		{int16(10), 10},
		{int8(10), 10},
		{int32(1), 1},
		{float32(1), 0},
		{"", 0},
		{map[string]interface{}{}, 0},
		{'d', 100},
	}

	for _, entry := range testcases {
		result := ConvertToInt(entry.integer)
		if result != entry.expectedValue {
			t.Errorf("FAIL: Integer %v expected with validity %d\n", entry.integer, entry.expectedValue)
		}
		fmt.Printf("PASS: Integer %v expected with validity %d\n", entry.integer, entry.expectedValue)
	}
}
