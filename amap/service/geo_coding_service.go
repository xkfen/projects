package service

import (
	"amap/api/request"
	"amap/tool/logger"
)
func GeoCodingService(key, privateKey, address string)(string, error){
	resp, err := request.GenCodingHttpRequest(key, privateKey, address)
	if err != nil {
		logger.Error("err", "GeoCodingService###调用第三方接口查询地理编码出错", err.Error())
		return "", err
	}
	if resp.Count == "1" {
		return  resp.GeoCodes[0].Location, nil
	}
	return "", nil
}
