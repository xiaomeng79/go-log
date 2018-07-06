//build pattern
package curl

type Builder interface {
	SetUrl(url string) Builder
	SetMethod(method string) Builder
	SetContentType(contentType string) Builder
	SetBody(body string) Builder
	SetHeader(k,v string) Builder
	AddHeader(k,v string) Builder
	Build() ICurl
}

type CurlBuilder struct {
	Curl *Curl
}

func (c *CurlBuilder) SetUrl(url string) Builder {
	if c.Curl == nil {
		c.Curl = New()
	}
	c.Curl.SetUrl(url)
	return c
}

func (c *CurlBuilder) SetMethod(method string) Builder {
	if c.Curl == nil {
		c.Curl = New()
	}
	c.Curl.SetMethod(method)
	return c
}

func (c *CurlBuilder) SetContentType(contentType string) Builder {
	if c.Curl == nil {
		c.Curl = New()
	}
	c.Curl.SetContentType(contentType)
	return c
}

func (c *CurlBuilder) SetBody(body string) Builder {
	if c.Curl == nil {
		c.Curl = New()
	}
	c.Curl.SetBody(body)
	return c
}

func (c *CurlBuilder) SetHeader(k,v string) Builder {
	if c.Curl == nil {
		c.Curl = New()
	}
	c.Curl.SetHeader(k,v)
	return c
}

func (c *CurlBuilder) AddHeader(k,v string) Builder {
	if c.Curl == nil {
		c.Curl = New()
	}
	c.Curl.AddHeader(k,v)
	return c
}

func (c *CurlBuilder) Build() ICurl {
	return c.Curl
}


