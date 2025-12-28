package biz02

import "context"

type Biz02 interface {
	Function01(ctx context.Context)
	Function02(ctx context.Context)
}
