package duface

import "context"

type (
	_LibraryConfig struct {
		ctx context.Context
	}
	LibraryOpt func(opt *_LibraryConfig)
)

func WithContext(ctx context.Context) LibraryOpt {
	return func(opt *_LibraryConfig) {
		opt.ctx = ctx
	}
}
