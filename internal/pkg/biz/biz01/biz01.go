package biz01

import "context"

type Biz01 interface {
	Function01(ctx context.Context)
	Function02(ctx context.Context)
}
