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
	var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImEzYjUwNjYxZmExNGMxOTg4MTM3MTFiMjlhY2VmOTdlIiwicGFzc3dvcmQiOiIzNTkyN2Y5Mzg0OWVkYWE0ZWY4NjIyMzE4ZjdlMjcyZCIs\nImV4cCI6MTcwMzYwOTU2MSwiaXNzIjoiZ2luLWJsb2cifQ._9R6b3B_uj8DO2K6SaZ2dCaO_g81h4ig8wXhACzdGQg"

	tokenClaims, err := util.Auth(token)
	//token, err := util.GenerateToken("wenli", "a3b50661fa14c198813711b29acef97e")
	if err != nil {
		panic(err)
	}
	rsp := &pb.HelloReply{}
	rsp.Msg = tokenClaims.Username
	return rsp, nil
}
