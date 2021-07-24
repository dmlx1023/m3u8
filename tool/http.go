package tool

import (
	"bytes"
	"fmt"
	"github.com/valyala/fasthttp"
	"io"
)

func Get(url string) ( io.Reader, error ) {

	state, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		return nil, err
	}
	if state != fasthttp.StatusOK{
		return nil, fmt.Errorf("http error: status code %d", state )
	}
	return  bytes.NewReader(resp), nil
}
