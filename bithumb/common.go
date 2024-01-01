package bithumb

type Tear float64

const (
	WHITE  Tear = 0.00003
	BLUE   Tear = 0.00005
	GREEN  Tear = 0.00008
	PURPLE Tear = 0.0001
	ORANGE Tear = 0.0001
	BLACK  Tear = 0.0001

	BUY  string = "bid"
	SELL string = "ask"
)

type ReqData struct {
	BaseUrl   string
	ApiKey    string
	ApiSecret string

	OrderBookCnt int

	BaseCurrency  string
	QuoteCurrency string
	TradingPair   string
}
