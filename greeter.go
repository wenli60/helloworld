package main

import (
	"context"
	"git.woa.com/test/helloworld/pkg/util"

	pb "git.woa.com/trpcprotocol/test/helloworld"
)

type greeterImpl struct {
	pb.UnimplementedGreeter
}

func (s *greeterImpl) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	//鉴权
	rsp := &pb.HelloReply{}
	tokenClaims, err := util.Auth(req.Token)
	if err != nil {
		rsp.Msg = err.Error()
		return rsp, nil
	}
	rsp.Msg = tokenClaims.Appid
	return rsp, nil
}
