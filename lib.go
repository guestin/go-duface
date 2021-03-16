package duface

type Library interface {

	// 人脸入库
	RegisterFace(userId string,
		imgData *ImageData,
		extParams *RegExtParams) (*RegisterFaceResponse, error)

	// 删除单个脸
	DeleteFace(userId string, faceToken string) (*BasicResponse, error)

	// 删除用户
	DeleteUser(userId string) (*BasicResponse, error)

	// 在指定分组中查找人脸
	Search(imgData *ImageData) (*SearchResponse, error)

	// 删库
	Drop() error
}
