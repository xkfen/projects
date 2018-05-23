package serviceV1

import (
	"testing"
	"gcoresys/common/logger"
	"encoding/base64"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
	"gcoresys/common/util"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"ocr/db/config"
)


func TestRun(t *testing.T) {
	logger.InitLogger(logger.LvlDebug, nil)
	config.GetOcrDbConfig("test")
	config.GetDb().LogMode(false)
	suite.Run(t, new(testingSuite))
}

type testingSuite struct {
	suite.Suite
}

// 单步测试
func oneStepTest(testMethods ...func()) {
	logger.InitLogger(logger.LvlDebug, nil)
	config.GetOcrDbConfig("test")
	config.GetDb().LogMode(false)

	for _, testMethods := range testMethods {
		config.ClearAllData()
		testMethods()
		config.ClearAllData()
	}
}

func (s *testingSuite) SetupTest() {
	config.ClearAllData()
}

func (s *testingSuite) TearDownTest() {
	config.ClearAllData()
}

func TestFetchAccessToken(t *testing.T) {
	oneStepTest(func() {
		clientId := "COrlv51qk5gQrWlNOwQXwUqh"
		clientSecret := "GQwlvQqXUX5t67UTRK9W5j8RGNOnZA24"
		accessTokenFetchUrl := "https://aip.baidubce.com/oauth/2.0/token"
		if accessToken , expiresIn, err := FetchAccessToken(clientId, clientSecret, accessTokenFetchUrl); err != nil {

		}else {
			logger.Info("msg", "access_token", accessToken)
			logger.Info("msg", "expires_in", expiresIn)
		}

	})
}

func TestGetOcrAccessToken(t *testing.T) {
	oneStepTest(func() {
		clientId := "EhtC2muicVwZdWfxUGABQLoA"
		clientSecret := "jScfnVpvFjgsb74M6X20O9I9zELA1KZI"
		_, err := GetOcrAccessToken(clientId, clientSecret)
		assert.NoError(t, err)
	})
}

func TestImageUrlEncode(t *testing.T){
	oneStepTest(func() {
		input, err := ioutil.ReadFile("/home/qydev/下载/c5b94c230a2f10d99b11ec64517b894c.jpg")
		if err !=  nil {
			logger.Error("err", "读取本地文件出错", err.Error())
		}
		base64Result := base64.StdEncoding.EncodeToString(input)
		logger.Info("info", "base64编码后的结果", base64Result)
	})

}

func getImage(imagePath string) (string){
	input, err := ioutil.ReadFile(imagePath)
	if err !=  nil {
		logger.Error("err", "读取本地文件出错", err.Error())
	}
	return base64.StdEncoding.EncodeToString(input)
}

// 测试身份证识别
func TestIDCardRecognize(t *testing.T) {
	oneStepTest(func() {
		token := "24.25e7d699e133765b2339e64ac3c4387f.2592000.1523446525.282335-10916026"
		idCardSide := "front"
		image := getImage("/home/qydev/下载/testImg/word.jpg")
		_ , err := IDCardRecognize(token, idCardSide, image, true, true)
		assert.NoError(t, err)
	})
}

// 银行卡识别
func TestBankCardRecognize(t *testing.T) {
	oneStepTest(func() {
		token := "24.25e7d699e133765b2339e64ac3c4387f.2592000.1523446525.282335-10916026"
		// 测试上传的不是正确的银行卡图片
		//image := getImage("/home/qydev/下载/testImg/word.jpg")
		image := getImage("/home/qydev/下载/testImg/bank2.jpg")
		_, err := BankCardRecognize(token, image)
		assert.NoError(t, err)
	})
}

// 测试文字识别高精度版
func TestGeneralWordsRecognize(t *testing.T) {
	oneStepTest(func() {
		token := "24.25e7d699e133765b2339e64ac3c4387f.2592000.1523446525.282335-10916026"
		image := getImage("/home/qydev/下载/testImg/bidcard.jpg")
		_,  err := GeneralWordsRecognize(token, image, "true", true)
		assert.NoError(t, err)
	})
}

// 测试文字识别(含位置高精度版)
func TestWordAccurateLocationRecognize(t *testing.T) {
	oneStepTest(func() {
		token := "24.25e7d699e133765b2339e64ac3c4387f.2592000.1523446525.282335-10916026"
		image := getImage("/home/qydev/下载/testImg/zx3.jpg")
		//token, image, recognizeGranularity, vertexesLocation, probability string, detectDirection bool
		recognizeGranularity := "small"
		vertexesLocation := "false"
		probability := "false"
		detectDirection := false
		result, err := WordAccurateLocationRecognize(token, image, recognizeGranularity, vertexesLocation, probability, detectDirection)
		assert.NoError(t, err)
		fmt.Println(result)
	})
}

// 测试表格文字识别提交接口
func TestTableWordsRecognizeSubmit(t *testing.T) {
	oneStepTest(func() {
		token := "24.25e7d699e133765b2339e64ac3c4387f.2592000.1523446525.282335-10916026"
		image := getImage("/home/qydev/下载/testImg/table2.jpg")
		_,  err := TableWordsRecognizeSubmit(token, image )
		assert.NoError(t, err)
	})
}

// 测试表格文字识别获取接口
func TestTableWordsRecognizeGet(t *testing.T) {
	oneStepTest(func() {
		token := "24.25e7d699e133765b2339e64ac3c4387f.2592000.1523446525.282335-10916026"
		requestId := "10916026_200309"
		resultType := "json"

		result,  err := TableWordsRecognizeGetResult(requestId, resultType, token )
		if err != nil {
			logger.Error("err", "扫描件识别获取出错", err.Error())
		}
		assert.NotEqual(t, nil, result)
		logger.Info(result)
	})
}

type JsonModel struct {
	Name string `json:"name"`
	Urls string `json:"urls"`
	Info Info `json:"info"`
}
type Info struct {
	QQ string `json:"qq"`
	Weixin string `json:"weixin"`
}
func TestReadJsonFile(t *testing.T){
	oneStepTest(func() {
		jsonFile, err  := ioutil.ReadFile("err.json")
		if err != nil {
			logger.Error("err", "读取json文件出错", err.Error())
		}
		jsonMap := make(map[string]interface{})
		err = json.Unmarshal(jsonFile, &jsonMap)
		if err != nil {
			logger.Error("err", "将json文件的数据解析为json格式出错", err.Error())

		}
		name := util.GetStrFromJson(jsonMap, "name")

		logger.Info(name)
	})
}

// 测试自定义模板识别
func TestCustomRecognize(t *testing.T) {
	oneStepTest(func() {
		token := "24.c51a05b99e6d333ded22aea1086d0b32.2592000.1523675583.282335-10929854"
		image := getImage("/home/qydev/下载/testImg/zx5.jpg")
		templateSign := "f83689ca47cf5568560f1006d74f4e89"
		result, err := Custom(token, image, templateSign)
		assert.NoError(t, err)
		fmt.Println(result)
	})
}

