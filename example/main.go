package main

import (
	"github.com/git-jock/jock-cli/proto"
	"github.com/git-jock/jock-cli/shared"
	"github.com/hashicorp/go-plugin"
)

type JP struct{}

func (JP) Run(req *proto.RunRequest) ([]string, error) {
	return append(req.Args, "changed"), nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.HandShake,
		Plugins: map[string]plugin.Plugin{
			"jk": &shared.JPGRPCPlugin{Impl: &JP{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
