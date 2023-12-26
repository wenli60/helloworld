package main

import (
	"context"

	pb "git.woa.com/trpcprotocol/test/helloworld"
)

type greeterImpl struct {
	pb.UnimplementedGreeter
}

func (s *greeterImpl) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	//鉴权

	rsp := &pb.HelloReply{}
	rsp.Msg = "ddddddddd"
	return rsp, nil
}
