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
	jwtPayLoad := util.JwtPayLoad{
		Appid:  "wenli",
		Appkey: "a3b50661fa14c198813711b29acef97e",
	}
	token, _ := util.GenToken(jwtPayLoad)
	err := util.Auth(token)
	rsp := &pb.HelloReply{}
	if err != nil {
		panic(err)
	}
	rsp.Msg = "dsdsd"
	return rsp, nil
}
