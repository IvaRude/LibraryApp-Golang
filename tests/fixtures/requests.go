package fixtures

import (
	"homework-3/tests/states"
	"net/http"
	"strings"
)

type RequestBuilder struct {
	req *http.Request
}

func NewRequestBuilder() *RequestBuilder {
	r, _ := http.NewRequest("GET", "/author/1", strings.NewReader(""))
	return &RequestBuilder{req: r}
}

func (b *RequestBuilder) Method(m string) *RequestBuilder {
	b.req.Method = m
	return b
}

func (b *RequestBuilder) P() *http.Request {
	return b.req
}

func (b *RequestBuilder) V() http.Request {
	return *b.req
}

func (b *RequestBuilder) Valid() *RequestBuilder {
	return NewRequestBuilder().Method(states.MethodGet)
}
