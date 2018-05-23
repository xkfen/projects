package request

import (
	"amap/tool"
	"amap/config"
	"errors"
	"encoding/json"
	"amap/tool/logger"
)

/**
	该api用于地理编码
 */

 // 地理编码请求参数
 type GenCodingReq struct {
 	// 应用的key(必填)
 	Key string `json:"key"`
 	// 结构化地址信息（必填）
 	Address string `json:"address"`
 	//// 指定查询的城市
 	//City string `json:"city"`
 	//// 批量查询控制
 	//Batch bool `json:"batch"`
 	//// 数字签名
 	//Sig string `json:"sig"`
 	//// 返回数据格式类型
 	//Output string `json:"output"`
 	//// 回调函数
 	//Callback string `json:"callback"`
 }
 // 地理编码响应参数
type GenCodingResp struct {
	// 返回结果状态值:0 表示请求失败；1 表示请求成功。
	Status string `json:"status"`
	// 返回结果数目:结果个数
	Count string `json:"count"`
	// 返回状态说明当 status 为 0 时，info 会返回具体错误原因，否则返回“OK”
	Info string `json:"info"`
	// 地理编码信息列表
	GeoCodes []GeoCodes `json:"geocodes"`
}

// todo 注意，以下结构体的某些字段是不确定的，比如Street,Number，有的传入的地址是完全合法的的格式，但是有的又没有街道等
type GeoCodes struct {
	// 结构化地址信息:省份＋城市＋区县＋城镇＋乡村＋街道＋门牌号码
	FormattedAddress string `json:"formatted_address"`
	// 地址所在的省份
	Province string `json:"province"`
	// 地址所在的城市名
	City string `json:"city"`
	// 城市编码
	//CityCode string `json:"city_code"`
	//// 地址所在的区
	//District string `json:"district"`
	//// 地址所在的乡镇
	//Township []TownShip `json:"township"`
	//// 街道
	//Street []Street `json:"street"`
	//// 门牌
	//Number []Number `json:"number"`
	//// 区域编码
	//Adcode string `json:"adcode"`
	// 坐标点
	Location string `json:"location"`
	// 匹配级别
	//Level string `json:"level"`
}

// 有的有street，有的没有，所有street是不确定的类型
type Street struct {
	Street string `json:"street"`
}
type Number struct {
	Number string `json:"number"`
}

type TownShip struct {

}

func GenCodingHttpRequest(key , privateKey, address string)( GenCodingResp, error){
	req := GenCodingReq {
		Key:key,
		Address:address,
	}
	//params := map[string]string {
	//	"key": key,
	//	"address": address,
	//}
	// 生成数字签名
	//sig := tool.GenSign(privateKey, params)
	//logger.Info("msg", "sig", sig)
	// 调用接口
	//byteData, err := tool.Get(config.GeoCoding, req, sig, config.ApiSource)
	byteData, err := tool.Get(config.GeoCoding, req, config.ApiSource)
	if err != nil {
		return GenCodingResp{}, errors.New("请求第三方接口失败["+err.Error()+"]")
	}

	var resp GenCodingResp
	err = json.Unmarshal(byteData, &resp)
	if err != nil {
		return GenCodingResp{}, errors.New("将第三方响应结果解析为结构体出错:[" + err.Error() + "]")
	}

	logger.Info("msg", "响应体", tool.StringifyJson(resp))
	// status = 1表示成功
	if resp.Status != "1" && resp.Info != "OK"{
		return GenCodingResp{}, errors.New(tool.GetErrMsgByInfo(resp.Info))
	}
	return resp, nil
}