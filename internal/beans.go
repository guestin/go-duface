package internal

/*
  access_token: 要获取的 Access Token；
  expires_in: Access Token 的有效期 (秒为单位,一般为 1 个月)；
  其他参数忽略,暂时不用 ;
*/
type Credential struct {
	RefreshToken  string `json:"refresh_token"`
	ExpireIn      int64  `json:"expires_in"`
	Scope         string `json:"scope"`
	SessionKey    string `json:"session_key"`
	AccessToken   string `json:"access_token"`
	SessionSecret string `json:"session_secret"`
}
