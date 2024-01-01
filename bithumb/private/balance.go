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
	"strings"
	"time"
)

const (
	balanceEndPoint = "/info/balance"
)

type BalanceResponse struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}

func Balance(reqData bithumb.ReqData) (bp, base, quote string) {
	accountUrl := fmt.Sprintf("%s%s", reqData.BaseUrl, balanceEndPoint)

	// 1.API Nonce 생성
	nonce := fmt.Sprint(time.Now().UnixNano() / int64(time.Millisecond))

	// 2.RequestParam, Sign 할때랑 Body 담을때 같이 사용
	values := url.Values{}
	values.Set("endpoint", balanceEndPoint)
	values.Set("currency", "ALL")
	requestParams := values.Encode()

	// 3.API Sign 생성
	signature, err := GenerateAPISign(balanceEndPoint, requestParams, reqData.ApiSecret, nonce)
	if err != nil {
		fmt.Println("빗썸전용 API-Sign 생성 오류:", err)
		return "", "", ""
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

	var balanceResponse BalanceResponse
	err = json.Unmarshal(body, &balanceResponse)
	if err != nil {
		fmt.Println("JSON 디코딩 오류:", err)
		return "", "", ""
	}

	//  Bithumb BP 는 total_p
	totalP, ok := balanceResponse.Data["total_p"].(string)
	if ok {
		//fmt.Println("Bithumb BP:", totalP)
	}

	_baseCurrency := strings.ToLower(reqData.BaseCurrency)
	availableBaseCurrency, ok := balanceResponse.Data["available_"+_baseCurrency].(string)
	if ok {
		//fmt.Println("사용가능 코인 " + baseCurrency + ":" + availableBaseCurrency)
	}

	_quoteCurrency := strings.ToLower(reqData.QuoteCurrency)
	availableQuoteCurrency, ok := balanceResponse.Data["available_"+_quoteCurrency].(string)
	if ok {
		//fmt.Println("사용가능 코인 " + quoteCurrency + ":" + availableQuoteCurrency)
	}

	return totalP, availableBaseCurrency, availableQuoteCurrency
}
