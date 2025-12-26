package component01

import "context"

type Component01 interface {
	Function01(ctx context.Context)
	Function02(ctx context.Context)
}
