package utils

import (
	"fmt"
	"math/big"
	"testing"
)

func TestCalculatePrices(t *testing.T) {
	// 정밀도 설정된 big.Float 생성
	min, _ := new(big.Float).SetString("93.1")
	max, _ := new(big.Float).SetString("95.1")
	increment, _ := new(big.Float).SetString("0.05")

	values := CalculatePrices(min, max, increment)

	// 값 출력 시 소수점 이하 정밀도 지정
	for _, value := range values {
		fmt.Println(value) // 소수점 이하 2자리까지 출력
	}
}
