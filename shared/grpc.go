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
	if resp != nil {
		return resp.Log, err
	}
	return nil, err
}

type GRPCServer struct {
	proto.UnimplementedJPServer
	Impl JP
}

func (m *GRPCServer) Run(ctx context.Context, req *proto.RunRequest) (*proto.RunResponse, error) {
	l, err := m.Impl.Run(req.Args)
	if l != nil {
		return &proto.RunResponse{Log: l}, err
	}
	return nil, err
}
