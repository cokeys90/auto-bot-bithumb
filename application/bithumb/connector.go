package bithumb

// 외부 거래소 인터페이스
type RestConnector interface {
	Depth()      // 호가 목록 조회
	Ticker()     // 헌재가 조회
	Currencies() // 화폐 목록
}

func InitConnector() RestConnector {
	var conn RestConnector
	conn = NewBithumbRestConnector()
	return conn
}
