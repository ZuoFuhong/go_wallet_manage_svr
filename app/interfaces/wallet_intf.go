package interfaces

import (
	"context"
	"github.com/ZuoFuhong/go_wallet_manage_svr/errcode"
	"github.com/ZuoFuhong/go_wallet_manage_svr/pkg/log"
	pb "github.com/ZuoFuhong/grpc-standard-pb/go_wallet_manage_svr"
	"github.com/google/uuid"
)

// CreateWallet 创建钱包
func (s *GoWalletManageSvrImpl) CreateWallet(ctx context.Context, _ *pb.CreateWalletReq) (*pb.CreateWalletRsp, error) {
	address := uuid.New().String()
	log.DebugContextf(ctx, "create new wallet address = %s", address)
	rsp := &pb.CreateWalletRsp{
		Address: address,
	}
	return rsp, nil
}

// ImportWallet 导入钱包
func (s *GoWalletManageSvrImpl) ImportWallet(ctx context.Context, req *pb.ImportWalletReq) (*pb.ImportWalletRsp, error) {
	privKey := req.GetPrivateKey()
	log.DebugContextf(ctx, "Import wallet privKey: %s", privKey)
	if privKey == "" {
		return nil, errcode.ErrLogicParam
	}
	rsp := &pb.ImportWalletRsp{
		Address: uuid.New().String(),
	}
	return rsp, nil
}
