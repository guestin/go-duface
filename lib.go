package duface

type Library interface {

	// 建立库
	Create() error

	// 删库
	Drop() error

	// 人脸入库
	RegisterFace(userId string,
		imgData *ImageData,
		extParams *RegExtParams) (*RegisterFaceResult, error)

	// 删除单个脸
	DeleteFace(userId string, faceToken string) error

	// 用户人脸列表
	ListUserFaces(userId string) ([]*UserFaceItem, error)

	// 删除用户
	DeleteUser(userId string) error

	// 用户id列表
	ListUsers(start, length uint) ([]string, error)

	// 1:N 在指定分组中查找人脸
	Search(imgData *ImageData,
		extParam *SearchExtParams) ([]*ComparisonInfo, error)

	// M:N 在指定分组中查找人脸
	MultiSearch(imgData *ImageData,
		extParam *MultiSearchExtParams) ([]*MultiSearchResultItem, error)
}
