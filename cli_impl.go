package duface

import (
	"context"
	"github.com/guestin/go-duface/internal"
	"github.com/guestin/go-requests"
	"github.com/guestin/go-requests/opt"
	"github.com/guestin/mob/murl"
	"sync"
	"time"
)

type _ClientImpl struct {
	ctx          context.Context
	ApiKey       string
	AppSecretKey string
	locker       sync.Locker
	accessInfo   *internal.TimedAccessToken
}

func NewClient(ctx context.Context, apiKey, appSecret string) Client {
	return &_ClientImpl{
		ctx:          ctx,
		ApiKey:       apiKey,
		AppSecretKey: appSecret,
		locker:       new(sync.Mutex),
	}
}

func (this *_ClientImpl) getAccessToken() (string, error) {
	this.locker.Lock()
	defer this.locker.Unlock()
	if this.accessInfo != nil && !this.accessInfo.IsExpired() {
		return this.accessInfo.AccessToken, nil
	}
	_url, err := murl.MakeUrlString(internal.DuFaceServiceAuthUrl,
		murl.WithPath("token"),
		murl.WithQuery("grant_type", "client_credentials"),
		murl.WithQuery("client_id", this.ApiKey),
		murl.WithQuery("client_secret", this.AppSecretKey))
	if err != nil {
		return "", err
	}
	irsp, err := requests.Get(this.ctx, _url,
		opt.BindJSON(new(internal.Credential)))
	if err != nil {
		return "", err
	}
	r := irsp.(*internal.Credential)
	this.accessInfo = internal.NewAccessTokenInfo(r.AccessToken, r.ExpireIn, time.Hour)
	return this.accessInfo.AccessToken, nil

}

func (this *_ClientImpl) NewLibrary(groupId string) (Library, error) {
	return newLibrary(this, groupId), nil
}

func (this *_ClientImpl) ListLibraries(start, length int) ([]string, error) {
	accessToken, err := this.getAccessToken()
	if err != nil {
		return nil, err
	}
	req := struct {
		Start  int
		Length int
	}{
		Start:  start,
		Length: length,
	}
	_url, err := murl.MakeUrlString(internal.DuFaceBusinessUrlV3,
		murl.WithPath("faceset/group/getlist"),
		murl.WithQuery("access_token", accessToken))
	if err != nil {
		return nil, err
	}
	rsp := struct {
		BaseResponse
		Result struct {
			GroupIds []string `json:"group_id_list"`
		} `json:"result"`
	}{}
	_, err = requests.Post(this.ctx, _url,
		opt.BodyJSON(&req),
		opt.BindJSON(&rsp))
	if err != nil {
		return nil, err
	}
	if err = rsp.parseError(); err != nil {
		return nil, err
	}
	return rsp.Result.GroupIds, nil
}

func (this *_ClientImpl) FaceDetect(imgData *ImageData, params *DetectExtParams) (*FaceDetectResult, error) {
	accessToken, err := this.getAccessToken()
	if err != nil {
		return nil, err
	}
	_url, err := murl.MakeUrlString(internal.DuFaceBusinessUrlV3,
		murl.WithPath("detect"),
		murl.WithQuery("access_token", accessToken))
	req := struct {
		*ImageData
		*DetectExtParams
	}{
		ImageData:       imgData,
		DetectExtParams: params,
	}
	rsp := struct {
		BaseResponse
		Result *FaceDetectResult `json:"result"`
	}{}
	_, err = requests.Post(this.ctx, _url,
		opt.BodyJSON(&req),
		opt.BindJSON(&rsp))
	if err != nil {
		return nil, err
	}
	if err = rsp.parseError(); err != nil {
		return nil, err
	}
	return rsp.Result, nil
}
