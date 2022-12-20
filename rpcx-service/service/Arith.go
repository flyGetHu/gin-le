package service

import (
	"context"
	"model"
	"time"
)

type Arith int

func (t *Arith) Mul(ctx context.Context, args *model.Args, reply *model.Reply) error {
	reply.C = args.A * args.B
	reply.C = time.Now().Second()
	return nil
}
