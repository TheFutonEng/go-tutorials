package strings

import (
	"fmt"
	"strings"
)

func Stringidx(s, substr string) int {
	idx := strings.Index(s, substr)
	return idx
}

func Bookend(s string) string {
	leftBook := s[0:]
	rightBook := s[len(s)-1:]

	fmt.Println(leftBook)

	return string(leftBook + rightBook)
}
