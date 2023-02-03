package iteration

// const repeatCount = 5

// Repeat returns character repeatCount times
func Repeat(char string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += char
	}

	return repeated
}
