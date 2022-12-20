package rpcx_client

import (
	"config"
	"github.com/smallnest/rpcx/client"
)

var rpcxClient client.XClient

func RpcxClient() client.XClient {
	if rpcxClient == nil {
		d, _ := client.NewPeer2PeerDiscovery(config.RpcxAddress, "")
		rpcxClient = client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	}
	return rpcxClient
}
