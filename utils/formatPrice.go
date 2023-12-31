package utils

import (
	"strings"
)

// AddCommasToPrice 함수는 가격에 쉼표를 추가합니다.
func AddCommasToPrice(price string) string {
	parts := strings.Split(price, ".")
	intPart := parts[0]
	var formattedPrice string

	for i, c := range reverse(intPart) {
		if i > 0 && i%3 == 0 {
			formattedPrice = "," + formattedPrice
		}
		formattedPrice = string(c) + formattedPrice
	}

	if len(parts) > 1 {
		formattedPrice += "." + parts[1]
	}

	return formattedPrice
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
