package private

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	endPoint = "/info/account"
)

func AccountEndPoint() string {
	return endPoint
}

func GetAccountEncode(orderCurrency, quoteCurrency string) string {
	// URL 값을 인코딩합니다.
	values := url.Values{}
	values.Set("endpoint", endPoint)
	values.Set("order_currency", orderCurrency)
	values.Set("payment_currency", quoteCurrency)
	encodedURL := values.Encode()
	return encodedURL
}

func Account(apiKey, apiSecret, privateBaseURL, baseCurrency, quoteCurrency string) {
	accountUrl := fmt.Sprintf("%s%s", privateBaseURL, endPoint)

	// API Nonce 생성
	apiNonce := time.Now().UnixNano() / int64(time.Millisecond)
	strApiNonce := strconv.FormatInt(apiNonce, 10)

	// accountParamEncode
	accountParamEncode := GetAccountEncode(baseCurrency, quoteCurrency)

	// API Sign 생성
	signature, err := GenerateAPISign(endPoint, accountParamEncode, apiSecret, strApiNonce)
	if err != nil {
		fmt.Println("빗썸전용 API-Sign 생성 오류:", err)
		return
	}

	fmt.Println("빗썸전용 API-Sign:", signature)

	payload := strings.NewReader("order_currency=" + baseCurrency + "&payment_currency=" + quoteCurrency)

	req, _ := http.NewRequest("POST", accountUrl, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Api-Sign", signature)
	req.Header.Add("Api-Nonce", strApiNonce)
	req.Header.Add("Api-Key", apiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}
