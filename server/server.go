package server

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/syzhang42/go-fire/auth"
	"golang.org/x/sync/errgroup"
)

type ServerConfig struct {
	ServerCfg struct {
		Servers []string `toml:"servers" required:"true"`
	} `toml:"server" required:"true"`
}

var serCfg ServerConfig

func Run(cfgStr string) {
	_, err := toml.DecodeFile(cfgStr, &serCfg)
	auth.Must(err)

	sigctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cancel()
	egrp, ctx := errgroup.WithContext(sigctx)

	egrp.SetLimit(len(serCfg.ServerCfg.Servers))
	for _, serverName := range serCfg.ServerCfg.Servers {
		if server, ok := serverMap[serverName]; ok && server != nil {
			if server.Name() != serverName {
				panic(fmt.Sprintf("%v not Register", serverName))
			}
			server.Init(cfgStr)
		}
		panic(fmt.Sprintf("%v not Register", serverName))
	}

	for _, server := range serverMap {
		_server := server
		egrp.Go(func() error { return _server.Run(ctx) })
	}
	egrp.Wait()
}
