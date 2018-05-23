package service

import (
	//"amap/tool"
	//"amap/tool/logger"
	"amap/api/request"
)

// 公交路径规划
type BusInfo struct {
	//Route BusRouteInfo `json:"route"`
	// 公交换乘方案数目
	Count string `json:"count"`
	// 起点和终点的步行距离 单位：米
	Distance string `json:"distance"`
	// 出租车费用  元
	TaxiCost string `json:"taxi_cost"`
	Infos []BusRouteInfo `json:"infos"`
	//Bus []request.BusNavInfo `json:"bus"`
}
type BusRouteInfo struct {
	// 费用  元
	Cost string `json:"cost"`
	// 夜班车
	NightFlag string `json:"night_flag"`
	// 时间
	Duration string `json:"duration"`
	Distance string `json:"distance"`
	BusLineInfo []request.BusLineInfo `json:"bus_line_info"`
	//Bus []request.BusNavInfo `json:"bus"`
	
}

type TransitLine struct {
	// 公交路线名称
	Name string `json:"name"`
	// 公交路线id
	ID string `json:"id"`
	// 公交类型
	Type string `json:"type"`
	// 公交行驶距离 米
	Distance string `json:"distance"`
	// 预计行驶时间 秒
	Duration string `json:"duration"`
	// 此路段坐标集
	Polyline string `json:"polyline"`
	// 首班车时间 格式如：0600，代表06：00
	StartTime string `json:"start_time"`
	// 末班车时间  格式如：2300，代表23：00
	EndTime string `json:"end_time"`
	// 此段途经公交站数
	ViaNum string `json:"via_num"`
	// 公交乘车站信息
	DepartureStop request.BusDepartureStop `json:"departure_stop"`
	// 下车站信息
	ArrivalStop request.BusArrivalStop `json:"arrival_stop"`
	// 此段途经公交站点列表
	ViaStops []TransitStop `json:"via_stops"`
}

type TransitStop struct {
	Name string `json:"name"`
}
// 公交路径规划service
//func BusTransitPathPlanService(key, origin, destination, city string)(BusInfo, error){
//	resp, err := request.BusTransitPathPlanHttpRequest(key, origin, destination, city)
//	if err != nil {
//		logger.Error("err", "调用高德公交路径规划接口出错", err.Error())
//		return BusInfo{}, err
//	}
//
//	bus = parseStruct2Struct()
//
//}

// 解析结构体
func parseTransit2Bus(resp request.TransitResp)(bus BusInfo, err error){
	//logger.Debug("-------parseStruct2Struct-------")
	//bus.Count = resp.Count
	//var infos []BusRouteInfo
	//for _, route := range resp.Route {
	//	bus.Distance = route.Distance
	//	bus.TaxiCost = route.TaxiCost
	//	for _, transit := range route.Transits {
	//		for _, detail := range transit.Transit {
	//
	//			//routeInfo.Duration = detail.Duration
	//			for _, segment := range detail.Segments {
	//				for _, busSegment := range segment.Bus {
	//					for _, busLineInfo := range busSegment.BusLines {
	//						var busLine BusInfo
	//						// 费用
	//						busLine.Cost = detail.Cost
	//						// 夜班车
	//						busLine.NightFlag = detail.NightFlag
	//
	//						busLine.Duration = busLine.Duration
	//						busLine.Distance = busLine.Distance
	//						busLine.Name
	//					}
	//				}
	//			}
	//
	//			infos = append(infos, routeInfo)
	//		}
	//	}

		//plan.Duration = path.Duration
		//plan.Distance = path.Distance
		//// 将结构体解析为json
		//var steps []Step
		//for _, stepInfo := range path.Steps {
		//	var step Step
		//	byteData, mErr := json.Marshal(stepInfo)
		//	if mErr != nil {
		//		logger.Error("err", "结构体解析为json出错", mErr.Error())
		//		return
		//	}
		//	if pErr := tool.ParseJsonFromBytes(byteData, &step); pErr != nil {
		//		logger.Error("err", "parseStruct2Struct## 将二进制解析为结构体出错", pErr.Error())
		//		return
		//	}
		//	steps = append(steps, step)
		//}
		//plan.Steps = steps
	//}
	return
}