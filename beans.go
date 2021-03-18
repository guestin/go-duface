package duface

import "github.com/guestin/mob/merrors"

type BaseResponse struct {
	ErrCode int    `json:"error_code"`
	ErrMsg  string `json:"error_msg"`
	LogId   int64  `json:"log_id"`
	Ts      int64  `json:"timestamp"`
	Cached  int    `json:"cached"`
}

func (this *BaseResponse) parseError(expect ...int) error {
	fullExpect := append([]int{0}, expect...)
	for _, v := range fullExpect {
		if v == this.ErrCode {
			return nil
		}
	}
	return merrors.Errorf0(this.ErrCode, "[%d] %s", this.ErrCode, this.ErrMsg)
}

type FaceField string

const (
	AGE         FaceField = "age"
	BEAUTY      FaceField = "beauty"
	EXPRESSION  FaceField = "expression"
	FACE_SHAPE  FaceField = "face_shape"
	GENDER      FaceField = "gender"
	GLASSES     FaceField = "glasses"
	LANDMARK    FaceField = "landmark"
	LANDMARK150 FaceField = "landmark150"
	QUALITY     FaceField = "quality"
	EYE_STATUS  FaceField = "eye_status"
	EMOTION     FaceField = "emotion"
	FACE_TYPE   FaceField = "face_type"
	MASK        FaceField = "mask"
	SPOOFING    FaceField = "spoofing" // 合成图检测
)

type ImageTypes string

const (
	BASE64     ImageTypes = "BASE64"
	URL        ImageTypes = "URL"
	FACE_TOKEN ImageTypes = "FACE_TOKEN"
)

type ControlValues string

const (
	NONE    ControlValues = "NONE"
	DEFAULT               = NONE
	LOW     ControlValues = "LOW"
	NORMAL  ControlValues = "NORMAL"
	HIGH    ControlValues = "HIGH"
)

type ActionType string

const (
	APPEND  ActionType = "APPEND"
	REPLACE ActionType = "REPLACE"
)

type RegExtParams struct {
	// 用户资料，长度限制 256B 默认空
	UserInfo        string        `json:"user_info,omitempty"`
	QualityControl  ControlValues `json:"quality_control,omitempty"`
	LivenessControl ControlValues `json:"liveness_control,omitempty"`
	ActionType      ActionType    `json:"action_type,omitempty"`
}

type FaceLocation struct {
	Left     float32 `json:"left"`
	Top      float32 `json:"top"`
	Width    float32 `json:"width"`
	Height   float32 `json:"height"`
	Rotation int     `json:"rotation"`
}

type RegisterFaceResult struct {
	FaceToken string       `json:"face_token"`
	Location  FaceLocation `json:"location"`
}

type BasicResponse struct {
	ErrCode int    `json:"error_code"`
	LogId   int    `json:"log_id"`
	ErrMsg  string `json:"error_msg"`
}

// 通用搜索扩展参数
type SearchExtGeneric struct {
	QualityControl  ControlValues `json:"quality_control,omitempty"`
	LivenessControl ControlValues `json:"liveness_control,omitempty"`
	//当需要对特定用户进行比对时,指定 user_id 进行比对.即人脸认证功能.
	UserId string `json:"user_id,omitempty"`
	// 返回相似度最高的几个用户,默认为1,最多返回50个
	MaxUserNum int `json:"max_face_num,omitempty"`
	// 从指定的 group 中进行查找 用逗号分隔，上限 10 个
	// 由于library本身占用了一个,所以这里最多支持9个
	GroupIdList []string `json:"-"`
}

// 1:N 搜索扩展参数
type SearchExtParams struct {
	SearchExtGeneric
	// 0: 代表检测出的人脸按照人脸面积从大到小排列
	// 1: 代表检测出的人脸按照距离图片中心从近到远排列
	// 默认为 0
	FaceSortType int `json:"face_sort_type"`
}

