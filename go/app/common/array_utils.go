package common

func FillArrayWithValue(array []int, value int) {
	for i := range array {
		array[i] = value
	}
}
