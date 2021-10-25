package http

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeuszhao-hub/gokit/interfaces/iserver"
	"net/http"
	"time"
)

type Ihttpctx interface {
	iserver.Server
}

type httpctx struct {
	ctx context.Context
	srv *http.Server
}

type Options struct {
	Address string
	Port    int
}

func NewHttp(o Options, h http.Handler) Ihttpctx {

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", o.Address, o.Port),
		Handler: h,
	}

	return &httpctx{
		srv: srv,
		ctx: context.Background(),
	}
}

func (hctx *httpctx) Run() error {

	go func() error {
		if err := hctx.srv.ListenAndServe(); err != nil && errors.Cause(err) == http.ErrServerClosed {
			fmt.Errorf(err.Error())
			return err
		}
		return nil
	}()

	return nil
}

func (hctx *httpctx) Shutdown() error {
	ctx, cancel := context.WithTimeout(hctx.ctx, 30*time.Second)
	defer cancel()
	if err := hctx.srv.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
