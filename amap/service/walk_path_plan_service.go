package service

import (
	"amap/api/request"
	"amap/tool/logger"
	"encoding/json"
	"amap/tool"
)

// 路径方案
type Path struct {
	// 起点和终点的距离
	Distance string `json:"distance"`
	// 预计步行时间
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

// 步行路径规划
func WalkPathPlaningService(key, origin, destination string)(Path, error){
	// 调用接口
	resp, err := request.WalkHttpRequest(key, origin, destination)
	if err != nil {
		logger.Error("err", "WalkPathPlaningService###步行路径规划接口出错", err.Error())
		return Path{}, err
	}
	// 解析请求
	plan, err := parseStruct2Struct(resp)
	if err != nil {
		logger.Error("err", "WalkPathPlaningService###解析结构体出错", err.Error())
		return Path{}, err
	}
	return plan, nil
}


// 解析结构体
func parseStruct2Struct(resp request.WalkResp)(plan Path, err error){
	logger.Debug("-------parseStruct2Struct-------")
	for _, path := range resp.Route.Paths {

		plan.Duration = path.Duration
		plan.Distance = path.Distance
		// 将结构体解析为json
		var steps []Step
		for _, stepInfo := range path.Steps {
			var step Step
			byteData, mErr := json.Marshal(stepInfo)
			if mErr != nil {
				logger.Error("err", "结构体解析为json出错", mErr.Error())
				return
			}
			if pErr := tool.ParseJsonFromBytes(byteData, &step); pErr != nil {
				logger.Error("err", "parseStruct2Struct## 将二进制解析为结构体出错", pErr.Error())
				return
			}
			steps = append(steps, step)
		}
		plan.Steps = steps
	}
	return
}
