package main

import (
	_ "git.code.oa.com/trpc-go/trpc-filter/debuglog"
	_ "git.code.oa.com/trpc-go/trpc-filter/recovery"
	trpc "git.code.oa.com/trpc-go/trpc-go"
	"git.code.oa.com/trpc-go/trpc-go/log"
	"git.woa.com/test/helloworld/internal/models"
	pb "git.woa.com/trpcprotocol/test/helloworld"
)

func main() {
	s := trpc.NewServer()
	err := models.Setup()
	if err != nil {
		panic(err)
	}
	pb.RegisterGreeterService(s.Service("trpc.test.helloworld.Greeter"), &greeterImpl{})
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
