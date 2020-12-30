package builder

import (
	"io"
	"net/http"
	"net/url"
)

type builder struct {
	req    *http.Request
	method string
	url    *url.URL
	body   io.Reader
}

type Option func(builder *builder)

func New(method, host string, options ...Option) (*builder, error) {
	u, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	ret := &builder{method: method, url: u}
	for _, opt := range options {
		opt(ret)
	}
	req, err := http.NewRequest(ret.method, ret.url.String(), ret.body)
	if err != nil {
		return nil, err
	}
	ret.req = req
	return ret, nil
}

func WithBody(body io.Reader) Option {
	return func(builder *builder) {
		builder.body = body
	}
}

func WithQuery(query map[string]string) Option {
	return func(builder *builder) {
		parameters := url.Values{}
		for k, v := range query {
			parameters.Add(k, v)
		}
		builder.url.RawQuery = parameters.Encode()
	}
}

func (b *builder) Do() (resp *http.Response, err error) {
	client := http.Client{}
	return client.Do(b.req)
}
