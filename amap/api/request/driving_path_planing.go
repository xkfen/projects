package request

import (
	"amap/tool"
	"amap/config"
	"amap/tool/logger"
	"errors"
)

// 驾车路径规划请求参数
type DrivingReq struct {
	Key string `json:"key"`
	// 出发点坐标
	Origin string `json:"origin"`
	// 终点坐标
	Destination string `json:"destination"`
}

// 驾车路径规划响应参数
type DrivingResp struct {
	Status string `json:"status"`
	// 返回状态说明：status为0时，info返回错误原因，否则返回“OK”
	Info string `json:"info"`
	// 驾车路径规划方案数目
	Count string `json:"count"`
	// 驾车路径规划信息列表
	Route []DrivingRoute `json:"route"`
}

// 驾车路径规划列表
type DrivingRoute struct {
	// 起点坐标
	Origin string `json:"origin"`
	// 终点坐标
	Destination string `json:"destination"`
	// 打车费用
	TaxiCost string `json:"taxi_cost"`
	// 驾车换乘方案
	Paths []DrivingPath `json:"paths"`
}

type DrivingPath struct {
	// 驾车换乘方案
	Path string `json:"path"`
	// 行驶距离 米
	Distance string `json:"distance"`
	// 预计行驶时间 秒
	Duration string `json:"duration"`
	// 导航策略
	Strategy string `json:"strategy"`
	// 此导航方案道路收费 元
	Tolls string `json:"tolls"`
	// 限行结果：0 代表限行已规避或未限行，即该路线没有限行路段，1 代表限行无法规避，即该线路有限行路段
	Restriction string `json:"restriction"`
	//  红绿灯个数
	TrafficLights string `json:"traffic_lights"`
	// 收费路段距离
	TollDistance string `json:"toll_distance"`
	// 导航路段
	Steps []DrivingStep `json:"steps"`
}

type DrivingStep struct {
	//  行驶指示
	Instruction string `json:"instruction"`
	// 方向
	Orientation string `json:"orientation"`
	// 道路名称
	Road string `json:"road"`
	// 此段路距离 米
	Distance string `json:"distance"`
	// 此段收费 元
	Tolls string `json:"tolls"`
	// 收费路段距离
	TollDistance string `json:"toll_distance"`
	// 主要收费道路
	TollRoad string `json:"toll_road"`
	// 导航主要动作
	Action []DrivingAction `json:"action"`
	// 导航辅助动作
	AssistantAction []DrivingAssistantAction `json:"assistant_action"`
	// 驾车导航详细信息
	Tmcs DrivingTmcs `json:"tmcs"`
}

// 导航主要动作
type DrivingAction struct {
	Action string `json:"action"`
}

// 导航辅助动作
type DrivingAssistantAction struct {
	AssistantAction string `json:"assistant_action"`
}

// 驾车导航详细信息
type DrivingTmcs struct {
	// 此段路的长度
	Distance string `json:"distance"`
	// 此段路的交通情况:未知、畅通、缓行、拥堵
	Status string `json:"status"`
	// 此段路的轨迹
	PolyLine string `json:"polyline"`
}

func DrivingPathPlanHttpRequest(key, origin, destination string )(DrivingResp, error){
	req := DrivingReq{
		Key:key,
		Origin:origin,
		Destination:destination,
	}
	// 发送请求
	byteData, err := tool.Get(config.Driving, req, config.ApiSource)
	if err != nil {
		logger.Error("err", "调用高德地图驾车路径规划接口出错", err.Error())
		return DrivingResp{}, err
	}
	// 解析响应体
	var resp DrivingResp
	err = tool.ParseJsonFromBytes(byteData, &resp)
	if err != nil {
		logger.Error("err", "将驾车路径规划响应参数解析为结构体出错", err.Error())
		return DrivingResp{}, err
	}

	// 错误请求
	if resp.Status != "1" && resp.Info != "OK" {
		logger.Error("调用驾车路径规划响应出错哦")
		return DrivingResp{}, errors.New(tool.GetErrMsgByInfo(resp.Info))
	}

	// 成功请求
	return resp, nil
}