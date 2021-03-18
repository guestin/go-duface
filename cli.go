package duface

import "context"

type (
	_LibraryConfig struct {
		ctx context.Context
	}

	LibraryOpt func(opt *_LibraryConfig)

	Client interface {

		// 创建库实体
		NewLibrary(groupId string, opts ...LibraryOpt) (Library, error)

		// 读取库列表
		// @param offset 默认0,起始序列号
		// @param length 默认100,最大为1000
		ListLibraries(offset, length int) ([]string, error)

		// 人脸检测
		FaceDetect(imgData *ImageData, params *DetectExtParams) (*FaceDetectResult, error)

		// get current context
		GetContext() context.Context

		// get current access-token
		GetAccessToken() (string, error)
	}
)

func WithContext(ctx context.Context) LibraryOpt {
	return func(opt *_LibraryConfig) {
		opt.ctx = ctx
	}
}
