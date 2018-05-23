package serviceV1

import (
	"strings"
	"net/http"
	"gcoresys/common/logger"
	"io/ioutil"
	"bytes"
	"ocr/model"
	"encoding/json"
	"gcoresys/common/util"
	rcUtil "griskcontrol/util"
	"errors"
	"github.com/tidwall/gjson"
	"fmt"
)



func FetchAccessToken(clientId, clientSecret, accessTokenFetchUrl string) (string, float64, error) {
	// 得到access_token请求
	requestLine := strings.Join([]string{
		accessTokenFetchUrl,
		"?grant_type=client_credentials&client_id=",
		clientId,
		"&client_secret=",
		clientSecret}, "")

	resp, err := http.Get(requestLine)
	// 判断响应码
	if err != nil || resp.StatusCode != http.StatusOK {
		logger.Error("err", "发送GET请求获取百度AI失败", err.Error())
		return "", 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("err", "使用GET请求获取access_token body出错", err.Error())
		return "", 0, err
	}

	if bytes.Contains(body, []byte("access_token")) {
		atr := model.AccessTokenResp{}
		err = json.Unmarshal(body, &atr)
		if err != nil {
			logger.Error("err", "使用GET请求获取access_token 返回数据json解析出错", err.Error())
			return "", 0, err
		}
		return atr.AccessToken, atr.ExpiresIn, nil
	} else {
		errAtr := model.AccessTokenErrResp{}
		err = json.Unmarshal(body, &errAtr)
		logger.Error("err", "发送GET请求获取access_token 返回错误", err.Error())
		if err != nil {
			return "", 0, err
		}
		return "", 0, err
	}
}

func GetOcrAccessToken(clientId, clientSecret string) (res *model.AccessTokenResp, err error) {
	//values := url.Values{}
	//values.Set("grant_type", "client_credentials") //规定，无需修改
	//values.Add("client_id", clientId) //根据自己申请填写
	//values.Add("client_secret", clientSecret) //根据自己申请填写
	//resp, err := http.PostForm(accessTokenUrl, values)
	//
	//if err != nil {
	//	return "", err
	//}
	//
	//data, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return "", err
	//}
	//
	//var mapJson map[string]interface{}
	//err = json.Unmarshal(data, &mapJson)
	//if err != nil {
	//	return "", err
	//}

	//accessToken := util.GetStrFromJson(mapJson, "access_token")
	req := &model.AccessTokenReq{
		ClientSecret: clientSecret,
		ClientId:     clientId,
	}
	body, err := rcUtil.PostForm(model.AccessTokenUrl+"?grant_type=client_credentials&client_id="+clientId+"&client_secret="+clientSecret, req, "", false)
	if err != nil {
		return nil, errors.New("请求第三方接口[获取access_token]失败[" + err.Error() + "]")
	}

	resp := make(map[string]interface{})
	if err = json.Unmarshal(body, &resp); err != nil {
		logger.Error("IDCardRecognize#解析json报错", "err", err.Error(), "body", body)
		return nil, err
	}

	logger.Info("----", "resp", resp)
	// 判断是否有错误，有错误就将错误信息返回前端
	errCode := util.GetIntFromJson(resp, "error_code")
	if errCode != 0 {
		errMsg := util.GetStrFromJson(resp, "error_msg")
		if errMsg != "" {
			return nil, errors.New(errMsg)
		}
	}

	accessToken := util.GetStrFromJson(resp, "access_token")
	logger.Info("accessToken", "-----", accessToken)
	expiresIn := util.GetFloatFromJson(resp, "expires_in")
	res = &model.AccessTokenResp{
		BaseResp:    *util.GetBaseResp(nil, "access_token获取成功"),
		AccessToken: accessToken,
		ExpiresIn:   expiresIn,
	}
	return
}

