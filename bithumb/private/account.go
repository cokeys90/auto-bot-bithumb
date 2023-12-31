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
	accountEndPoint = "/info/account"
)

type AccountResponse struct {
	Status string      `json:"status"`
	Data   AccountData `json:"data"`
}

type AccountData struct {
	Created         string `json:"created"`
	AccountID       string `json:"account_id"`
	OrderCurrency   string `json:"order_currency"`
	PaymentCurrency string `json:"payment_currency"`
	TradeFee        string `json:"trade_fee"`
	Balance         string `json:"balance"`
}

func Account(reqData bithumb.ReqData) (tradeFee, userBalance string) {
	accountUrl := fmt.Sprintf("%s%s", reqData.BaseUrl, accountEndPoint)

	// 1.API Nonce 생성
	nonce := fmt.Sprint(time.Now().UnixNano() / int64(time.Millisecond))

	// 2.RequestParam, Sign 할때랑 Body 담을때 같이 사용
	values := url.Values{}
	values.Set("endpoint", accountEndPoint)
	values.Set("order_currency", reqData.BaseCurrency)
	values.Set("payment_currency", reqData.QuoteCurrency)
	requestParams := values.Encode()

	// 3.API Sign 생성
	signature, err := GenerateAPISign(accountEndPoint, requestParams, reqData.ApiSecret, nonce)
	if err != nil {
		fmt.Println("빗썸전용 API-Sign 생성 오류:", err)
		return "-1", ""
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

	var accountResponse AccountResponse
	err = json.Unmarshal(body, &accountResponse)
	if err != nil {
		fmt.Println("JSON 디코딩 오류:", err)
		return "-1", ""
	}

	return accountResponse.Data.TradeFee, accountResponse.Data.Balance
}
