package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

const configPrefix = "GO_APP"

var (
	// 환경변수 바인딩 정보 배열
	bindInfoArr = []KeyBindInfo{
		{
			BindKey:    "SERVER_PORT",
			JsonKey:    "server.port",
			Required:   true,
			RequireMsg: "서버 포트 필수",
		}, // 서버 포트
		{
			BindKey:    "SERVER_JWT_SECRET",
			JsonKey:    "server.jwtSecret",
			Required:   false,
			RequireMsg: "서버 JWT 비밀키 필수",
		}, // JWT 비밀키
		{
			BindKey:    "DB_JDBC_URL",
			JsonKey:    "database.jdbcUrl",
			Required:   false,
			RequireMsg: "DB JDBC URL 필수",
		}, // JDBC URL
		{
			BindKey:    "DB_PORT",
			JsonKey:    "database.port",
			Required:   false,
			RequireMsg: "DB 포트 필수",
		}, // DB 포트
		{
			BindKey:    "DB_SCHEME",
			JsonKey:    "database.scheme",
			Required:   false,
			RequireMsg: "DB 스키마 필수",
		}, // DB 명 (스키마 명)
		{
			BindKey:    "DB_USERNAME",
			JsonKey:    "database.username",
			Required:   false,
			RequireMsg: "DB 접속 계정 필수",
		}, // DB 접속 계정
		{
			BindKey:    "DB_PASSWORD",
			JsonKey:    "database.password",
			Required:   false,
			RequireMsg: "DB 접속 비밀번호 필수",
		}, // DB 접속 비밀번호
	}
)

type KeyBindInfo struct {
	BindKey    string // 환경변수 키
	JsonKey    string // 바인딩 키
	Required   bool   // 필수 여부
	RequireMsg string // 필수 여부 메시지
}

var config ApplicationInfo

type ApplicationInfo struct {
	isLoaded bool         // 설정이 로드되었는지 여부
	Server   ServerInfo   `json:"server"`
	Database DatabaseInfo `json:"database"`
}

type ServerInfo struct {
	Port      int    `json:"port"`
	JwtSecret string `json:"jwtSecret"`
}

type DatabaseInfo struct {
	JdbcUrl  string `json:"jdbcUrl"`
	Port     int    `json:"port"`
	Scheme   string `json:"scheme"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func loadConfig() ApplicationInfo {
	// 설정 초기화
	_initConfig()
	// 설정 읽기
	config = ApplicationInfo{}
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("fatal error config (to struct failed): %w", err))
	}
	if config.Server.Port != 0 {
		config.isLoaded = true
		viper.Reset()
	}
	return config
}

// _initConfig 설정 초기화
func _initConfig() {
	// 환경변수 설정
	viper.SetEnvPrefix(configPrefix)
	for _, bindInfo := range bindInfoArr {
		_ = viper.BindEnv(bindInfo.BindKey) // 환경변수 바인딩
		if bindInfo.Required {              // 필수 여부 체크
			if !viper.IsSet(bindInfo.BindKey) {
				log.Panic(fmt.Errorf("fatal error config (required): %s", bindInfo.RequireMsg))
			}
		}
		viper.Set(bindInfo.JsonKey, viper.Get(bindInfo.BindKey)) // 바인딩된 환경변수를 JSON Key에 바인딩
	}

	for _, key := range viper.AllKeys() {
		fmt.Printf("%s: %v\n", key, viper.Get(key))
	}

}

// GetConfig 어플리케이션 설정 반환 (없으면 설정 로드 후 반환)
func GetConfig() ApplicationInfo {
	if !config.isLoaded {
		loadConfig()
	}
	return config
}
