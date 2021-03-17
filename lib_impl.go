package duface

import (
	"github.com/guestin/go-duface/internal"
	"github.com/guestin/go-requests"
	"github.com/guestin/go-requests/opt"
	"github.com/guestin/mob/murl"
	"strings"
)

type _LibraryImpl struct {
	cli     *_ClientImpl
	groupId string
}

func newLibrary(cli *_ClientImpl, groupId string) Library {
	return &_LibraryImpl{
		cli:     cli,
		groupId: groupId,
	}
}

func (this *_LibraryImpl) RegisterFace(userId string, imgData *ImageData, extParams *RegExtParams) (*RegisterFaceResult, error) {
	accessToken, err := this.cli.getAccessToken()
	if err != nil {
		return nil, err
	}
	_url, err := murl.MakeUrlString(internal.DuFaceBusinessUrlV3,
		murl.WithPath("faceset/user/add"),
		murl.WithQuery("access_token", accessToken))
	req := struct {
		*ImageData
		UserId  string `json:"user_id"`
		GroupId string `json:"group_id"`
		*RegExtParams
	}{
		ImageData:    imgData,
		UserId:       userId,
		GroupId:      this.groupId,
		RegExtParams: extParams,
	}
	rsp := struct {
		BaseResponse
		Result *RegisterFaceResult `json:"result" validate:"omitempty,required"`
	}{}
	if _, err = requests.Post(this.cli.ctx, _url,
		opt.BodyJSON(&req), opt.BindJSON(&rsp)); err != nil {
		return nil, err
	}
	if err := rsp.intoError(); err != nil {
		return nil, err
	}
	return rsp.Result, nil
}

func (this *_LibraryImpl) DeleteFace(userId string, faceToken string) error {
	accessToken, err := this.cli.getAccessToken()
	if err != nil {
		return err
	}
	_url, err := murl.MakeUrlString(internal.DuFaceBusinessUrlV3,
		murl.WithPath("faceset/face/delete"),
		murl.WithQuery("access_token", accessToken))
	req := struct {
		UserId    string `json:"user_id"`
		GroupId   string `json:"group_id"`
		FaceToken string `json:"face_token"`
	}{
		UserId:    userId,
		GroupId:   this.groupId,
		FaceToken: faceToken,
	}
	rsp := BaseResponse{}
	if _, err = requests.Post(this.cli.ctx, _url,
		opt.BodyJSON(&req), opt.BindJSON(&rsp)); err != nil {
		return err
	}
	if err := rsp.intoError(); err != nil {
		return err
	}
	return nil
}

func (this *_LibraryImpl) ListUserFaces(userId string) ([]*UserFaceItem, error) {
	accessToken, err := this.cli.getAccessToken()
	if err != nil {
		return nil, err
	}
	_url, err := murl.MakeUrlString(internal.DuFaceBusinessUrlV3,
		murl.WithPath("faceset/face/getlist"),
		murl.WithQuery("access_token", accessToken))
	req := struct {
		UserId  string `json:"user_id"`
		GroupId string `json:"group_id"`
	}{
		UserId:  userId,
		GroupId: this.groupId,
	}
	rsp := struct {
		BaseResponse
		Result struct {
			UserFaceItems []*UserFaceItem `json:"face_list"`
		} `json:"result"`
	}{}
	_, err = requests.Post(this.cli.ctx, _url,
		opt.BodyJSON(&req), opt.BindJSON(&rsp))
	if err != nil {
		return nil, err
	}
	if err := rsp.intoError(); err != nil {
		return nil, err
	}
	return rsp.Result.UserFaceItems, nil
}

func (this *_LibraryImpl) DeleteUser(userId string) error {
	accessToken, err := this.cli.getAccessToken()
	if err != nil {
		return err
	}
	_url, err := murl.MakeUrlString(internal.DuFaceBusinessUrlV3,
		murl.WithPath("faceset/user/delete"),
		murl.WithQuery("access_token", accessToken))
	req := struct {
		UserId  string `json:"user_id"`
		GroupId string `json:"group_id"`
	}{
		UserId:  userId,
		GroupId: this.groupId,
	}
	rsp := BaseResponse{}
	if _, err = requests.Post(this.cli.ctx, _url,
		opt.BodyJSON(&req), opt.BindJSON(&rsp)); err != nil {
		return err
	}
	if err := rsp.intoError(); err != nil {
		return err
	}
	return nil
}

