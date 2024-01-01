package utils

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

// RandomPrice 함수는 CalculatePrices의 결과 배열에서 랜덤한 하나의 값을 반환합니다.
func RandomPrice(minPrice, maxPrice, priceIncrement *big.Float) *big.Float {
	prices := CalculatePrices(minPrice, maxPrice, priceIncrement)

	if len(prices) == 0 {
		return nil // 가격 배열이 비어있을 경우 nil 반환
	}

	// 랜덤 시드 설정
	rand.Seed(time.Now().UnixNano())

	// 랜덤 인덱스 선택
	randomIndex := rand.Intn(len(prices))

	// 랜덤 가격 반환
	return prices[randomIndex]
}

// calculatePrices 함수는 주어진 최소가격, 최대가격, 가격 증가치에 따라 가격 리스트를 생성합니다.
func CalculatePrices(minPrice, maxPrice, priceIncrement *big.Float) []*big.Float {
	var prices []*big.Float

	// 최소값이 최대값보다 크면 빈 리스트 반환
	if minPrice.Cmp(maxPrice) > 0 {
		fmt.Println("calculatePrices()", "minPrice is bigger than maxPrice => 최소값이 최대값보다 크다.")
		return prices
	}

	// 생성할 가격의 범위가 priceIncrement 보다 작으면 빈 리스트 반환
	if new(big.Float).Sub(maxPrice, minPrice).Cmp(priceIncrement) <= 0 {
		fmt.Println("calculatePrices()", "maxPrice - minPrice is smaller than priceIncrement => 생성할 가격의 범위가 priceIncrement 보다 작다.")
		return prices
	}

	// 초기 currentPrice 설정
	currentPrice := new(big.Float).Mul(new(big.Float).Quo(minPrice, priceIncrement).SetMode(big.ToZero), priceIncrement)
	currentPrice = new(big.Float).Add(currentPrice, priceIncrement)

	// increment의 소수점 이하 자리수 계산
	incrementString := priceIncrement.Text('f', -1)
	dotIndex := strings.Index(incrementString, ".")
	decimalPlaces := 0
	if dotIndex != -1 {
		decimalPlaces = len(incrementString) - dotIndex - 1
	}

	// minPrice < currentPrice < maxPrice 범위 내에서 가격 리스트 생성
	for currentPrice.Cmp(maxPrice) < 0 {
		savePrice := truncateBigFloat(currentPrice, decimalPlaces)
		prices = append(prices, new(big.Float).Set(savePrice))
		currentPrice = new(big.Float).Add(currentPrice, priceIncrement)
	}

	return prices
}

// truncateBigFloat 함수는 big.Float 값을 주어진 소수점 자리수로 자릅니다.
func truncateBigFloat(f *big.Float, precision int) *big.Float {
	multiplier := new(big.Float).SetFloat64(math.Pow(10, float64(precision)))
	truncated := new(big.Float).Mul(f, multiplier)
	truncatedInt, _ := truncated.Int(nil) // 정수로 변환 (내림)
	return new(big.Float).Quo(new(big.Float).SetInt(truncatedInt), multiplier)
}
