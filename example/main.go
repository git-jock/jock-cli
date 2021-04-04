package main

import (
	"github.com/git-jock/jock-cli/shared"
	"github.com/hashicorp/go-plugin"
)

type JP struct{}

func (JP) Run(args []string) ([]string, error) {
	return append(args, "changed"), nil
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
