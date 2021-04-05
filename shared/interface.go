package shared

import (
	"context"
	"github.com/git-jock/jock-cli/proto"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

var HandShake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "JOCK_PLUGIN",
	MagicCookieValue: "woof",
}

var PluginMap = map[string]plugin.Plugin{
	"grpcPlugin": &JPGRPCPlugin{},
}

type JP interface {
	Run(*proto.RunRequest) ([]string, error)
}

type JPGRPCPlugin struct {
	plugin.Plugin
	Impl JP
}

func (p *JPGRPCPlugin) GRPCServer(_ *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterJPServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *JPGRPCPlugin) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewJPClient(c)}, nil
}
