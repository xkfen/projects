package service

import (
	"testing"
	"amap/tool/logger"
	"amap/config"
	"github.com/stretchr/testify/assert"
	"fmt"
	"amap/tool"
)

func TestWalkPathPlaningService(t *testing.T) {
	logger.InitLogger(logger.LvlDebug, nil)
	origin, originErr := GeoCodingService(config.Key, config.PrivateKey, "广东省深圳市宝安区西乡街道麻布新村1巷6号801")
	assert.NoError(t, originErr)
	destination, destErr := GeoCodingService(config.Key, config.PrivateKey, "广东省深圳市宝安区新安街道众里创业社区403")
	assert.NoError(t, destErr)
	plan, err := WalkPathPlaningService(config.Key, origin, destination)
	assert.NoError(t, err)
	fmt.Println(tool.StringifyJson(plan))
}