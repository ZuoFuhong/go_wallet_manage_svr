package main

import (
	"context"
	"fmt"
	codec "github.com/ZuoFuhong/grpc-middleware/encoding/json"
	pb "github.com/ZuoFuhong/grpc-standard-pb/go_wallet_manage_svr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"testing"
)

func Test_RPC(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:1025", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	stub := pb.NewGoWalletManageSvrClient(conn)
	rpcRsp, err := stub.ImportWallet(context.Background(), &pb.ImportWalletReq{
		PrivateKey: "0x01c4bda0939df07a31e3738c6c1e1d5905c9f229e6ffa1922557308a62efb23f",
	}, grpc.CallContentSubtype(codec.Name)) // 指定 JSON 编码
	if err != nil {
		// Resolve grpc errcode
		if rpcErr, ok := status.FromError(err); ok {
			fmt.Println(rpcErr)
		}
		return
	}
	fmt.Println("rpcRsp: ", rpcRsp.GetAddress())
}

func Test_LoadBalance(t *testing.T) {
	conn, err := grpc.Dial("monica://Production/go_wallet_manage_svr",
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"weighted_round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatal(err)
	}
	stub := pb.NewGoWalletManageSvrClient(conn)
	rpcRsp, err := stub.ImportWallet(context.Background(), &pb.ImportWalletReq{
		PrivateKey: "0x01c4bda0939df07a31e3738c6c1e1d5905c9f229e6ffa1922557308a62efb23f",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("rpcRsp: ", rpcRsp.GetAddress())
}
