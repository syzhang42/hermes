package server

/*
如何扩展server?
1、func init 调用 Register上报你的 server实例
2、提供Init接口初始化你的server实例
3、提供Run-->阻塞
*/
import (
	"context"
	"sync"
)

type Server interface {
	Name() string
	Init(cfgStr string)
	Run(ctx context.Context) error
}

var (
	serverMap  = make(map[string]Server, 0)
	providerMu = sync.Mutex{}
)

func Register(srv Server) {
	providerMu.Lock()
	defer providerMu.Unlock()
	if srv == nil {
		return
	}
	serverMap[srv.Name()] = srv
}
