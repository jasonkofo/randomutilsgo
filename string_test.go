package randomutilsgo

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateString(t *testing.T) {
	var (
		PWD_LEN = 20
		start   = time.Now()
		pwd     = GenerateRandomString(PWD_LEN)
	)
	fmt.Printf("INFO: Generating random string with %d digits took %f seconds\n", PWD_LEN, time.Since(start).Seconds())
	if len(pwd) != PWD_LEN {
		t.Errorf("FAIL: Did not generate the correct number of digits")
	}
	fmt.Printf("PASS: Generated random string '%s'\n", pwd)
}

func TestIsEmail(t *testing.T) {
	emails := []struct {
		email         string
		ExpectedValue bool
	}{
		{"jason@gmail.com", true},
		{"jason@@gmail.com", false},
		{"person123@hotmail.com", true},
		{"person123hotmail.com", false},
		{"person_123@gmail.com", true},
		{"person_123@gmail.c", false},
		{"", false},
	}

	for _, entry := range emails {
		if IsEmail(entry.email) != entry.ExpectedValue {
			t.Errorf("FAIL: Email %s expected with validity %t\n", entry.email, entry.ExpectedValue)
		}
		fmt.Printf("PASS: Email %s expected with validity %t\n", entry.email, entry.ExpectedValue)
	}
}

func TestIsSimpleString(t *testing.T) {
	emails := []struct {
		value         string
		ExpectedValue bool
	}{
		{"jason@gmail.com", false},
		{"1234567890", true},
		{"abcdefghijklmnopqrstuvwxyz", true},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", true},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890", true},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567 890", false},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghi_jklmnopqrstuvwxyz1234567890", true},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghi_ jklmnopqrstuvwxyz1234567890", false},
	}

	for _, entry := range emails {
		if IsSimpleString(entry.value) != entry.ExpectedValue {
			t.Errorf("FAIL: String %s expected with validity %t\n", entry.value, entry.ExpectedValue)
		}
		fmt.Printf("PASS: String %s expected with validity %t\n", entry.value, entry.ExpectedValue)
	}
}
