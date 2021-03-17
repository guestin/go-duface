package duface

type Client interface {

	// 创建库实体
	NewLibrary(groupId string) (Library, error)

	// 读取库列表
	// @param offset 默认0,起始序列号
	// @param length 默认100,最大为1000
	ListLibraries(offset, length int) ([]string, error)

	// 人脸检测
	FaceDetect(imgData *ImageData, params *DetectExtParams) (*FaceDetectResult, error)
}
