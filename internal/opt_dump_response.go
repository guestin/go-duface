package internal

import (
	"bytes"
	"fmt"
	"github.com/guestin/go-requests/opt"
	"io"
)

func DumpResponse(tag string) opt.Option {
	return func(context *opt.RequestContext) error {
		old := context.ResponseHandler
		context.ResponseHandler = func(statusCode int, stream io.Reader) (interface{}, error) {
			fmt.Printf("[%s] status code=%d\n", tag, statusCode)
			respBytes, err := io.ReadAll(stream)
			if err != nil {
				fmt.Printf("[%s] read response error:%s\n", tag, err)
				return nil, err
			}
			fmt.Printf("[%s] response:%s\n", tag, respBytes)
			if old != nil {
				return old(statusCode, bytes.NewReader(respBytes))
			}
			return nil, nil
		}
		return nil
	}

}
