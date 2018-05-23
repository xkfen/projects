package serviceV1

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// 测试身份证识别
func TestDriverLicenseRecognize(t *testing.T) {
	oneStepTest(func() {
		token := "24.25e7d699e133765b2339e64ac3c4387f.2592000.1523446525.282335-10916026"
		image := getImage("/home/qydev/下载/testImg/word.jpg")
		_ , err := DriverLicenseRecognize(token,  image, true)
		assert.NoError(t, err)
	})
}