package randomutilsgo

import (
	"math"
	"math/rand"
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
