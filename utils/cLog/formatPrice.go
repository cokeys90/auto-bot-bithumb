package cLog

import (
	"github.com/shopspring/decimal"
	"math/rand"
	"strconv"
	"strings"
	"time"
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

// getPriceUnit 함수는 주어진 가격 문자열에 대한 호가 단위를 반환합니다.
// 가격 문자열을 실수로 파싱하지 못하면 에러를 반환합니다.
func PriceUnit(priceStr string) (float64, error) {
	// 문자열을 실수로 변환합니다.
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, err
	}

	// 변환된 가격을 사용하여 호가 단위를 결정합니다.
	switch {
	case price < 1.0: // 1원 미만
		return 0.0001, nil
	case price < 10.0: // 1원 이상 10원 미만
		return 0.001, nil
	case price < 100.0: // 10원 이상 100원 미만
		return 0.01, nil
	case price < 1000.0: // 100원 이상 1,000원 미만
		return 1, nil
	case price < 5000.0: // 1,000원 이상 5,000원 미만
		return 5, nil
	case price < 10000.0: // 5,000원 이상 10,000원 미만
		return 5, nil
	case price < 50000.0: // 10,000원 이상 50,000원 미만
		return 10, nil
	case price < 100000.0: // 50,000원 이상 100,000원 미만
		return 50, nil
	case price < 500000.0: // 100,000원 이상 500,000원 미만
		return 100, nil
	case price < 1000000.0: // 500,000원 이상 1,000,000원 미만
		return 500, nil
	default: // 1,000,000원 이상
		return 1000, nil
	}
}

func DecimalRnd(min, max float64) decimal.Decimal {
	rand.Seed(time.Now().UnixNano()) // 랜덤 시드 초기화

	minValue := decimal.NewFromFloat(min) // 최소 값
	maxValue := decimal.NewFromFloat(max) // 최대 값

	// 최소 값과 최대 값 사이의 랜덤 Decimal 생성
	randomDecimal := minValue.Add(decimal.NewFromFloat(rand.Float64()).Mul(maxValue.Sub(minValue))).Round(4)

	return randomDecimal
}
func IntRnd(min, max int) int {
	rand.Seed(time.Now().UnixNano()) // 랜덤 시드 초기화
	return rand.Intn(max-min+1) + min
}
