package internal

import "time"

type TimedAccessToken struct {
	AccessToken  string
	ExpireAt     time.Time
	AheadTimeout time.Duration // 提前超时时间
}

func NewAccessTokenInfo(token string, expireIn int64, aheadTimeout time.Duration) *TimedAccessToken {
	nowTs := time.Now().Unix()
	return &TimedAccessToken{
		AccessToken:  token,
		ExpireAt:     time.Unix(nowTs+expireIn, 0),
		AheadTimeout: aheadTimeout,
	}
}

func (this *TimedAccessToken) IsExpired() bool {
	return this.ExpireAt.Before(time.Now().Add(this.AheadTimeout))
}
