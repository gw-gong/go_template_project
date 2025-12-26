package component02

import "context"

type Component02 interface {
	Function01(ctx context.Context)
	Function02(ctx context.Context)
}
