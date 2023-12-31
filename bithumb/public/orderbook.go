package public

import (
	"encoding/json"
	"fmt"
	"github.com/cokeys90/auto-bot-bithumb/utils"
	"io/ioutil"
	"net/http"
)

type OrderBookData struct {
	Timestamp       string      `json:"timestamp"`
	PaymentCurrency string      `json:"payment_currency"`
	OrderCurrency   string      `json:"order_currency"`
	Bids            []OrderItem `json:"bids"`
	Asks            []OrderItem `json:"asks"`
}

type OrderItem struct {
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

type OrderBookResponse struct {
	Status string        `json:"status"`
	Data   OrderBookData `json:"data"`
}

func OrderBook(publicBaseURL string, currency string, orderBookCnt int) {
	url := fmt.Sprintf("%s/orderbook/%s?count=%d", publicBaseURL, currency, orderBookCnt)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("API 요청 중 오류 발생:", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("응답 데이터 읽기 오류:", err)
		return
	}

	var orderBookResponse OrderBookResponse
	err = json.Unmarshal(body, &orderBookResponse)
	if err != nil {
		fmt.Println("JSON 디코딩 오류:", err)
		return
	}

	var orderBookData = orderBookResponse.Data

	for i, ask := range orderBookData.Asks {
		askPrice := utils.AddCommasToPrice(ask.Price)
		askQty := utils.AddCommasToPrice(ask.Quantity)
		fmt.Printf("매도 #%d - 가격: %s, 수량: %s\n", i+1, askPrice, askQty)
	}

	for i, bid := range orderBookData.Bids {
		// 가격에 쉼표 추가
		bidPrice := utils.AddCommasToPrice(bid.Price)
		bidQty := utils.AddCommasToPrice(bid.Quantity)
		fmt.Printf("매수 #%d - 가격: %s, 수량: %s\n", i+1, bidPrice, bidQty)
	}
}
