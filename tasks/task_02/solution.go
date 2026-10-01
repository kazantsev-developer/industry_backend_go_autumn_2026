package main

func rotateRunes(s string, shift int) string {
	runes := []rune(s)
	n := len(runes)
	if n == 0 {
		return ""
	}

	shift %= n
	if shift < 0 {
		shift += n
	}

	result := make([]rune, n)
	copy(result, runes[shift:])
	copy(result[n-shift:], runes[:shift])
	return string(result)
}