func (this *_LibraryImpl) ListUsers(start, length uint) ([]string, error) {
	accessToken, err := this.cli.getAccessToken()
	if err != nil {
		return nil, err
	}
	_url, err := murl.MakeUrlString(internal.DuFaceBusinessUrlV3,
		murl.WithPath("faceset/group/getusers"),
		murl.WithQuery("access_token", accessToken))
	req := struct {
		GroupId string `json:"group_id"`
		Start   uint   `json:"start"`
		Length  uint   `json:"length"`
	}{
		GroupId: this.groupId,
		Start:   start,
		Length:  length,
	}
	rsp := struct {
		BaseResponse
		Result struct {
			UserIds []string `json:"user_id_list"`
		} `json:"result"`
	}{}
	_, err = requests.Post(this.cli.ctx, _url,
		opt.BodyJSON(&req), opt.BindJSON(&rsp))
	if err != nil {
		return nil, err
	}
	if err := rsp.intoError(); err != nil {
		return nil, err
	}
	return rsp.Result.UserIds, nil
}

func (this *_LibraryImpl) Search(imgData *ImageData, extParam *SearchExtParams) ([]*ComparisonInfo, error) {
	accessToken, err := this.cli.getAccessToken()
	if err != nil {
		return nil, err
	}
	_url, err := murl.MakeUrlString(internal.DuFaceBusinessUrlV3,
		murl.WithPath("search"),
		murl.WithQuery("access_token", accessToken))
	req := struct {
		Image       string     `json:"image"`
		ImageType   ImageTypes `json:"image_type"`
		GroupIdList string     `json:"group_id_list"`
		*SearchExtParams
	}{
		Image:           imgData.Data,
		ImageType:       imgData.Type,
		GroupIdList:     strings.Join([]string{this.groupId}, ","),
		SearchExtParams: extParam,
	}
	rsp := struct {
		BaseResponse
		Result struct {
			FaceToken       string            `json:"face_token"`
			ComparisonInfos []*ComparisonInfo `json:"user_list"`
		} `json:"result"`
	}{}
	_, err = requests.Post(this.cli.ctx, _url,
		opt.BodyJSON(&req), opt.BindJSON(&rsp))
	if err != nil {
		return nil, err
	}
	if err := rsp.intoError(); err != nil {
		return nil, err
	}
	return rsp.Result.ComparisonInfos, nil
}

func (this *_LibraryImpl) MultiSearch(
	imgData *ImageData,
	extParam *MultiSearchExtParams) ([]*MultiSearchResultItem, error) {
	accessToken, err := this.cli.getAccessToken()
	if err != nil {
		return nil, err
	}
	_url, err := murl.MakeUrlString(internal.DuFaceBusinessUrlV3,
		murl.WithPath("multi-search"),
		murl.WithQuery("access_token", accessToken))
	req := struct {
		Image       string     `json:"image"`
		ImageType   ImageTypes `json:"image_type"`
		GroupIdList string     `json:"group_id_list"`
		*MultiSearchExtParams
	}{
		Image:                imgData.Data,
		ImageType:            imgData.Type,
		GroupIdList:          strings.Join([]string{this.groupId}, ","),
		MultiSearchExtParams: extParam,
	}
	rsp := struct {
		BaseResponse
		Result struct {
			FaceNum                int                      `json:"face_num"`
			MultiSearchResultItems []*MultiSearchResultItem `json:"face_list"`
		} `json:"result"`
	}{}
	_, err = requests.Post(this.cli.ctx, _url,
		opt.BodyJSON(&req), opt.BindJSON(&rsp))
	if err != nil {
		return nil, err
	}
	if err := rsp.intoError(); err != nil {
		return nil, err
	}
	return rsp.Result.MultiSearchResultItems, nil
}

func (this *_LibraryImpl) Drop() error {
	return this.groupCtl("delete")
}

func (this *_LibraryImpl) Create() error {
	return this.groupCtl("add")
}

func (this *_LibraryImpl) groupCtl(act string) error {
	accessToken, err := this.cli.getAccessToken()
	if err != nil {
		return err
	}
	_url, err := murl.MakeUrlString(internal.DuFaceBusinessUrlV3,
		murl.WithPath("faceset/group"),
		murl.WithPath(act),
		murl.WithQuery("access_token", accessToken))
	req := struct {
		GroupId string `json:"group_id"`
	}{
		GroupId: this.groupId,
	}
	rsp := BaseResponse{}
	if _, err := requests.Post(this.cli.ctx, _url,
		opt.BodyJSON(&req), opt.BindJSON(&rsp)); err != nil {
		return err
	}
	return rsp.intoError()
}
