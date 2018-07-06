package curl

import (
	"net/http"
	"time"
	"io"
	"strings"
	"errors"
	"github.com/xiaomeng79/go-utils/httpclient"
	"io/ioutil"
)

//这些是要记录到日志的东西
type Curl struct {
	Method string `json:"method"`
	Url string `json:"url"`
	StatusCode int `json:"status"`
	Delay int64 `json:"delay"`
	Request *Request `json:"request"`
	Response *Response `json:"response"`
}

//请求
type Request struct {
	Header http.Header
	Body string
	ContentType string
	ContentLength int64
}

//响应
type Response struct {
	Header http.Header
	Body string
	ContentType string
	ContentLength int64
}


//定义错误
var (
	NotAllowMethod = errors.New("method not allow")
)

type ICurl interface {
	Do() error //去请求
}

//新建一个curl
func New() *Curl {
	req := &Request{
		Header:make(http.Header),
	}
	res := &Response{
		Header:make(http.Header),
	}
	return &Curl{
		Request:req,
		Response:res,
	}
}
//注意调用者做异常处理
func (c *Curl) Do() error {
	//记录请求时间,毫秒
	s_time := time.Now().UnixNano()/1e6
	defer func() {
		c.Delay = time.Now().UnixNano()/1e6 - s_time
	}()
	//生成请求内容
	var input io.Reader
	switch c.Method {
	case "GET","OPTIONS","HEAD","TRACE","CONNECT":
		input = nil
	case "POST","PUT","DELETE":
		input = strings.NewReader(c.Request.Body)
	default :
		return NotAllowMethod
	}
	//生成请求
	req, err := http.NewRequest(c.Method,c.Url,input)
	if err != nil {
		return err
	}
	c.Request.ContentLength = req.ContentLength
	req.Header = c.Request.Header
	c.Request.ContentType = req.Header.Get("Content-Type")
	//执行请求
	//1.生成请求客户端
	hc := httpclient.New()
	res,err := hc.Do(req)
	if err != nil {
		return err
	}
	//读取响应体
	body,err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	c.StatusCode = res.StatusCode
	c.Response.Header = res.Header
	c.Response.ContentLength = res.ContentLength
	c.Response.ContentType = res.Header.Get("Content-Type")
	c.Response.Body = string(body)
	defer res.Body.Close()
	return nil

}
//设置content-type
func (c *Curl) SetContentType(contentType string) {
	c.Request.ContentType = contentType
}

//设置body
func (c *Curl) SetBody(body string) {
	c.Request.Body = body
}
//设置method
func (c *Curl) SetMethod(method string) {
	c.Method = method
}
//设置url
func (c *Curl) SetUrl(url string) {
	c.Url = url
}
//增加头信息
func (c *Curl) AddHeader(k,v string) {
	c.Request.Header.Add(k,v)
}
//设置头信息
func (c *Curl) SetHeader(k,v string) {
	c.Request.Header.Set(k,v)
}
