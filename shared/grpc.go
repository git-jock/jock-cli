package shared

import (
	"github.com/git-jock/jock-cli/proto"
	"golang.org/x/net/context"
)

type GRPCClient struct {
	client proto.JPClient
}

func (m *GRPCClient) Run(args []string) ([]string, error) {
	resp, err := m.client.Run(context.Background(), &proto.RunRequest{
		Args: args,
	})
	return resp.Log, err
}

type GRPCServer struct {
	proto.UnimplementedJPServer
	Impl JP
}

func (m *GRPCServer) Run(_ context.Context, req *proto.RunRequest) (*proto.RunResponse, error) {
	l, err := m.Impl.Run(req.Args)
	return &proto.RunResponse{Log: l}, err
}