// 检查调用第三方接口有没有出错,出错直接返回
func checkErr(resp map[string]interface{}) (err error) {
	// 判断是否有错误，有错误就将错误信息返回前端
	errCode := util.GetIntFromJson(resp, "error_code")
	if errCode != 0 {
		errMsg := util.GetStrFromJson(resp, "error_msg")
		if errMsg != "" {
			return errors.New(errMsg)
		}
	}
	jsonDataByte, err := ioutil.ReadFile("1.json")
	if err != nil {
		return err
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(jsonDataByte, &jsonMap)
	//string:=strconv.Itoa(int)
	if err != nil {
		return err
	}
	jsonMatchArray := util.GetArrFromJson(jsonMap, "matches")
	for _, array := range jsonMatchArray {
		dataJson := util.JsonToMap(array)
		errorCode := util.GetIntFromJson(dataJson, "error_code")
		errMsg := util.GetStrFromJson(dataJson, "error_msg")
		if errCode == errorCode {
			return errors.New(errMsg)
		}
	}

	return
}

// 身份证识别
func IDCardRecognize(token, idCardSide, image string, detectDirection, detectRick bool) (res *model.IdCardRecognizeRespData, err error) {
	req := &model.IdCardRecognizeReq{
		DetectDirection: detectDirection,
		DetectRisk:      detectRick,
		IdCardSide:      idCardSide,
		Image:           image,
	}
	body, err := rcUtil.PostForm(model.IdCardUrl+"?access_token="+token, req, "", true)
	if err != nil {
		return nil, errors.New("请求第三方接口[身份证识别]失败[" + err.Error() + "]")
	}
	// 将得到的json解析到map里面，方便后面拿数据
	resp := make(map[string]interface{})
	if err = json.Unmarshal(body, &resp); err != nil {
		logger.Error("IDCardRecognize#解析json报错", "err", err.Error(), "body", body)
		return nil, err
	}

	logger.Info("----", "resp", resp)
	// 判断是否有错误，有错误就将错误信息返回前端
	errCode := util.GetIntFromJson(resp, "error_code")
	if errCode != 0 {
		errMsg := util.GetStrFromJson(resp, "error_msg")
		if errMsg != "" {
			return nil, errors.New(errMsg)
		}
	}

	logger.Info("errorCode", "-----", errCode)
	//checkErr(resp)

	// 转为系统需要的数据格式
	res, err = getIdCardResp(resp)
	return
}

// 得到身份证状态
func getIdCardStatus(status string) (idCardStatus string) {
	switch status {
	case model.NormalStatus:
		idCardStatus = model.Normal
	case model.ReversedSideStatus:
		idCardStatus = model.ReversedSide
	case model.NonIdCardStatus:
		idCardStatus = model.NonIdCard
	case model.BlurredStatus:
		idCardStatus = model.Blurred
	case model.OverExposureStatus:
		idCardStatus = model.OverExposure
	default:
		idCardStatus = model.Unknown
	}
	return
}

// 得到身份证类型
func getIdCardType(riskType string) (idCardType string) {
	switch riskType {
	case model.NormalType:
		idCardType = model.N
	case model.CopyType:
		idCardType = model.Copy
	case model.TemporaryType:
		idCardType = model.Temporary
	case model.ScreenType:
		idCardType = model.Screen
	default:
		idCardType = model.UnknownType
	}
	return
}

// 根据调用第三方接口返回的Json解析为系统需要的数据模型
func getIdCardResp(resp map[string]interface{}) (response *model.IdCardRecognizeRespData, err error) {
	// 转为系统需要的数据格式
	wordsResult := util.GetJsonFromJson(resp, "words_result")
	// 首先判断身份证类型，如果类型不对，后面的逻辑就不再继续
	// 身份证类型
	riskType := util.GetStrFromJson(resp, "risk_type")
	if riskType != model.NormalType {
		return nil, errors.New(model.UnknownType)
	}
	idCardType := getIdCardType(riskType)
	// 姓名
	nameJson := util.GetJsonFromJson(wordsResult, "姓名")
	name := util.GetStrFromJson(nameJson, "words")
	// 身份证号码
	idCardNumJson := util.GetJsonFromJson(wordsResult, "公民身份号码")
	idCardNumber := util.GetStrFromJson(idCardNumJson, "words")
	// 性别
	sexJson := util.GetJsonFromJson(wordsResult, "性别")
	sex := util.GetStrFromJson(sexJson, "words")
	// 民族
	nationJson := util.GetJsonFromJson(wordsResult, "民族")
	nation := util.GetStrFromJson(nationJson, "words")
	// 住址
	locationJson := util.GetJsonFromJson(wordsResult, "住址")
	location := util.GetStrFromJson(locationJson, "words")
	// 签发日期json
	issueDateJson := util.GetJsonFromJson(wordsResult, "签发日期")
	// 签发日期
	issueDate := util.GetStrFromJson(issueDateJson, "words")
	// 签发机关json
	issueOrgJson := util.GetJsonFromJson(wordsResult, "签发机关")
	// 签发机关
	issueOrg := util.GetStrFromJson(issueOrgJson, "words")
	// 失效日期json
	expirationDateJson := util.GetJsonFromJson(wordsResult, "失效日期")
	// 失效日期
	expirationDate := util.GetStrFromJson(expirationDateJson, "words")
	// 身份证状态
	status := util.GetStrFromJson(resp, "image_status")
	idCardStatus := getIdCardStatus(status)

	// 得到想要的数据结构
	response = &model.IdCardRecognizeRespData{
		BaseResp:          *util.GetBaseResp(nil, "身份证信息获取成功"),
		Name:              name,
		IdCardNumber:      idCardNumber,
		Sex:               sex,
		Nation:            nation,
		Location:          location,
		IssueDate:         issueDate,
		IssueOrganization: issueOrg,
		ExpirationDate:    expirationDate,
		Status:            idCardStatus,
		RiskType:          idCardType,
	}
	return
}

// 银行卡识别
func BankCardRecognize(token, image string) (res *model.BankCardRecognizeResp, err error) {
	req := &model.BankCardRecognizeReq{
		Image: image,
	}
	body, err := rcUtil.PostForm(model.BankCardUrl+"?access_token="+token, req, "", false)
	if err != nil {
		return nil, errors.New("请求第三方接口[银行卡识别]失败[" + err.Error() + "]")
	}

	// 将调用第三方接口返回的Json解析为map
	resp := make(map[string]interface{})
	if err = json.Unmarshal(body, &resp); err != nil {
		logger.Error("BankCardRecognize#解析json报错", "err", err.Error(), "body", body)
		return nil, err
	}

	// 判断是否有错误，有错误就将错误信息返回前端
	//checkErr(resp)
	//errCode := util.GetIntFromJson(resp, "error_code")
	//if errCode != 0 {
	// errMsg := util.GetStrFromJson(resp, "error_msg")
	// if errMsg != "" {
	//	 return nil, errors.New(errMsg)
	// }
	//}
	// 将响应参数转为自己需要的数据
	res, err = getBankCardResp(resp)
	return
}

func getBankCardResp(resp map[string]interface{}) (response *model.BankCardRecognizeResp, err error) {
	resultJson := util.GetJsonFromJson(resp, "result")
	bankCardType := util.GetIntFromJson(resultJson, "bank_card_type")
	logger.Info("上传的银行卡类型", "---", bankCardType)
	if bankCardType == model.Zero {
		return nil, errors.New(model.BankCardTypeZero)
	}
	bankType := getBankCardType(bankCardType)
	bankName := util.GetStrFromJson(resultJson, "bank_name")
	bankCardNumber := util.GetStrFromJson(resultJson, "bank_card_number")

	response = &model.BankCardRecognizeResp{
		BaseResp:       *util.GetBaseResp(nil, "身份证信息获取成功"),
		BankCardNumber: bankCardNumber,
		BankName:       bankName,
		BankCardType:   bankType,
	}
	return
}

// 得到当前银行卡类型
func getBankCardType(bankCardType int) (bankType string) {
	switch bankCardType {
	case model.One:
		bankType = model.BankCardTypeOne
	case model.Two:
		bankType = model.BankCardTypeTwo
	default:
		bankType = model.BankCardTypeZero
	}
	return
}

// 文字识别(高精度版)
func GeneralWordsRecognize(token, image, probability string, detectDirection bool) (result string, err error) {
	req := &model.GeneralWordsRecognizeReq{
		Image:           image,
		DetectDirection: detectDirection,
		Probability:     probability,
	}

	body, err := rcUtil.PostForm(model.GeneralWordsUrl+"?access_token="+token, req, "", false)
	if err != nil {
		return "", errors.New("请求第三方接口[文字识别(高精度版)]失败[" + err.Error() + "]")
	}

	resp := make(map[string]interface{})
	if err = json.Unmarshal(body, &resp); err != nil {
		logger.Error("GeneralWordsRecognize#解析json报错", "err", err.Error(), "body", body)
		return "", err
	}

	// 判断是否有错误，有错误就将错误信息返回前端
	checkErr(resp)
	//errCode := util.GetIntFromJson(resp, "error_code")
	//if errCode != 0 {
	//	errMsg := util.GetStrFromJson(resp, "error_msg")
	//	if errMsg != "" {
	//		return "", errors.New(errMsg)
	//	}
	//}
	wordsResultJsonArray := util.GetArrFromJson(resp, "words_result")
	for _, wordsResultJson := range wordsResultJsonArray {
		wordsMap := util.JsonToMap(wordsResultJson)
		words := util.GetStrFromJson(wordsMap, "words")
		result += words

	}

	logger.Info("msg", "result", result)
	return
}

// 文字识别(含位置高精度版)
func WordAccurateLocationRecognize(token, image, recognizeGranularity, vertexesLocation, probability string, detectDirection bool) (string, error) {
	req := &model.WordAccurateLocationRecognizeReq{
		Image:                image,
		RecognizeGranularity: recognizeGranularity,
		VertexesLocation:     vertexesLocation,
		Probability:          probability,
		DetectDirection:      detectDirection,
	}

	body, err := rcUtil.PostForm(model.AccurateWordUrl+"?access_token="+token, req, "", false)
	if err != nil {
		return "", errors.New("调用第三方接口[文字识别(含位置高精度版)]出错,[" + err.Error() + "]")
	}

	return string(body), err
}

// 表格文字识别提交接口
func TableWordsRecognizeSubmit(token, image string) (string, error) {
	req := &model.TableWordsSubmitReq{
		Image: image,
	}
	body, err := rcUtil.PostForm(model.TableWordsSubmitUrl+"?access_token="+token, req, "", false)
	if err != nil {
		return "", errors.New("请求第三方接口[表格文字识别提交]失败[" + err.Error() + "]")
	}

	resp := make(map[string]interface{})
	if err = json.Unmarshal(body, &resp); err != nil {
		logger.Error("TableWordsRecognizeSubmit#解析json报错", "err", err.Error(), "body", body)
		return "", err
	}

	// 判断是否有错误，有错误就将错误信息返回前端
	errCode := util.GetIntFromJson(resp, "error_code")
	if errCode != 0 {
		errMsg := util.GetStrFromJson(resp, "error_msg")
		if errMsg != "" {
			return "", errors.New(errMsg)
		}
	}
	return string(body), err
}

// 表格文字识别获取接口
func TableWordsRecognizeGetResult(requestId, resultType, token string) (string, error) {
	req := &model.TableWordsGetReq{
		RequestId:  requestId,
		ResultType: resultType,
	}
	body, err := rcUtil.PostForm(model.TableWordsGetUrl+"?access_token="+token, req, "", false)
	if err != nil {
		return "", errors.New("请求第三方接口[表格文字识别获取]失败[" + err.Error() + "]")
	}

	resp := make(map[string]interface{})
	if err = json.Unmarshal(body, &resp); err != nil {
		logger.Error("TableWordsRecognizeGetResult#解析json报错", "err", err.Error(), "body", body)
		return "", err
	}
	//todo
	// 将响应参数转为自己需要的数据

	// 判断是否有错误，有错误就将错误信息返回前端
	errCode := util.GetIntFromJson(resp, "error_code")
	if errCode != 0 {
		errMsg := util.GetStrFromJson(resp, "error_msg")
		if errMsg != "" {
			return "", errors.New(errMsg)
		}
	}
	return string(body), err
}

// 自定义模板识别
func CustomRecognize(token, image, templateSign string) (response *model.CustomRecognizeResp, err error) {
	req := &model.CustomRecognizeReq{
		Image:        image,
		TemplateSign: templateSign,
	}
	//  调用第三方接口
	body, err := rcUtil.PostForm(model.CustomRecognizeUrl+"?access_token="+token, req, "", false)
	if err != nil {
		return nil, errors.New("请求第三方接口[自定义模板识别获取]失败[" + err.Error() + "]")
	}

	resStr := string(body)
	// 使用第三方库gjson解析
	response = getCustomRecognizeResp(resStr)
	return
}

func Custom(token, image, templateSign string) (string, error) {
	req := &model.CustomRecognizeReq{
		Image:        image,
		TemplateSign: templateSign,
	}
	//  调用第三方接口
	body, err := rcUtil.PostForm(model.CustomRecognizeUrl+"?access_token="+token, req, "", false)
	if err != nil {
		return "", errors.New("请求第三方接口[自定义模板识别获取]失败[" + err.Error() + "]")
	}

	//resStr := string(body)
	// 使用第三方库gjson解析
	//response = getCustomRecognizeResp(resStr)
	logger.Info("msg", "-----", string(body))
	fmt.Println(string(body))
	return string(body), err
}

// 自定义模板识别
func getCustomRecognizeResp(respStr string) (response *model.CustomRecognizeResp) {
	name := gjson.Get(respStr, `data.ret.#[word_name="name"].word`)
	idCard := gjson.Get(respStr, `data.ret.#[word_name="certificateNumber"].word`)
	phoneNumber := gjson.Get(respStr, `data.ret.#[word_name="phoneNumber"].word`)
	birthday := gjson.Get(respStr, `data.ret.#[word_name="birthday"].word`)
	certificateType := gjson.Get(respStr, `data.ret.#[word_name="certificateType"].word`)
	sex := gjson.Get(respStr, `data.ret.#[word_name="sex"].word`)
	marriage := gjson.Get(respStr, `data.ret.#[word_name="marriage"].word`)
	response = &model.CustomRecognizeResp{
		Name:              name.Str,
		PhoneNumber:       phoneNumber.Str,
		CertificateType:   certificateType.Str,
		CertificateNumber: idCard.Str,
		Birthday:          birthday.Str,
		Sex:               sex.Str,
		Marriage:          marriage.Str,
	}
	return
}
