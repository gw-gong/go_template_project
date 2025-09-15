package mysql

type mysqlClient struct {}

func newMysqlClient() *mysqlClient {
	return &mysqlClient{}
}

func (m *mysqlClient) Close() error {
	return nil
}

func (m *mysqlClient) Setxxxx() {}

func (m *mysqlClient) Getxxxx() {}
