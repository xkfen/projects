package service

import (
	"testing"
	"amap/tool"
	"amap/config"
	"github.com/stretchr/testify/assert"
	"amap/tool/logger"
)

func TestGeoCodingService1(t *testing.T) {
	logger.InitLogger(logger.LvlDebug, nil)
	if tool.IsDev() {
		location, err := GeoCodingService(config.Key, config.PrivateKey, "广东省深圳市宝安区西乡街道麻布新村1巷6号801")
		assert.NoError(t, err)
		// 113.870939,22.570183
		logger.Info("msg", "location", location)
	}
}

func TestGeoCodingService2(t *testing.T) {
	logger.InitLogger(logger.LvlDebug, nil)
	if tool.IsDev() {
		location, err := GeoCodingService(config.Key, config.PrivateKey, "广东省深圳市宝安区新安街道众里创业社区403")
		assert.NoError(t, err)
		// 113.873779,22.562423
		logger.Info("msg", "location", location)
	}
}