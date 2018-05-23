package request

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

}

// 换乘路段列表
type SegmentInfo struct {
	// 此路段公交导航信息
	Walking string `json:"walking"`
	// 此路段公交导航信息
	Bus string `json:"bus"`
	// 地铁入口
	Entrance string `json:"entrance"`
	// 地铁出口
	Exit string `json:"exit"`
	// 乘坐火车的信息
	Railway string `json:"railway"`
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