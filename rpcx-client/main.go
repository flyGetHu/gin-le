package main

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"log"
	"model"
	"rpcx-client/rpcx_client"
	"time"
)

func main() {
	args := &model.Args{
		A: 10,
		B: 20,
	}

	reply := &model.Reply{}
	calls := make(chan *client.Call, 1)
	for {
		_, err := rpcx_client.RpcxClient().Go(context.Background(), "Mul", args, reply, calls)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}
		replyCall := <-calls
		if replyCall.Error != nil {
			log.Fatalf("failed to call: %v", replyCall.Error)
		} else {
			log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		}
		<-time.After(time.Second)
	}
}
