package main

import (
	"context"
	"git.woa.com/test/helloworld/pkg/util"
	"time"

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
	time.Sleep(time.Second * 3)
	err := util.Auth(token)
	rsp := &pb.HelloReply{}
	if err != nil {
		rsp.Msg = err.Error()
		return rsp, nil
	}
	rsp.Msg = "dsdsd"
	return rsp, nil
}
