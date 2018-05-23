package request

import (
	"amap/tool"
	"amap/config"
	"amap/tool/logger"
	"encoding/json"
	"errors"
)
/**
   步行路径规划
 */

 // 步行路径规划请求参数
 type WalkReq struct {
 	// 请求服务权限标识
 	Key string `json:"key"`
 	// 出发点
 	Origin string `json:"origin"`
 	// 目的地
 	Destination string `json:"destination"`
 }
 
 // 步行路径规划响应参数
 type WalkResp struct {
 	// 返回状态:1：成功；0：失败
 	Status string `json:"status"`
 	// 返回的状态信息
 	Info string `json:"info"`
 	// 返回结果总数目
 	Count string `json:"count"`
 	// 路线信息列表
 	Route Route `json:"route"`
 }

 // 路线信息
 type Route struct {
 	// 起点坐标
	Origin string `json:"origin"`
	// 终点坐标
	Destination string `json:"destination"`
	// 步行方案
	Paths []Path `json:"paths"`
	
 }
 
 // 路径方案
 type Path struct {
 	// 起点和终点的距离 单位：米
 	Distance string `json:"distance"`
 	// 预计步行时间(单位：秒)
 	Duration string `json:"duration"`
 	// 步行结果列表
	Steps []Step `json:"steps"`
 }

 // 步行结果
 type Step struct {
	//Step StepPlan `json:"step"`
	 // 路段步行指示
	 Instruction string `json:"instruction"`
	 // 道路名称(有可能为string，也有可能是array)
	 //Road string `json:"road"`
	 // 此路段距离(米)
	 Distance string `json:"distance"`
	 // 此路段预计步行时间
	 Duration string `json:"duration"`
	 // 方向(有可能为string，也有可能是array)
	 //Orientation string `json:"orientation"`
	 // 此路段坐标点
	 //PolyLine string `json:"polyline"`
	 // 步行主要动作
	 Action string `json:"action"`
	 // 步行辅助动作(是数组)
	 //AssistantAction []AssistantAction `json:"assistant_action"`
 }

 type AssistantAction struct {
	 AssistantAction string `json:"assistant_action"`
 }
 // 步行方案
 type StepPlan struct {
 	// 路段步行指示
 	Instruction string `json:"instruction"`
 	// 道路名称
 	Road string `json:"road"`
 	// 此路段距离(米)
 	Distance string `json:"distance"`
 	// 此路段预计步行时间
 	Duration string `json:"duration"`
 	// 方向
 	Orientation string `json:"orientation"`
 	// 此路段坐标点
 	PolyLine string `json:"polyline"`
 	// 步行主要动作
 	Action string `json:"action"`
 	// 步行辅助动作
 	AssistantAction string `json:"assistant_action"`
 }

 // 步行路径规划
 func WalkHttpRequest(key, origin, destination string)(WalkResp, error){
 	req := WalkReq{
 		Key:key,
 		Origin:origin,
 		Destination:destination,
	}
	// 调用接口
 	byteData, err := tool.Get(config.Walking, req, config.ApiSource)
 	if err != nil {
 		logger.Error("err", "调用[高德]步行路径规划出错", err.Error())
 		return WalkResp{}, err
	}
	// 解析响应参数
	var resp WalkResp
	err = json.Unmarshal(byteData, &resp)
	if err != nil {
		logger.Error("err", "高德地图步行路径规划响应参数解析为结构体出错", err.Error())
		return WalkResp{}, err
	}

	// 判断接口成功与否
	if resp.Status != "1" && resp.Info != "OK" {
		return WalkResp{}, errors.New(tool.GetErrMsgByInfo(resp.Info))
	}

	return resp, nil
 }