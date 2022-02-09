package handlers

import (
	"io"

	"github.com/sirupsen/logrus"

	"net/http"
)

type Proxy struct {
	logger *logrus.Logger
}

func NewProxy(logger *logrus.Logger) *Proxy {
	return &Proxy{logger}
}

//nolint:varnamelen
func (p *Proxy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	resp, err := p.HandleHTTP(r)
	if err != nil {
		p.logger.Error(err)
	}
	defer resp.Body.Close()

	for header, values := range resp.Header {
		for _, value := range values {
			rw.Header().Add(header, value)
		}
	}
	rw.WriteHeader(resp.StatusCode)

	if _, err = io.Copy(rw, resp.Body); err != nil {
		p.logger.Error(err)
	}
}
