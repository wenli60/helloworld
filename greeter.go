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
	//token, err := util.GenerateToken("wenli", "a3b50661fa14c198813711b29acef97e")
	rsp := &pb.HelloReply{}
	tokenClaims, err := util.Auth("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBpZCI6ImEzYjUwNjYxZmExNGMxOTg4MTM3MTFiMjlhY2VmOTdlIiwiYXBwa2V5IjoiMzU5MjdmOTM4NDllZGFhNGVmODYyMjMxOGY3ZTI3MmQiLCJleHAi\nOjE3MDM3MDAzMDksImlzcyI6Imdpbi1ibG9nIn0.iQyEFAZThIVKpMVUIZSGFSychGAKy1HuvQe26bj1JkM")
	if err != nil {
		panic(err)
	}
	rsp.Msg = tokenClaims.Appid
	return rsp, nil
}
