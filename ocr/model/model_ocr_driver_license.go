package model

// 驾驶证识别请求参数
type DriverLicenseReq struct {
	// (必填)图像数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
	Image string `json:"image"`
	// (非必填)是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括:- true：检测朝向；- false：不检测朝向。
	DetectDirection bool `json:"detect_direction"`
}

// 驾驶证识别响应参数
type DriverLicenseResp struct {
	// 驾驶证号
	DriverNumber string `json:"driver_number"`
	// 姓名
	Name string `json:"name"`
	// 有效期限
	ExpirationDate string `json:"expiration_date"`
	// 准驾车型
	DriverType string `json:"driver_type"`
	// 有效起始日期
	EffectiveStartDate string `json:"effective_start_date"`
	// 住址
	Location string `json:"location"`
	// 国籍
	Nationality string `json:"nationality"`
	// 出生日期
	Birthday string `json:"birthday"`
	// 性别
	Sex string `json:"sex"`
	// 初次领证日期
	IssueDate string `json:"issue_date"`


	// 唯一的log id，用于问题定位
	LogId int `json:"log_id"`
	// 识别结果数，表示words_result的元素个数
	WordsResultNum int `json:"words_result_num"`
	// 识别结果数组
	WordsResult []interface{} `json:"words_result"`
	// 识别结果字符串
	Words string `json:"words"`
}