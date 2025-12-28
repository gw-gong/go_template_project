package mysql

import "context"

type Test01DbManager interface {
	Close() error
	Setxxxx(ctx context.Context)
	Getxxxx(ctx context.Context)
}
