package internal

import (
	"fmt"
	"github.com/guestin/go-requests/opt"
	"io"
)

func DumpResponse(tag string, skip bool) opt.Option {
	return func(reqCtx *opt.RequestContext) error {
		reqCtx.InstallResponseHandler(
			func(statusCode int, stream io.Reader, previousValue interface{}) (interface{}, error) {
				if skip {
					return previousValue, nil
				}
				fmt.Printf("[%s] status code=%d\n", tag, statusCode)
				respBytes, err := io.ReadAll(stream)
				if err != nil {
					fmt.Printf("[%s] read response error:%s\n", tag, err)
					return nil, err
				}
				fmt.Printf("[%s] response:%s\n", tag, respBytes)
				return nil, nil
			}, opt.HEAD)
		return nil
	}
}
