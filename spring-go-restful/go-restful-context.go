package SpringGoRestful

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/emicklei/go-restful"
	SpringTrace "github.com/go-spring/go-spring/spring-trace"
	SpringWeb "github.com/go-spring/go-spring/spring-web"
)

const (
	HeaderUpgrade            = "Upgrade"
	HeaderContentType        = "Content-Type"
	HeaderContentDisposition = "Content-Disposition"
	HeaderXForwardedProto    = "X-Forwarded-Proto"
	HeaderXForwardedProtocol = "X-Forwarded-Protocol"
	HeaderXForwardedSsl      = "X-Forwarded-Ssl"
	HeaderXUrlScheme         = "X-Url-Scheme"
)

type Context struct {
	*SpringTrace.DefaultTraceContext

	// go-restful 上下文对象
	GoRestfulRequest  *restful.Request
	GoRestfulResponse *restful.Response

	// 处理器 Path
	HandlerPath string

	// Web 处理函数
	HandlerFun SpringWeb.Handler

	paramNames  []string
	paramValues []string
}

func (c *Context) NativeContext() interface{} {
	return c.GoRestfulRequest
}

func (c *Context) Get(key string) interface{} {
	return c.GoRestfulRequest.Attribute(key)
}

func (c *Context) Set(key string, val interface{}) {
	c.GoRestfulRequest.SetAttribute(key, val)
}

func (c *Context) Request() *http.Request {
	return c.GoRestfulRequest.Request
}

func (c *Context) IsTLS() bool {
	return c.GoRestfulRequest.Request.TLS != nil
}

func (c *Context) IsWebSocket() bool {

	// NOTE: 这一段逻辑使用 gin 的实现
	if strings.Contains(strings.ToLower(c.GoRestfulRequest.HeaderParameter("Connection")), "upgrade") &&
		strings.ToLower(c.GoRestfulRequest.HeaderParameter("Upgrade")) == "websocket" {
		return true
	}
	return false

	// NOTE: 这一段逻辑使用 echo 的实现
	//upgrade := c.GoRestfulRequest.Request.Header.Get(HeaderUpgrade)
	//return upgrade == "websocket" || upgrade == "Websocket"
}

func (c *Context) Scheme() string {
	r := c.GoRestfulRequest.Request
	// Can't use `r.Request.URL.Scheme`
	// See: https://groups.google.com/forum/#!topic/golang-nuts/pMUkBlQBDF0

	if r.TLS != nil {
		return "https"
	}

	if scheme := r.Header.Get(HeaderXForwardedProto); scheme != "" {
		return scheme
	}

	if scheme := r.Header.Get(HeaderXForwardedProtocol); scheme != "" {
		return scheme
	}

	if ssl := r.Header.Get(HeaderXForwardedSsl); ssl == "on" {
		return "https"
	}

	if scheme := r.Header.Get(HeaderXUrlScheme); scheme != "" {
		return scheme
	}
	return "http"
}

func (c *Context) ClientIP() string {
	// TODO 参考gin的实现
	panic("implement me")
}

func (c *Context) Path() string {
	return c.HandlerPath
}

func (c *Context) Handler() SpringWeb.Handler {
	return c.HandlerFun
}

func (c *Context) ContentType() string {
	return c.GoRestfulRequest.HeaderParameter("Content-Type")
}

func (c *Context) GetHeader(key string) string {
	return c.GoRestfulRequest.HeaderParameter(key)
}

func (c *Context) GetRawData() ([]byte, error) {
	panic("implement me")
}

func (c *Context) PathParam(name string) string {
	panic("implement me")
}

func (c *Context) PathParamNames() []string {
	panic("implement me")
}

func (c *Context) PathParamValues() []string {
	panic("implement me")
}

func (c *Context) QueryParam(name string) string {
	panic("implement me")
}

func (c *Context) QueryParams() url.Values {
	panic("implement me")
}

func (c *Context) QueryString() string {
	panic("implement me")
}

func (c *Context) FormValue(name string) string {
	panic("implement me")
}

func (c *Context) FormParams() (url.Values, error) {
	panic("implement me")
}

func (c *Context) FormFile(name string) (*multipart.FileHeader, error) {
	panic("implement me")
}

func (c *Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	panic("implement me")
}

func (c *Context) MultipartForm() (*multipart.Form, error) {
	panic("implement me")
}

func (c *Context) Cookie(name string) (*http.Cookie, error) {
	panic("implement me")
}

func (c *Context) Cookies() []*http.Cookie {
	panic("implement me")
}

func (c *Context) Bind(i interface{}) error {
	panic("implement me")
}

func (c *Context) Status(code int) {
	panic("implement me")
}

func (c *Context) Header(key, value string) {
	panic("implement me")
}

func (c *Context) SetCookie(cookie *http.Cookie) {
	panic("implement me")
}

func (c *Context) NoContent(code int) {
	panic("implement me")
}

func (c *Context) String(code int, format string, values ...interface{}) {
	panic("implement me")
}

func (c *Context) HTML(code int, html string) {
	panic("implement me")
}

func (c *Context) HTMLBlob(code int, b []byte) {
	panic("implement me")
}

func (c *Context) JSON(code int, i interface{}) {
	panic("implement me")
}

func (c *Context) JSONPretty(code int, i interface{}, indent string) {
	panic("implement me")
}

func (c *Context) JSONBlob(code int, b []byte) {
	panic("implement me")
}

func (c *Context) JSONP(code int, callback string, i interface{}) {
	panic("implement me")
}

func (c *Context) JSONPBlob(code int, callback string, b []byte) {
	panic("implement me")
}

func (c *Context) XML(code int, i interface{}) {
	panic("implement me")
}

func (c *Context) XMLPretty(code int, i interface{}, indent string) {
	panic("implement me")
}

func (c *Context) XMLBlob(code int, b []byte) {
	panic("implement me")
}

func (c *Context) Blob(code int, contentType string, b []byte) {
	panic("implement me")
}

func (c *Context) Stream(code int, contentType string, r io.Reader) {
	panic("implement me")
}

func (c *Context) File(file string) {
	panic("implement me")
}

func (c *Context) Attachment(file string, name string) {
	panic("implement me")
}

func (c *Context) Inline(file string, name string) {
	panic("implement me")
}

func (c *Context) Redirect(code int, url string) {
	panic("implement me")
}

func (c *Context) SSEvent(name string, message interface{}) {
	panic("implement me")
}

func (c *Context) Error(err error) {
	panic("implement me")
}
