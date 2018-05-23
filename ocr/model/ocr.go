package model

import (
	"gcoresys/common/http"
)

const (
	AccessTokenUrl      = "https://aip.baidubce.com/oauth/2.0/token"
	IdCardUrl           = "https://aip.baidubce.com/rest/2.0/ocr/v1/idcard"
	GeneralWordsUrl     = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic"
	AccurateWordUrl     = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate"
	TableWordsSubmitUrl = "https://aip.baidubce.com/rest/2.0/solution/v1/form_ocr/request"
	TableWordsGetUrl    = "https://aip.baidubce.com/rest/2.0/solution/v1/form_ocr/get_request_result"
	BankCardUrl         = "https://aip.baidubce.com/rest/2.0/ocr/v1/bankcard"
	CustomRecognizeUrl  = "https://aip.baidubce.com/rest/2.0/solution/v1/iocr/recognise"
	DriverLicenseUrl  = "https://aip.baidubce.com/rest/2.0/ocr/v1/driving_license"
)

// 银行卡类型
const (
	BankCardTypeZero = "上传银行卡不能识别"
	BankCardTypeOne = "借记卡"
	BankCardTypeTwo = "信用卡"
	Zero = 0
	One = 1
	Two = 2
)

// 身份证识别
const (
	// 身份证状态
	NormalStatus = "normal"
	ReversedSideStatus = "reversed_side"
	NonIdCardStatus = "non_idcard"
	BlurredStatus = "blurred"
	OverExposureStatus = "over_exposure"
	Normal = "正常识别"
	ReversedSide = "未摆正身份证"
	NonIdCard = "上传的图片中不包含身份证"
	Blurred = "身份证模糊"
	OverExposure = "身份证关键字段反光或过曝"
	Unknown = "未知状态"

	// 身份证类型
	NormalType = "normal"
	CopyType = "copy"
	TemporaryType = "temporary"
	ScreenType = "screen"

	N = "正常身份证"
	Copy = "复印件"
	Temporary = "临时身份证"
	Screen = "翻拍"
	UnknownType = "上传的身份证非法"

)

