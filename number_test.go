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