// M:N 搜索扩展参数
type MultiSearchExtParams struct {
	SearchExtGeneric
	// 匹配阈值（设置阈值后，score 低于此阈值的用户信息将不会返回） 最大 100 最小 0 默认 80
	// 此阈值设置得越高，检索速度将会越快，推荐使用默认阈值 80
	MatchThreshold int `json:"match_threshold,omitempty"`
}

type ImageData struct {
	// 图片信息 (总数据大小应小于 10M),图片上传方式根据 image_type 来判断
	Data string `json:"image"`
	/*
		图片类型
		BASE64: 图片的 base64 值,base64 编码后的图片数据,编码后的图片大小不超过 2M；
		URL: 图片的 URL 地址 (可能由于网络等原因导致下载图片时间过长)；
		FACE_TOKEN: 人脸图片的唯一标识,调用人脸检测接口时,会为每个人脸图片赋予一个唯一的 FACE_TOKEN,
		同一张图片多次检测得到的 FACE_TOKEN 是同一个.
	*/
	Type ImageTypes `json:"image_type"`
}

type ComparisonInfo struct {
	GroupId  string  `json:"group_id"`
	UserId   string  `json:"user_id"`
	UserInfo string  `json:"user_info"`
	Score    float32 `json:"score"`
}

type MultiSearchResultItem struct {
	FaceToken       string            `json:"face_token"`
	FaceLocation    FaceLocation      `json:"location"`
	ComparisonInfos []*ComparisonInfo `json:"user_list"`
}

type DetectExtParams struct {
	/*
			包括 age,beauty,expression,face_shape,gender,glasses,landmark,landmark150,
			quality,eye_status,emotion,face_type,mask,spoofing 信息
			逗号分隔.
		    默认只返回 face_token,人脸框,概率和旋转角度
	*/
	FaceFields FaceField `json:"face_field,omitempty"`
	/*
		最多处理人脸的数目,默认值为 1,根据人脸检测排序类型检测图片中排序第一的人脸（默认为人脸面积最大的人脸）,最大值 120
	*/
	MaxFaceNum uint32 `json:"max_face_num,omitempty"`
	/*
		人脸的类型
		LIVE 表示生活照:通常为手机,相机拍摄的人像图片,或从网络获取的人像图片等
		IDCARD 表示身份证芯片照:二代身份证内置芯片中的人像照片
		WATERMARK 表示带水印证件照:一般为带水印的小图,如公安网小图
		CERT 表示证件照片:如拍摄的身份证,工卡,护照,学生证等证件图片
		默认 LIVE
	*/
	FaceType        string        `json:"face_type,omitempty"`
	LivenessControl ControlValues `json:"liveness_control,omitempty"`
	/*
		人脸检测排序类型
		0: 代表检测出的人脸按照人脸面积从大到小排列
		1: 代表检测出的人脸按照距离图片中心从近到远排列
		默认为 0
	*/
	FaceSortType int `json:"face_sort_type,omitempty"`
}

type FaceAngle struct {
	Yaw   float32 `json:"yaw"`   // 三维旋转之左右旋转角 [-90 (左), 90 (右)]
	Pitch float32 `json:"pitch"` // 三维旋转之俯仰角度 [-90 (上), 90 (下)]
	Roll  float32 `json:"roll"`  // 平面内旋转角 [-180 (逆时针), 180 (顺时针)]
}

type FaceAttribute struct {
	Type        string  `json:"type"`        // 分类
	Probability float32 `json:"probability"` // probability: 范围[0~1],0 最小,1 最大.
}

// [0,1] 取值,越接近 0 闭合的可能性越大
type EyeStatus struct {
	Left  float32 `json:"left_eye"`
	Right float32 `json:"right_eye"`
}

type LandmarkPoint struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type FaceOcclusion struct {
	LeftEye    float32 `json:"left_eye"`    // 左眼
	RightEye   float32 `json:"right_eye"`   // 右眼
	Nose       float32 `json:"nose"`        // 鼻子
	Mouth      float32 `json:"mouth"`       // 嘴巴
	LeftCheck  float32 `json:"left_check"`  // 左脸颊
	RightCheck float32 `json:"right_check"` // 右脸颊
	Chin       float32 `json:"chin"`        // 下巴
}

