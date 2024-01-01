package public

import (
	"encoding/json"
	"fmt"
	"github.com/cokeys90/auto-bot-bithumb/bithumb"
	"io/ioutil"
	"net/http"
)

type TickerData struct {
	OpeningPrice     string `json:"opening_price"`
	ClosingPrice     string `json:"closing_price"`
	MinPrice         string `json:"min_price"`
	MaxPrice         string `json:"max_price"`
	UnitsTraded      string `json:"units_traded"`
	AccTradeValue    string `json:"acc_trade_value"`
	PrevClosingPrice string `json:"prev_closing_price"`
	UnitsTraded24H   string `json:"units_traded_24H"`
	AccTradeValue24H string `json:"acc_trade_value_24H"`
	Fluctate24H      string `json:"fluctate_24H"`
	FluctateRate24H  string `json:"fluctate_rate_24H"`
	Date             string `json:"date"`
}

type TickerResponse struct {
	Status string     `json:"status"`
	Data   TickerData `json:"data"`
}

func Ticker(reqData bithumb.ReqData) (price string) {
	// ticker 조회
	url := fmt.Sprintf("%s/public/ticker/%s", reqData.BaseUrl, reqData.TradingPair)
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

	// JSON 응답 데이터를 구조체로 디코딩
	var tickerResponse TickerResponse
	err = json.Unmarshal(body, &tickerResponse)
	if err != nil {
		fmt.Println("JSON 디코딩 오류:", err)
		return
	}

	//// 가격에 쉼표 추가
	//closingPriceWithCommas := utils.AddCommasToPrice(tickerResponse.Data.ClosingPrice)
	//
	//// 디코딩한 데이터와 현재 시간 출력
	//fmt.Printf("현재 가격: %s\n", closingPriceWithCommas)

	return tickerResponse.Data.ClosingPrice
}
