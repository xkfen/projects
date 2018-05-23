package request

import (
	"amap/tool"
	"amap/tool/logger"
	"amap/config"
	"errors"
)

// 骑行路径规划请求参数
type BicyclingReq struct {
	// 在高德官网申请的key
	Key string `json:"key"`
	// 起点坐标
	Origin string `json:"origin"`
	// 终点坐标
	Destination string `json:"destination"`
}

// 骑行路径规划响应参数
type BicyclingResp struct {
	// 返回结果,0:成功
	ErrCode int `json:"errcode"`
	// 具体错误原因
	ErrDetail string `json:"errdetail"`
	// 返回状态说明:OK 成功
	ErrMsg string `json:"errmsg"`
	Data BicycleData `json:"data"`
}

type BicycleData struct {
	// 起点坐标
	Origin string `json:"origin"`
	// 终点坐标
	Destination string `json:"destination"`
	Paths []BicyclePath `json:"paths"`
}

type BicyclePath struct {
	// 起终点的骑行距离
	Distance int `json:"distance"`
	// 起终点的骑行时间
	Duration int `json:"duration"`
	// 具体骑行结果
	Steps []BicycleStep `json:"steps"`
}

type BicycleStep struct {
	// 路段骑行指示
	Instruction string `json:"instruction"`
	// 此段路道路名称
	Road string `json:"road"`
	// 此段路骑行距离
	Distance int `json:"distance"`
	// 此段路骑行方向
	Orientation string `json:"orientation"`
	// 此段路骑行耗时 秒
	Duration string `json:"duration"`
	// 此段路骑行主要动作
	Action string `json:"action"`
	// 此段路骑行辅助动作
	AssistantAction string `json:"assistant_action"`
}

// 骑行路径规划
func BicyclingPathPlanHttpRequest(key, origin, destination string)(BicyclingResp, error){
	req := BicyclingReq{
		Key:key,
		Origin:origin,
		Destination:destination,
	}

	// 调用接口
	byteData, err := tool.Get(config.Bicycling, req, config.ApiSource)
	if err != nil {
		logger.Error("err", "高德地图骑行路径规划接口出错哦", err.Error())
		return BicyclingResp{}, err
	}

	// 解析响应
	var resp BicyclingResp
	err = tool.ParseJsonFromBytes(byteData, &resp)
	if err != nil {
		logger.Error("err", "解析骑行路径规划响应为结构体出错哦",err.Error())
		return BicyclingResp{}, err
	}

	// 错误请求
	if resp.ErrCode != 0 && resp.ErrMsg != "OK" {
		logger.Error("骑行路径规划接口出错了")
		return BicyclingResp{}, errors.New(tool.GetErrMsgByInfo(resp.ErrMsg))
	}

	// 成功请求
	return resp, nil
}