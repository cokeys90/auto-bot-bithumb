package private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cokeys90/auto-bot-bithumb/bithumb"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	tradeCancelEndPoint = "/trade/cancel"
)

type TradeCancelResponse struct {
	Status string `json:"status"`
}

func TradeCancel(reqData bithumb.ReqData, tradeType, orderId string) (order_id string) {
	accountUrl := fmt.Sprintf("%s%s", reqData.BaseUrl, tradeCancelEndPoint)

	// 1.API Nonce 생성
	nonce := fmt.Sprint(time.Now().UnixNano() / int64(time.Millisecond))

	// 2.RequestParam, Sign 할때랑 Body 담을때 같이 사용
	values := url.Values{}
	values.Set("endpoint", tradeCancelEndPoint)
	values.Set("order_currency", reqData.BaseCurrency)
	values.Set("payment_currency", reqData.QuoteCurrency)
	values.Set("order_id", orderId)
	values.Set("type", tradeType)
	requestParams := values.Encode()

	// 3.API Sign 생성
	signature, err := GenerateAPISign(tradeCancelEndPoint, requestParams, reqData.ApiSecret, nonce)
	if err != nil {
		fmt.Println("빗썸전용 API-Sign 생성 오류:", err)
		return ""
	}

	req, _ := http.NewRequest("POST", accountUrl, bytes.NewBufferString(requestParams))

	req.Header.Add("Content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(requestParams)))
	req.Header.Add("Api-Sign", signature)
	req.Header.Add("Api-Nonce", nonce)
	req.Header.Add("Api-Key", reqData.ApiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	//fmt.Println(string(body))

	var tradeResponse TradeCancelResponse
	err = json.Unmarshal(body, &tradeResponse)
	if err != nil {
		fmt.Println("JSON 디코딩 오류:", err)
		return ""
	}

	return
}
