package main

import (
	"fmt"
	"github.com/cokeys90/auto-bot-bithumb/bithumb/private"
	"time"
)

const (
	publicBaseURL  = "https://api.bithumb.com/public"
	privateBaseURL = "https://api.bithumb.com"
	baseCurrency   = "BTC" // 종목심볼. ex)XRP
	quoteCurrency  = "KRW" // 결제통화 심볼. ex)KRW
	orderBookCnt   = 1     // 조회 개수

	currency = baseCurrency + "_" + quoteCurrency // bithumb은 {baseCurrency}_{quoteCurrency}

	apiKey    = "1"
	apiSecret = "1"
)

func PrintCurrentTimeAndCurrency() {
	// 현재 시간을 원하는 형식으로 포맷팅
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// 주문 통화와 결제 통화를 출력
	fmt.Printf("%s - 주문 통화: %s, 결제 통화: %s\n", currentTime, baseCurrency, quoteCurrency)
}

func main() {
	for {
		fmt.Println("************")
		//PrintCurrentTimeAndCurrency()
		//public.Ticker(publicBaseURL, currency)
		//public.OrderBook(publicBaseURL, currency, orderBookCnt)
		//fmt.Println("")

		private.Account(apiKey, apiSecret, privateBaseURL, baseCurrency, quoteCurrency)

		// 1초 대기
		time.Sleep(1 * time.Second)
	}
}