type FaceQuality struct {
	// 人脸各部分遮挡的概率,范围 [0~1],0 表示完整,1 表示不完整
	Occlusion FaceOcclusion `json:"occlusion"`
	// 人脸模糊程度,范围 [0~1],0 表示清晰,1 表示模糊
	Blur float32
	// 取值范围在 [0~255], 表示脸部区域的光照程度 越大表示光照越好
	Illumination float32
	// 人脸完整度,0 或 1, 0 为人脸溢出图像边界,1 为人脸都在图像边界内
	Completeness int
}

type FaceDetectionItem struct {
	// 人脸图片的唯一标识 （人脸检测 face_token 有效期为 60min）
	Token string `json:"face_token"`
	// 人脸在图片中的位置
	Location FaceLocation `json:"location"`
	// 人脸置信度,范围[0~1],代表这是一张人脸的概率,0 最小,1 最大.其中返回 0 或 1 时,数据类型为 Integer
	Probability float32 `json:"face_probability"`
	// 人脸旋转角度参数
	Angle FaceAngle `json:"angle"`
	// 年龄 ,当 face_field 包含 age 时返回
	Age *float32 `json:"age"`
	// 美丑打分,范围 0-100,越大表示越美.当 face_fields 包含 beauty 时返回
	Beauty *int `json:"beauty"`
	// type: none: 不笑；smile: 微笑；laugh: 大笑
	Expression *FaceAttribute `json:"expression"` // 表情,当 face_field 包含 expression 时返回
	// type: square: 正方形 triangle: 三角形 oval: 椭圆 heart: 心形 round: 圆形
	Shape *FaceAttribute `json:"face_shape"` // 脸型,当 face_field 包含 face_shape 时返回
	// type: male: 男性 female: 女性
	Gender *FaceAttribute `json:"gender"` // 性别,face_field 包含 gender 时返回
	// type: none: 无眼镜,common: 普通眼镜,sun: 墨镜
	Glasses   *FaceAttribute `json:"glasses"`    // 性别,face_field 包含 glasses 时返回
	EyeStatus *EyeStatus     `json:"eye_status"` // 双眼状态（睁开 / 闭合） face_field 包含 eye_status 时返回
	// type: angry: 愤怒 disgust: 厌恶 fear: 恐惧 happy: 高兴 sad: 伤心 surprise: 惊讶 neutral: 无表情 pouty: 撅嘴 grimace: 鬼脸
	Emotion *FaceAttribute `json:"emotion"` // 情绪 face_field 包含 emotion 时返回
	// type: human: 真实人脸 cartoon: 卡通人脸
	Type *FaceAttribute `json:"face_field"` //真实人脸 / 卡通人脸 face_field 包含 face_type 时返回
	// type: 没戴口罩 / 戴口罩 取值 0 或 1 0 代表没戴口罩 1 代表戴口罩
	Mask        *FaceAttribute  `json:"mask"`        // 口罩识别 face_field 包含 mask 时返回
	Landmark    []LandmarkPoint `json:"landmark"`    // 4 个关键点位置,左眼中心,右眼中心,鼻尖,嘴中心.face_field 包含 landmark 时返回
	Landmark72  []LandmarkPoint `json:"landmark72"`  // 72 个特征点位置 face_field 包含 landmark72 时返回
	Landmark150 []LandmarkPoint `json:"landmark150"` // 150 个特征点位置 face_field 包含 landmark150 时返回
	Quality     *FaceQuality    `json:"quality"`     //人脸质量信息.face_field 包含 quality 时返回
	Spoofing    float32         `json:"spoofing"`    // 判断图片是否为合成图,注意:官方文档对该字段的取值范围没有说明
}

type FaceDetectResult struct {
	Num     int                  `json:"face_num"`
	Results []*FaceDetectionItem `json:"face_list"`
}

type UserFaceItem struct {
	FaceToken  string `json:"face_token"`
	CreateTime string `json:"ctime"`
}
