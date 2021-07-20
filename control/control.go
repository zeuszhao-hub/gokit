package control

import (
	"git.neigou.com/zhaofuchun/gokit/interfaces/iserver"
)

type Control struct {
	svr []iserver.Server
}

func NewControl() *Control {
	return &Control{}
}

// RegisterServer 注册服务接管，先注册的将被优先RUN和优先Shutdown
func (a *Control) RegisterServer(svr ...iserver.Server) {
	for _, v := range svr {
		a.svr = append(a.svr, v)
	}
}

func (a *Control) Run() {
	for _, v := range a.svr {
		v.Run()
	}
}

func (a *Control) Shutdown() {
	for _, v := range a.svr {
		v.Shutdown()
	}
}
