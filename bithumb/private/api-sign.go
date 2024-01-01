package private

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

func GenerateAPISign(endPoint, requestParams, apiSecret, apiNonce string) (string, error) {
	// 서명 문자열을 생성합니다.
	//signature := fmt.Sprintf("%s0%s0%s", tradeCancelEndPoint, requestParams, apiNonce)
	signature := endPoint + string(0) + requestParams + string(0) + apiNonce

	hmacParsed := hmac.New(sha512.New, []byte(apiSecret))
	hmacParsed.Write([]byte(signature))

	hexData := hex.EncodeToString(hmacParsed.Sum(nil))
	byteHexData := []byte(hexData)
	hmacParsed.Reset()

	result := base64.StdEncoding.EncodeToString(byteHexData)
	return result, nil
}
