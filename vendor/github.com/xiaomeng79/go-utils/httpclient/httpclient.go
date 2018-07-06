//使用单例和选项设计模式写一个http请求客户端
package httpclient

import (
	"net/http"
	"sync"
	"time"
)
//选项
type Option func(*Options)
type Options struct {
	Timeout time.Duration //总体超时时间
}
//设置超时
func SetTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.Timeout = timeout
	}
}

var (
	once sync.Once
	instance *http.Client
)
//创建一个单例的http客户端
func New(opts ...Option) *http.Client {
	defaultOptions := &Options{
		Timeout:30 * time.Second,//默认30s超时
	}
	for _,opt := range opts {
		opt(defaultOptions)
	}
	once.Do(func() {
		instance = &http.Client{
			Timeout:defaultOptions.Timeout,
			Transport:http.DefaultTransport,
		}
	})
	return instance
}


