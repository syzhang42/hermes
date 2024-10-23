package httpserver

import (
	"context"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/syzhang42/go-fire/auth"
	"github.com/syzhang42/hermes/server"
	"github.com/syzhang42/hermes/server/httpserver/internal"
)

type HttpServer struct {
	Config struct {
		Ip   string `toml:"ip"`
		Port string `toml:"port" default:"9958"`
	} `toml:"http_server"`

	api *internal.ApiManager
}

func init() {
	server.Register(&HttpServer{})
}

func (h *HttpServer) Name() string {
	return "http_server"
}

func (h *HttpServer) Init(cfgStr string) {
	_, err := toml.DecodeFile(cfgStr, h)
	auth.Must(err)
	h.api = internal.NewApiManager()
	h.api.Init()
}

func (h *HttpServer) Run(ctx context.Context) error {
	h.api.Run(fmt.Sprintf("%v:%v", h.Config.Ip, h.Config.Port))
	return nil
}
