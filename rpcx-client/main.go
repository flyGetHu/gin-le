package main

import (
	"context"
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
	for {
		call, err := rpcx_client.RpcxClient().Go(context.Background(), "Mul", args, reply, nil)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		replyCall := <-call.Done
		if replyCall.Error != nil {
			log.Fatalf("failed to call: %v", replyCall.Error)
		} else {
			log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		}
		<-time.After(time.Second)
	}
}
