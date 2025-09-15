package mysql

type XxxDbManager interface {
	Close() error
	Setxxxx()
	Getxxxx()
}

func NewXxxDbManager() XxxDbManager {
	return newMysqlClient()
}