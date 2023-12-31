package private

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

func GenerateAPISign(endPoint, requestParams, apiSecret, apiNonce string) (string, error) {
	// 서명 문자열을 생성합니다.
	signature := fmt.Sprintf("%s;%s;%s", endPoint, requestParams, apiNonce)

	// HMAC-SHA512 해시를 계산합니다.
	key := []byte(apiSecret)
	h := hmac.New(sha512.New, key)
	h.Write([]byte(signature))
	signatureBytes := h.Sum(nil)

	// Base64로 변환하여 반환합니다.
	signatureBase64 := base64.StdEncoding.EncodeToString(signatureBytes)
	return signatureBase64, nil
}
