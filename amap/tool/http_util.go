package tool

import (
	"bytes"
	"errors"
	"fmt"
	"amap/tool/logger"
	"io/ioutil"
	"net/http"
	"net/url"
)

// toWho:日志标识第三方名称
func PostJSON(reqUrl string, jsonParams interface{}, toWho string) ([]byte, error) {
	jsonB := StringifyJsonToBytes(jsonParams)
	if len(string(jsonB)) > 500 {
		logger.Info("["+toWho+"]请求参数,部分:"+string(jsonB)[:500], "url", reqUrl)
	} else {
		logger.Info("["+toWho+"]请求参数:"+string(jsonB), "url", reqUrl)
	}
	req, rErr := http.NewRequest("POST", reqUrl, bytes.NewBuffer(jsonB))
	if rErr != nil {
		return nil, rErr
	}
	req.Header.Set("Content-Type", "application/json")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		logger.Error("["+toWho+"]请求不成功:", "err", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		logger.Error("["+toWho+"]读取响应体报错:", "err", err2.Error())
		return nil, err2
	}
	if len(string(b)) < 500 {
		logger.Info("[" + toWho + "]请求结果:" + string(b))
	} else {
		logger.Info("[" + toWho + "]请求结果:" + string(b)[:500])
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("[" + toWho + "]请求不成功:" + string(b))
	}
	return b, nil
}

// method:请求方法
// Content-Type:application/json
// toWho:日志标识第三方名称
func HttpJSON(method string, reqUrl string, jsonParams interface{}, toWho string) ([]byte, error) {
	jsonB := StringifyJsonToBytes(jsonParams)
	logger.Info("["+toWho+"]请求参数:"+string(jsonB), "method", method, "url", reqUrl)
	req, rErr := http.NewRequest(method, reqUrl, bytes.NewBuffer(jsonB))
	if rErr != nil {
		return nil, rErr
	}
	req.Header.Set("Content-Type", "application/json")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		logger.Error("["+toWho+"]请求不成功:", "err", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		logger.Error("["+toWho+"]读取响应体报错:", "err", err2.Error())
		return nil, err2
	}
	logger.Info("[" + toWho + "]请求结果:" + string(b))
	if resp.StatusCode != 200 {
		return nil, errors.New("[" + toWho + "]请求不成功:" + string(b))
	}
	return b, nil
}

// example: http://host:port/uri/?param1=1&param2=2
func Get(reqUrl string, jsonParams interface{},toWho string) ([]byte, error) {
	var params url.Values = url.Values{}
	var jsonObj map[string]interface{}
	jsonB := StringifyJsonToBytes(jsonParams)
	logger.Info("["+toWho+"]请求参数:"+string(jsonB), "url", reqUrl)
	ParseJsonFromBytes(jsonB, &jsonObj)
	for k, v := range jsonObj {
		params.Set(k, fmt.Sprintf("%v", v))
	}
	logger.Debug("["+toWho+"]get请求url:", "url", reqUrl+"?"+params.Encode())
	req, rErr := http.NewRequest("GET", reqUrl+"?"+params.Encode(), nil)
	if rErr != nil {
		return nil, rErr
	}
	//req.Header.Set("sign", sig)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		logger.Error("["+toWho+"]请求不成功:", "err", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		logger.Error("["+toWho+"]读取响应体报错:", "err", err2.Error())
		return nil, err2
	}
	logger.Info("[" + toWho + "]请求结果:" + string(b))
	if resp.StatusCode != 200 {
		return nil, errors.New("[" + toWho + "]请求不成功:" + string(b))
	}
	return b, nil
}