// 自定义识别
const (
	Name = "姓名"
	Marriage = "婚姻状况"
	CertificateNumber = "证件号码"
	PhoneNumber = "手机号码"
)
// 鉴权请求参数
type AccessTokenReq struct {
	ClientId string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// 鉴权响应参数
type AccessTokenResp struct {
	http.BaseResp
	// token
	AccessToken string `json:"access_token"`
	// 过期时间
	ExpiresIn float64 `json:"expires_in"`
}

// 获取access_token 返回err resp
type AccessTokenErrResp struct {
	Error string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// 身份证识别请求参数
type IdCardRecognizeReq struct {
	// 是否检测图像朝向，默认不检测，即：false(非必填)
	DetectDirection bool `json:"detect_direction"`
	// front：身份证正面；back：身份证背面(必填)
	IdCardSide string `json:"id_card_side"`
	// 图像数据，base64编码后进行urlencode，
	// 要求base64编码和urlencode后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	// 必填
	Image string `json:"image"`
	// 否开启身份证风险类型(身份证复印件、临时身份证、身份证翻拍、修改过的身份证)功能，
	// 默认不开启，即：false。可选值:true-开启；false-不开启
	// 非必填
	DetectRisk bool `json:"detect_risk"`

}

// 身份证识别响应参数
type IdCardRecognizeRespData struct {
	http.BaseResp
	// 姓名
	Name string `json:"name"`
	// 身份证号码
	IdCardNumber string `json:"id_card_number"`
	// 性别
	Sex string `json:"sex"`
	// 民族
	Nation string `json:"nation"`
	// 住址
	Location string `json:"location"`
	// 签发日期
	IssueDate string `json:"issue_date"`
	// 签发机关
	IssueOrganization string `json:"issue_organization"`
	//  失效日期
	ExpirationDate string `json:"expiration_date"`
	// 身份证状态
	Status string `json:"status"`
	// 身份证类型： normal-正常身份证；copy-复印件；temporary-临时身份证；screen-翻拍；unknow-其他未知情况
	RiskType string `json:"risk_type"`
}
//type IdCardRecognizeRespData struct {
//	http.BaseResp
//	// 图像方向，当detect_direction=true时存在,1:未定义;0：正向;1:逆时针90度;2:逆时针180度;3:逆时针270度
//	Direction int `json:"direction"`
//	/**
//	normal-识别正常
//	reversed_side-未摆正身份证
//	non_idcard-上传的图片中不包含身份证
//	blurred-身份证模糊
//	over_exposure-身份证关键字段反光或过曝
//	unknown-未知状态
//	 */
//	ImageStatus string `json:"image_status"`
//	/**
//		输入参数 detect_risk = true 时，则返回该字段识别身份证类型: normal-正常身份证；copy-复印件；temporary-临时身份证；screen-翻拍；unknow-其他未知情况
//	 */
//	RiskType string `json:"risk_type"`
//	/**
//	如果参数 detect_risk = true 时，则返回此字段。如果检测身份证被编辑过，该字段指定编辑软件名称，如:Adobe Photoshop CC 2014 (Macintosh),如果没有被编辑过则返回值无此参数
//	 */
//	EditTool string `json:"edit_tool"`
//	/**
//	唯一的log id，用于问题定位
//	 */
//	LogId uint64 `json:"log_id"`
//	/**
//	定位和识别结果数组
//	 */
//	WordsResult []string `json:"words_result"`
//	/**
//	识别结果数，表示words_result的元素个数
//	 */
//	WordsResultNum uint32 `json:"words_result_num"`
//	/**
//	位置数组（坐标0点为左上角）
//	 */
//	Location string `json:"location"`
//	/**
//	表示定位位置的长方形左上顶点的水平坐标
//	 */
//	Left uint32 `json:"left"`
//	/**
//	表示定位位置的长方形左上顶点的垂直坐标
//	 */
//	Top uint32 `json:"top"`
//	/**
//	表示定位位置的长方形的宽度
//	 */
//	Width uint32 `json:"width"`
//	/**
//	表示定位位置的长方形的高度
//	 */
//	Height uint32 `json:"height"`
//	/**
//	识别结果字符串
//	 */
//	Words  string `json:"words"`
//}

// 银行卡识别请求参数
type BankCardRecognizeReq struct {
	//图像数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	Image string `json:"image"`
}

// 银行可识别响应参数
//type BankCardRecognizeResp struct {
//	// 请求标识码，随机数，唯一。
//	LogId int `json:"log_id"`
//	// 返回结果
//	Result interface{} `json:"result"`
//	// 银行卡卡号
//	BankCardNumber string `json:"bank_card_number"`
//	// 银行名，不能识别时为空
//	BankName string `json:"bank_name"`
//	// 银行卡类型，0:不能识别; 1: 借记卡; 2: 信用卡
//	BankCardType int `json:"bank_card_type"`
//}

type BankCardRecognizeResp struct {
	http.BaseResp
	// 返回结果
	// 银行卡卡号
	BankCardNumber string `json:"bank_card_number"`
	// 银行名，不能识别时为空
	BankName string `json:"bank_name"`
	// 银行卡类型，0:不能识别; 1: 借记卡; 2: 信用卡
	BankCardType string `json:"bank_card_type"`
}

// 通用文字识别(高精度版)请求参数
type GeneralWordsRecognizeReq struct {
	// 图像数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	Image string `json:"image"`
	/**
	是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括:
	- true：检测朝向；
	- false：不检测朝向。
	 */
	DetectDirection bool `json:"detect_direction"`
	// 是否返回识别结果中每一行的置信度
	Probability string `json:"probability"`
}

// 文字识别高精度版响应参数
type GeneralWordsRecognizeResp struct {
	/**
	图像方向，当detect_direction=true时存在。
	- -1:未定义，
	- 0:正向，
	- 1: 逆时针90度，
	- 2:逆时针180度，
	- 3:逆时针270度
	 */
	//Direction int `json:"direction"`
	///**
	//唯一的log id，用于问题定位
	// */
	//LogId int `json:"log_id"`
	//// 识别结果数组
	//WordsResult []string `json:"words_result"`
	//// 识别结果数，表示words_result的元素个数
	//WordsResultNum int `json:"words_result_num"`
	// 识别结果字符串
	Words string `json:"words"`
	// 识别结果中每一行的置信度值，包含average：行置信度平均值，variance：行置信度方差，min：行置信度最小值
	//Posibility interface{} `json:"posibility"`
}

// 文字识别(含位置高精度版)
type WordAccurateLocationRecognizeReq struct {
	// 图像数据(必填)，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	Image string `json:"image"`
	// 是否定位单字符位置，big：不定位单字符位置，默认值；small：定位单字符位置（可选值：big、small）
	RecognizeGranularity string `json:"recognize_granularity"`
	/**
	是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括:
- true：检测朝向；
- false：不检测朝向。
	 */
	DetectDirection bool `json:"detect_direction"`
	// 是否返回文字外接多边形顶点位置，不支持单字位置。默认为false，可选值(true, false)
	VertexesLocation string `json:"vertexes_location"`
	// 是否返回识别结果中每一行的置信度 可选值true.false
	Probability string `json:"probability"`
}

// 文字识别(含位置高精度版)
type WordAccurateLocationRecognizeResp struct {
	
} 

//表格文字识别提交请求参数
type TableWordsSubmitReq struct {
	//图像数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	Image string `json:"image"`
}

// 表格文字识提交别响应参数
type TableWordsSubmitResp struct {
	// 唯一的log id，用于问题定位
	LogId int `json:"log_id"`
	// 返回的结果
	Result []interface{} `json:"result"`
	// 该图片对应请求的request_id
	RequestId string `json:"request_id"`
}

//表格文字识别获取请求参数
type TableWordsGetReq struct {
	// 发送表格文字识别请求时返回的request id
	RequestId string `json:"request_id"`
	// 期望获取结果的类型，取值为“excel”时返回xls文件的地址，取值为“json”时返回json格式的字符串,默认为”excel”
	ResultType string `json:"result_type"`
}
// 表格文字识别获取响应参数
type TableWordsGetResp struct {
	// 唯一的log id，用于问题定位
	LogId int `json:"log_id"`
	// 返回的结果
	Result []interface{} `json:"result"`
	// 识别结果字符串，如果request_type是excel，则返回excel的文件下载地址，如果request_type是json，则返回json格式的字符串
	ResultData string `json:"result_data"`
	// 表格识别进度（百分比）
	Percent int `json:"percent"`
	// 该图片对应请求的request_id
	RequestId string `json:"request_id"`
	// 识别状态，1：任务未开始，2：进行中,3:已完成
	RetCtring string `json:"ret_msg"`
}


// 自定义模板识别请求参数
type CustomRecognizeReq struct {
	// 图像数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	Image string `json:"image"`
	// 您在自定义文字识别平台制作的模版的ID，举例：Nsdax2424asaAS791823112
	TemplateSign string `json:"templateSign"`
}


// 自定义模板响应参数
type CustomRecognizeResp struct {
	//// 0代表成功，如果有错误码返回可以参考下方错误码列表排查问题
	//ErrorCode int `json:"error_code"`
	//// 如果error_code具体的失败信息，可以参考下方错误码列表排查问题
	//ErrorMsg string `json:"error_msg"`
	//// 识别返回的结果
	//Data map[string]interface{} `json:"data"`
	////表示是否结构话成功，true为成功，false为失败；成功时候，返回结构化的识别结果；失败时，如果能识别，返回类似通用文字识别的结果，如果不能识别，返回空
	//IsStructured bool `json:"is_structured"`
	//// 调用的日志id
	//LogId string `json:"log_id"`
	//// isStructured 为 true 时存在，表示字段的名字；如果 isStructured 为 false 时，不存在
	//WordName string `json:"word_name"`
	//// 识别的字符串或单字
	//Word string `json:"word"`
	//// 字符串或单字所在矩形框
	//Rect map[string]interface{} `json:"rect"`
	//Charset []map[string]interface{} `json:"charset"`
	// 姓名
	Name string `json:"name"`
	// 证件类型
	CertificateType string `json:"certificate_type"`
	// 证件号码
	CertificateNumber string `json:"certificate_number"`
	// 性别
	Sex string `json:"sex"`
	// 出生日期
	Birthday string `json:"birthday"`
	// 婚姻状况
	Marriage string `json:"marriage"`
	// 手机号码
	PhoneNumber string `json:"phone_number"`

} 