package randomutilsgo

import (
	"math"
	"math/rand"
	"regexp"
	"time"
)

var upperCaseRune []rune = []rune{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L', 'M',
	'N', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

var lowerCaseRune []rune = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j', 'k', 'm', 'n',
	'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}

var numberRune []rune = []rune{
	'1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
}

var symbolRune []rune = []rune{
	'_', '+', '@', '$', '%', '*', '#',
}

var runeOfRunes = [][]rune{
	upperCaseRune, lowerCaseRune, numberRune, symbolRune,
}

func floorAndMinimum(x int, limit int) int {
	y := math.Floor(float64(x))
	return int(math.Min(y, float64(limit)))
}

func randomIntegerInBounds(max int) int {
	rand.Seed(time.Now().UnixNano())
	return floorAndMinimum(rand.Intn(max), max-1)
}

func GenerateRandomString(length int) string {
	str := ""
	for i := 0; i < length; i++ {
		var (
			arrlen   = len(runeOfRunes)
			jdx      = randomIntegerInBounds(arrlen)
			arrinlen = len(runeOfRunes[jdx])
			idx      = randomIntegerInBounds(arrinlen)
		)
		str += string(runeOfRunes[jdx][idx])
	}
	return str
}

func GenerateRandomStringForUseAsPassword() string {
	str := ""
	uci := randomIntegerInBounds(len(upperCaseRune))
	str += string(upperCaseRune[uci])
	for i := 0; i < 5; i++ {
		lci := randomIntegerInBounds(len(lowerCaseRune))
		str += string(lowerCaseRune[lci])
	}
	for i := 0; i < 2; i++ {
		sci := randomIntegerInBounds(len(symbolRune))
		str += string(symbolRune[sci])
	}
	for i := 0; i < 4; i++ {
		nci := randomIntegerInBounds(len(numberRune))
		str += string(numberRune[nci])
	}
	return str
}

// IsEmail lets us know whether the supplied string is an email field that is
// alphanumeric, i.e. contains uppercase letters, lower case letters, numbers
// and underscores
func IsEmail(email string) bool {
	re := regexp.MustCompile(`^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$`)
	return re.Match([]byte(email))
}

// IsSimpleString lets us know whether the supplied string is a simple string
// that is alphanumeric, i.e. contains uppercase letters, lower case letters,
// numbers and underscores
func IsSimpleString(str string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9_]*$")
	return re.Match([]byte(str))
}
