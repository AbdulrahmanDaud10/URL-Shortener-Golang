package api

import (
	"math"
	"strings"
)

// characters used for Base62 encoding
const (
	Base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func Base62Encode(num int) string {
	encoded := ""
	for num > 0 {
		remainder := num % 62
		num /= 62
		encoded = string(Base62Chars[remainder]) + encoded
	}
	return encoded
}
func Base62Decode(str string) int {
	decoded := 0
	for i := 0; i < len(str); i++ {
		pos := strings.Index(Base62Chars, string(str[i]))
		decoded += pos * int(math.Pow(62, float64(i)))
	}
	return decoded
}
