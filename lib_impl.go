package duface

type _LibraryImpl struct {
	cli     Client
	groupId string
}

func newLibrary(cli Client, groupId string) Library {
	return &_LibraryImpl{
		cli:     cli,
		groupId: groupId,
	}
}

func (this *_LibraryImpl) RegisterFace(userId string, imgData *ImageData, extParams *RegExtParams) (*RegisterFaceResponse, error) {
	panic("implement me")
}

func (this *_LibraryImpl) DeleteFace(userId string, faceToken string) (*BasicResponse, error) {
	panic("implement me")
}

func (this *_LibraryImpl) DeleteUser(userId string) (*BasicResponse, error) {
	panic("implement me")
}

func (this *_LibraryImpl) Search(imgData *ImageData) (*SearchResponse, error) {
	panic("implement me")
}

func (this *_LibraryImpl) Drop() error {
	panic("implement me")
}
