package main

import (
	"fmt"
	"github.com/cokeys90/auto-bot-bithumb/bithumb"
	"github.com/cokeys90/auto-bot-bithumb/bithumb/private"
	"github.com/cokeys90/auto-bot-bithumb/bithumb/public"
	"github.com/cokeys90/auto-bot-bithumb/utils"
	"github.com/cokeys90/auto-bot-bithumb/utils/cLog"
	"github.com/shopspring/decimal"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

const (
	baseUrl       = "https://api.bithumb.com"
	baseCurrency  = "META" // 종목심볼. ex)XRP
	quoteCurrency = "KRW"  // 결제통화 심볼. ex)KRW
	orderBookCnt  = 1      // 조회 개수

	tradingPair = baseCurrency + "_" + quoteCurrency // bithumb은 {baseCurrency}_{quoteCurrency}

	apiKey    = "1"
	apiSecret = "1"

	userTear = bithumb.BLUE
)

func PrintCurrentTimeAndCurrency() {
	// 현재 시간을 원하는 형식으로 포맷팅
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// 주문 통화와 결제 통화를 출력
	fmt.Printf("%s_(%s)_", currentTime, tradingPair)
}

func main() {
	totalRewardBP := decimal.NewFromFloat(0)

	for {
		sleepTime := cLog.IntRnd(500, 9000)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)

		reqData := bithumb.ReqData{
			BaseUrl:       baseUrl,
			ApiKey:        apiKey,
			ApiSecret:     apiSecret,
			OrderBookCnt:  orderBookCnt,
			BaseCurrency:  baseCurrency,
			QuoteCurrency: quoteCurrency,
			TradingPair:   tradingPair,
		}

		//fmt.Println("\n************")
		PrintCurrentTimeAndCurrency()

		tradeFee, _ := private.Account(reqData)
		if tradeFee != "0" {
			fmt.Println("(0원 일때만 실행) 수수료 : " + tradeFee)
			continue // 수수료가 0이 아니면 자산에 로스가 발생할 수 있다.
		}

		BP, baseAmount, quoteAmount := private.Balance(reqData)

		// BP 하루에 10만개 todo: BP 최대치 도달 관리는 나중에
		// tip:내가 사둔 COIN 개수 만큼 거래함, 5:5비율 대충 원화가 많게 맞춰두면됨

		fBP, err := strconv.ParseFloat(BP, 64)
		if err != nil {
			fmt.Println("BP 문자열을 부동소수점으로 변환하는 중 에러 발생:", err)
			return
		}

		fBaseAmount, err := strconv.ParseFloat(baseAmount, 64)
		if err != nil {
			fmt.Println("baseAmount 문자열을 부동소수점으로 변환하는 중 에러 발생:", err)
			return
		}

		fQuoteAmount, err := strconv.ParseFloat(quoteAmount, 64)
		if err != nil {
			fmt.Println("quoteAmount 문자열을 부동소수점으로 변환하는 중 에러 발생:", err)
			return
		}

		price := public.Ticker(reqData)
		sAskPrice, sBidPrice := public.OrderBook(reqData)

		// 호가 단위 계산
		unit, err := cLog.PriceUnit(price)
		if err != nil {
			continue // 계산실패? ㅠㅠ
		}

		askPrice, err := strconv.ParseFloat(sAskPrice, 64)
		if err != nil {
			continue
		}
		bidPrice, err := strconv.ParseFloat(sBidPrice, 64)
		if err != nil {
			continue
		}

		// 매도 - 매수 > 호가 단위
		if askPrice-bidPrice <= unit {
			fmt.Println("빈틱 없음...ㅠㅠ")
			continue // 빈틱 없음 ... ㅠ
		}

		// 매도: 보유한 코인 있는만큼 빈틱가격에 전량 매도

		dAskPrice := decimal.NewFromFloat(askPrice)
		dUnit := decimal.NewFromFloat(unit)

		// 빈틱 초기화(기본값)
		autoPrice := dAskPrice.Sub(dUnit)
		sAutoPrice := autoPrice.String()

		// 빈틱에서 가격 정하기
		bUnit := new(big.Float).SetFloat64(unit)
		bBidPrice := new(big.Float).SetFloat64(bidPrice)
		bAskPrice := new(big.Float).SetFloat64(askPrice)
		prices := utils.CalculatePrices(bBidPrice, bAskPrice, bUnit)

		if len(prices) != 0 { // 없으면 그냥 매도에서 한틱 아래로 거래, 빈틱 이 있는데 없다고 나올때?
			// 랜덤 시드 설정
			rand.Seed(time.Now().UnixNano())

			// 랜덤 인덱스 선택
			randomIndex := rand.Intn(len(prices))

			// 랜덤 가격 반환
			sAutoPrice = prices[randomIndex].String()
		}

		// 총 가격 (가격 * 수량)
		// 에서 최소 거래 금액에 맞는 수량 구하기
		// 5000 / 가격 = 최소 수량
		d5000won := decimal.NewFromFloat(5000)
		minAmount := d5000won.Div(dAskPrice) // 최소수량
		fMinAmount, _ := minAmount.Float64()
		rndAmount := cLog.DecimalRnd(fMinAmount, fBaseAmount)

		sellOrderId := private.Trade(reqData, bithumb.SELL, sAutoPrice, rndAmount.String())
		time.Sleep(30 * time.Millisecond)
		// 매수: 매도 한 물량 만큼 구매
		buyOrderId := private.Trade(reqData, bithumb.BUY, sAutoPrice, rndAmount.String())
		time.Sleep(30 * time.Millisecond)

		// 주문 취소
		private.TradeCancel(reqData, bithumb.SELL, sellOrderId)
		private.TradeCancel(reqData, bithumb.BUY, buyOrderId)

		// 기본정보 출력
		printMsg := fmt.Sprintf("BP:%d/%s:%d/%s:%d", int(fBP), baseCurrency, int(fBaseAmount), quoteCurrency, int(fQuoteAmount))

		// 매도/매수 주문 정보

		printMsg += fmt.Sprintf("/(주문)가격/수량:%s/%s", sAutoPrice, rndAmount.String())

		// 예상 누적 BP
		_userTear := decimal.NewFromFloat(float64(userTear))
		reward := rndAmount.Mul(_userTear).Mul(autoPrice)
		totalRewardBP = totalRewardBP.Add(reward)
		printMsg += fmt.Sprintf(" =>(예상)누적BP:%s", totalRewardBP.String())

		fmt.Println(printMsg)
	}
}
