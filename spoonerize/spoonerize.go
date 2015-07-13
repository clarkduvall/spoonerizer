package spoonerize

import (
	"unicode"
	"unicode/utf8"
)

func isSpace(c byte) bool {
	return unicode.IsSpace(rune(c))
}

func isLetter(c byte) bool {
	return unicode.IsLetter(rune(c))
}

func isNotLetter(c byte) bool {
	return !isLetter(c)
}

func isUpper(c byte) bool {
	return unicode.IsUpper(rune(c))
}

func toUpper(c byte) byte {
	return byte(unicode.ToUpper(rune(c)))
}

func toLower(c byte) byte {
	return byte(unicode.ToLower(rune(c)))
}

func swapLetters(bytes []byte, i1 int, s1 int, i2 int, s2 int) {
	// Prevent allocations by using an array.
	var prefix1, prefix2 [3]byte
	copy(prefix1[:], bytes[i1:i1+s1])
	copy(prefix2[:], bytes[i2:i2+s2])

	diff := s2 - s1
	copy(bytes[i1+s1+diff:i2+diff], bytes[i1+s1:i2])

	upper1, upper2 := isUpper(prefix1[0]), isUpper(prefix2[0])
	copy(bytes[i2+diff:i2+diff+s1], prefix1[:])
	copy(bytes[i1:i1+s2], prefix2[:])
	if upper1 != upper2 {
		if upper1 {
			bytes[i1] = toUpper(bytes[i1])
			bytes[i2+diff] = toLower(bytes[i2+diff])
		} else {
			bytes[i1] = toLower(bytes[i1])
			bytes[i2+diff] = toUpper(bytes[i2+diff])
		}
	}
}

type pred func(byte) bool

func advanceUntil(bytes []byte, i int, fn pred) int {
	for l := len(bytes); i < l && !fn(bytes[i]); i++ {
	}
	return i
}

func getNextPrefix(bytes []byte, i int) (int, int, int) {
	start := advanceUntil(bytes, i, isLetter)
	end := advanceUntil(bytes, start, isNotLetter)
	size := end - start
	i = advanceUntil(bytes, end, isSpace)

	if size == 0 {
		return -1, -1, i
	}

	// Ignore non ASCII.
	if !utf8.FullRune(bytes[start : start+1]) {
		return -1, -1, i
	}

	if size > maxStopwordLength {
		size = maxStopwordLength
	}

	var lower [maxStopwordLength]byte
	for i := 0; i < size; i++ {
		lower[i] = byte(unicode.ToLower(rune(bytes[start+i])))
	}

	if vowels[lower[0]] {
		return -1, -1, i
	}

	if stopwords[string(lower[:size])] {
		return -1, -1, i
	}

	if size > 2 && trigraphs[string(lower[:3])] {
		return start, 3, i
	}

	if size > 1 && digraphs[string(lower[:2])] {
		return start, 2, i
	}

	return start, 1, i
}

func Spoonerize(textBytes []byte) []byte {
	start, size, tempStart, tempSize := -1, 0, -1, 0
	for l, i := len(textBytes), 0; i < l; i++ {
		tempStart, tempSize, i = getNextPrefix(textBytes, i)
		if tempStart < 0 || tempSize <= 0 {
			continue
		}

		if start < 0 || size <= 0 {
			start, size = tempStart, tempSize
		} else {
			swapLetters(textBytes, start, size, tempStart, tempSize)
			start, size = -1, -1
		}
	}
	return textBytes
}
