package mysql

import "context"

type Test02DbManager interface {
	Close() error
	Setxxxx(ctx context.Context)
	Getxxxx(ctx context.Context)
}
