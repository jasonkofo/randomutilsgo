package randomutilsgo

import (
	"fmt"
	"testing"
	"time"
)

func TestGeneratePassword(t *testing.T) {
	var (
		PWD_LEN = 100
		start   = time.Now()
		pwd     = GenerateRandomString(PWD_LEN)
	)
	fmt.Printf("Generating password with %d digits took %f seconds\n", PWD_LEN, time.Since(start).Seconds())
	if len(pwd) != PWD_LEN {
		t.Errorf("Did not generate the correct number of digits")
	}
}
