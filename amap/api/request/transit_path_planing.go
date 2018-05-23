package request

import (
	"amap/config"
	"amap/tool"
	"amap/tool/logger"
	"errors"
)

// 公交路径规划请求参数
type TransitReq struct {
	// 请求服务权限标识:用户在高德地图官网申请Web服务API类型KEY （Y）
	Key string `json:"key"`
	//  出发点（Y）
	Origin string `json:"origin"`
	// 目的地（Y）
	Destination string `json:"destination"`
	// 城市/跨城规划时的起点城市（Y）
	City string `json:"city"`
	// 跨城公交规划时的终点城市 (跨城必填)
	Cityd string `json:"cityd"`
	// 返回结果详略:可选值：base(default)/all base:返回基本信息(默认值)；all：返回全部信息
	Extensions string `json:"extensions"`
	// 公交换乘策略
	/**
	可选值：
		0：最快捷模式(默认)

		1：最经济模式

		2：最少换乘模式

		3：最少步行模式

		5：不乘地铁模式
	 */
	Strategy string `json:"strategy"`
	// 是否计算夜班车：0：不计算夜班车（默认） 1：计算夜班车
	NightFlag string `json:"nightflag"`
	// 出发日期:根据出发时间和日期筛选可乘坐的公交路线，格式：date=2014-3-19
	Date string `json:"date"`
	// 出发时间:根据出发时间和日期筛选可乘坐的公交路线，格式：time=22:34
	Time string `json:"time"`
}

// 公交路径规划响应参数
type TransitResp struct {
	// 1：成功；0：失败
	Status string `json:"status"`
	// 返回的状态信息
	Info string `json:"info"`
	// 公交换乘方案数目
	Count string `json:"count"`
	// 公交换乘信息列表
	Route []TransitRoute `json:"route"`

}

//  公交换乘信息列表
type TransitRoute struct {
	// 起点坐标
	Origin string `json:"origin"`
	// 终点坐标
	Destination string `json:"destination"`
	// 起点和终点的步行距离 单位：米
	Distance string `json:"distance"`
	// 出租车费用 元
	TaxiCost string `json:"taxi_cost"`
	// 公交换乘方案列表
	Transits []TransitInfo `json:"transits"`
}

type TransitInfo struct {
	Transit []TransitDetail `json:"transit"`
}

type TransitDetail struct {
	// 此换乘方案价格 单位：元
	Cost string `json:"cost"`
	// 此换乘方案预期时间 单位：秒
	Duration string `json:"duration"`
	// 是否是夜班车：0：非夜班车；1：夜班车
	NightFlag string `json:"nightflag"`
	// 此方案总步行距离 单位：米
	WalkingDistance string `json:"walking_distance"`
	// 换乘路段列表
	Segments []SegmentInfo `json:"segments"`
}

// 换乘路段列表
type SegmentInfo struct {
	// 此路段公交导航信息
	Walking []WalkingNavInfo `json:"walking"`
	// 此路段公交导航信息
	Bus []BusNavInfo `json:"bus"`
	// 地铁入口
	Entrance []SubwayInfo `json:"entrance"`
	// 地铁出口
	Exit []SubwayInfo `json:"exit"`
	// 乘坐火车的信息
	Railway []RailwayInfo `json:"railway"`
}

// 步行导航信息列表
type WalkingNavInfo struct {
	// 起点坐标
	Origin string `json:"origin"`
	// 终点坐标
	Destination string `json:"destination"`
	// 每段线路步行距离 单位：米
	Distance string `json:"distance"`
	// 步行预计时间 单位：秒
	Duration string `json:"duration"`
	// 步行路段列表
	Steps []Step `json:"steps"`
}

// 公交导航信息列表
type BusNavInfo struct {
	// 步行路段列表
	BusLines []BusLineInfo `json:"buslines"`
}

// 步行路段列表
type BusLineInfo struct {
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
	DepartureStop BusDepartureStop `json:"departure_stop"`
	// 下车站信息
	ArrivalStop BusArrivalStop `json:"arrival_stop"`
	// 此段途经公交站点列表
	ViaStops []BusViaStopInfo `json:"via_stops"`
}

