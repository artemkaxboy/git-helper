package common

func RoundStrings(strings []string, maxLength []int) []string {
	for i, s := range strings {

		utfS := []rune(s)

		if len(utfS) > maxLength[i] {
			strings[i] = string(utfS[:maxLength[i]-3]) + "..."
		}
	}

	return strings
}
