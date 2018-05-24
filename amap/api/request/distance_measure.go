package request

import (
	"amap/config"
	"amap/tool/logger"
	"amap/tool"
	"errors"
)
// 距离测量api

// 距离测量请求参数
type DistanceMeasureReq struct {
	Key string `json:"key"`
	// 出发点
	Origin string `json:"origin"`
	// 目的地
	Destination string `json:"destination"`
}

// 距离测量响应参数
type DistanceMeasureResp struct {
	// 值为0或1，0表示请求失败；1表示请求成功
	Status string `json:"status"`
	// 返回状态说明，status为0时，info返回错误原；否则返回“OK”
	Info string `json:"info"`
	// 距离信息列表
	Results []DistanceMeasureResult `json:"results"`
}

type DistanceMeasureResult struct {
	Result ResultInfo `json:"result"`
}

type ResultInfo struct {
	// 起点坐标，起点坐标序列号（从１开始）
	OriginId string `json:"origin_id"`
	// 终点坐标，终点坐标序列号（从１开始） 
	DestId string `json:"dest_id"`
	// 路径距离，单位：米
	Distance string `json:"distance"`
	// 预计时间
	Duration string `json:"duration"`
	// 仅在出错的时候显示该字段。大部分显示“未知错误”:由于此接口支持批量请求，建议不论批量与否用此字段判断请求是否成功
	Info string `json:"info"`
	/**
	仅在出错的时候显示此字段。
	在驾车模式下：
	1，指定地点之间没有可以行车的道路
	2，起点/终点 距离所有道路均距离过远（例如在海洋/矿业）
	3，起点/终点不在中国境内
	 */
	 Code string `json:"code"`
}

// 距离测量  (todo 默认为驾车模式,需要传入type来指定)
func DistanceHttpRequest(key, origin, destination string)(DistanceMeasureResp, error){
	req := DistanceMeasureReq {
		Key:key,
		Origin:origin,
		Destination:destination,
	}
	// 路径测量接口
	byteData, err := tool.Get(config.DistanceMeasure, req, config.ApiSource)
	if err != nil {
		logger.Error("err", "调用高德路径测量接口出错", err.Error())
		return DistanceMeasureResp{}, err
	}

	// 解析路径测量响应参数
	var resp DistanceMeasureResp
	if err = tool.ParseJsonFromBytes(byteData, &resp); err != nil {
		logger.Error("err", "解析高德路径规划响应参数为结构体出错", err.Error())
		return DistanceMeasureResp{}, err
	}

	// 请求不成功
	if resp.Status != "1" && resp.Info != "OK" {
		logger.Error("路径测量接口请求失败")
		return DistanceMeasureResp{}, errors.New(tool.GetErrMsgByInfo(resp.Info))
	}

	//  请求成功
	return resp, nil
}