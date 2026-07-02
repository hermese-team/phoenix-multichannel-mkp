package mysql

type Config struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
}
