package main

func main() {
}

func sameEnd(str string) string {
	length := len([]rune(str))
	half := length / 2

	left := ""
	right := ""

	size := 0

	for i := 0; i < half; i++ {
		left += string(str[i])
		right = string(str[length-1-i]) + right
		if left == right {
			size = len([]rune(left))
		}
	}
	result := str[0:size]

	return result
}