/// 公交乘车站信息
type BusDepartureStop struct {
	// 站点名字
	Name string `json:"name"`
	// 站点id
	ID string `json:"id"`
	// 站点经纬度
	Location string `json:"location"`
}

// 下车站信息
type BusArrivalStop struct {
	// 站点名字
	Name string `json:"name"`
	// 站点id
	ID string `json:"id"`
	// 站点经纬度
	Location string `json:"location"`
}

// 此段途经公交站点列表
type BusViaStopInfo struct {
	// 途径公交站点信息
	Name string `json:"name"`
	// 公交站点编号
	ID string `json:"id"`
	// 公交站点经纬度
	Location string `json:"location"`
}

// 地铁出入口信息
type SubwayInfo struct {
	Location string `json:"location"`
}

// 火车信息
type RailwayInfo struct {
	// 线路id编号
	ID string `json:"id"`
	// 该线路车段耗时
	Time string `json:"time"`
	// 线路名称
	Name string `json:"name"`
	// 线路车次号
	Trip string `json:"trip"`
	// 该item换乘段的行车总距离
	Distance string `json:"distance"`
	// 线路车次类型
	Type string `json:"type"`
	// 火车始发站信息
	DepartureStop DepartureStopInfo `json:"departure_stop"`
	// 火车到站信息
	ArrivalStop ArrivalStopInfo `json:"arrival_stop"`
	// 途径站点 extensions=all时返回
	ViaStop ViaStopInfo `json:"via_stop"`
	// 聚合的备选方案，extensions=all时返回
	Alters AltersInfo `json:"alters"`
	// 仓位及价格信息
	Spaces SpacesInfo `json:"spaces"`

}

// 火车站始发信息
type DepartureStopInfo struct {
	// 上车站点ID
	ID string `json:"id"`
	// 上车站点名称
	Name string `json:"name"`
	// 上车站点经纬度
	Location string `json:"location"`
	// 上车站点所在城市的adcode
	AdCode string `json:"adcode"`
	// 上车点发车时间
	Time string `json:"time"`
	// 是否始发站，1表示为始发站，0表示非始发站
	Start string `json:"start"`
}

// 火车到站信息
type ArrivalStopInfo struct {
	// 下车站点ID
	ID string `json:"id"`
	// 下车站点名称
	Name string `json:"name"`
	// 下车站点经纬度
	Location string `json:"location"`
	// 下车站点所在城市的adcode
	AdCode string `json:"adcode"`
	// 到站时间，如大于24:00，则表示跨天
	Time string `json:"time"`
	// 是否为终点站，1表示为终点站，0表示非终点站
	End string `json:"end"`
}

// 途径站点 extensions=all时返回
type ViaStopInfo struct {
	// 途径站点的名称
	Name string `json:"name"`
	// 途径站点的ID
	ID string `json:"id"`
	// 途径站点的坐标点
	Location string `json:"location"`
	// 途径站点的进站时间，如大于24:00,则表示跨天
	Time string `json:"time"`
	// 途径站点的停靠时间，单位：分钟
	Wait string `json:"wait"`
}

// 聚合的备选方案
type AltersInfo struct {
	// 备选方案ID
	ID string `json:"id"`
	// 备选线路名称
	Name string `json:"name"`
}

// 仓位及价格信息
type SpacesInfo struct {
	// 仓位编码
	Code string `json:"code"`
	// 仓位费用
	Cost string `json:"cost"`
}

// 公交导航路径规划
func BusTransitPathPlanHttpRequest(key, origin, destination, city string )(TransitResp, error){
	req := TransitReq {
		Key:key,
		Origin:origin,
		Destination:destination,
		City:city,
	}
	byteData, err := tool.Get(config.Transit, req, config.ApiSource)
	if err != nil {
		logger.Error("err", "调用高德地图公交路径规划出错", err.Error())
		return TransitResp{}, err
	}
	var resp TransitResp
	err = tool.ParseJsonFromBytes(byteData, &resp)
	if err != nil {
		logger.Error("err", "将公交路径规划返回的参数解析为结构体出错", err.Error())
		return TransitResp{}, err
	}

	if resp.Status != "1" && resp.Info != "OK" {
		return TransitResp{}, errors.New(tool.GetErrMsgByInfo(resp.Info))
	}
	return resp,  nil
}