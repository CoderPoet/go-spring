package SpringGoRestful

import (
	"context"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/go-spring/go-spring/spring-web"
)

type Container struct {
	HttpServer *http.Server
	WebServer  *restful.WebService
}

func (c *Container) Stop() {
	c.HttpServer.Shutdown(context.TODO())
}

func (c *Container) Start(address string) error {
	c.HttpServer = &http.Server{Addr: address, Handler: restful.DefaultContainer}
	return c.HttpServer.ListenAndServe()
}

func (c *Container) StartTLS(address string, certFile, keyFile string) error {
	c.HttpServer = &http.Server{Addr: address, Handler: restful.DefaultContainer}
	return c.HttpServer.ListenAndServeTLS(certFile, keyFile)
}

func (c *Container) GET(path string, fn SpringWeb.Handler, filters ...SpringWeb.Filter) {
	c.WebServer.GET(path).Filter(HandlerWrapper(fn, filters...))
}

func (c *Container) POST(path string, fn SpringWeb.Handler, filters ...SpringWeb.Filter) {
	c.WebServer.POST(path).Filter(HandlerWrapper(fn, filters...))
}

func (c *Container) PATCH(path string, fn SpringWeb.Handler, filters ...SpringWeb.Filter) {
	c.WebServer.PATCH(path).Filter(HandlerWrapper(fn, filters...))
}

func (c *Container) PUT(path string, fn SpringWeb.Handler, filters ...SpringWeb.Filter) {
	c.WebServer.PUT(path).Filter(HandlerWrapper(fn, filters...))
}

func (c *Container) DELETE(path string, fn SpringWeb.Handler, filters ...SpringWeb.Filter) {
	c.WebServer.DELETE(path).Filter(HandlerWrapper(fn, filters...))
}

func (c *Container) HEAD(path string, fn SpringWeb.Handler, filters ...SpringWeb.Filter) {
	c.WebServer.HEAD(path).Filter(HandlerWrapper(fn, filters...))
}

func (c *Container) OPTIONS(path string, fn SpringWeb.Handler, filters ...SpringWeb.Filter) {
	// go-restful 并没有该方法
	panic("implement me")
}

func NewContainer() *Container {
	ws := new(restful.WebService)
	restful.Add(ws)

	return &Container{
		WebServer: ws,
	}
}

func HandlerWrapper(fn SpringWeb.Handler, filters ...SpringWeb.Filter) restful.FilterFunction {
	return func(req *restful.Request, res *restful.Response, chain *restful.FilterChain) {

	}
}
