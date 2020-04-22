package grpc

import (
	"github.com/open-kingfisher/king-utils/common/log"
	"google.golang.org/grpc"
)

func ClientDial(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc server not connect: %v", err)
	}
	return conn
}
