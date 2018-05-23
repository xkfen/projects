package serviceV1

import (
	"ocr/model"
	rcUtil "griskcontrol/util"
	"errors"
	"encoding/json"
	"gcoresys/common/logger"
)

func DriverLicenseRecognize(token ,image string , detectDirection bool)(string , error){
	req := &model.DriverLicenseReq{
		Image:image,
		DetectDirection:detectDirection,
	}
	body , err := rcUtil.PostForm(model.DriverLicenseUrl + "?access_token=" + token, req, "", false)

	if err != nil {
		return "", errors.New("调用第三方接口[驾驶证识别]出错:["+err.Error()+"]")
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return "", errors.New("DriverLicenseRecognize调用身份证识别返回结果解析json出错:["+err.Error()+"]")
	}

	logger.Info("解析后的模型", "resp", resp)

	//errMsg := util.GetStrFromJson(resp, "")
	return string(body), err
}