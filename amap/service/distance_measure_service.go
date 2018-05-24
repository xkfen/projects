package service

import (
	"amap/api/request"
	"amap/tool/logger"
)

type DistanceMeasureResult struct {
	// 距离 米
	Distance string `json:"distance"`
	// 预计时间  秒
	Duration string `json:"duration"`
}

// 距离测量
func DistanceMeasureService(key, origin, destination string)(distance, duration string , err error){
	resp, respErr := request.DistanceHttpRequest(key, origin, destination)
	if respErr != nil {
		logger.Error("err", "DistanceMeasureService##距离测量接口出错", respErr.Error())
		return
	}
	// todo 以下只是针对resp.Results数组为1的情况，不能统一做成这样，实际需要的时候记得来修改
	for _, result := range resp.Results {
		distance = result.Result.Distance
		duration = result.Result.Duration
	}
	// todo resp.Results是数组，需要根据需求来展示
	return
}