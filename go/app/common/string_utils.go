package common

const defaultSuffix = "..."

func RoundStrings(strings []string, maxLength []int) []string {
	for i, s := range strings {

		utfS := []rune(s)

		if len(utfS) > maxLength[i] {
			suffixToAdd := defaultSuffix

			if maxLength[i] < len(defaultSuffix) {
				suffixToAdd = defaultSuffix[:maxLength[i]]
			}
			suffixLen := len(suffixToAdd)

			strings[i] = string(utfS[:maxLength[i]-suffixLen]) + suffixToAdd
		}
	}

	return strings
}
