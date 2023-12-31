package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

type Config struct {
	JdbcUrl  string
	Port     int
	Scheme   string
	Username string
	Password string
}

var database *sql.DB

func InitDatabase(_config Config) {
	//loc, err := time.LoadLocation("Asia/Seoul")
	//if err != nil {
	//	panic(err)
	//}

	cfg := mysql.NewConfig()
	cfg.User = _config.Username
	cfg.Passwd = _config.Password
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%d", _config.JdbcUrl, _config.Port)
	cfg.Collation = "utf8mb4_general_ci"
	//cfg.Loc = loc
	cfg.DBName = _config.Scheme
	cfg.ParseTime = true

	connector, err := mysql.NewConnector(cfg)
	if err != nil {
		panic(err)
	}

	rawDb := sql.OpenDB(connector)

	database = rawDb
}

// GetDatabase DB 커넥션 반환
func GetDatabase() *sql.DB {
	return database
}

// IDatabaseCaller Database 호출자 인터페이스 (쿼리 호출 또는 Tx 호출에 사용)
type IDatabaseCaller interface {
	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row

	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}
