package tool

import (
	"encoding/json"
	"amap/tool/logger"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
	"sort"
	"strings"
)

// 只有为true的时候才调用接口
func IsDev () bool{
	return true
}

/**

func genAccountString(agentID string) string {
	return url.QueryEscape(agentID)
}

func genSecrest(agentID, password string) string {
	hash := sha256.New()
	hash.Write([]byte(url.QueryEscape(agentID+password)))
	md := hash.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(md))
}

func genSign(result string) string {
	hash := sha256.New()
	hash.Write([]byte(result))
	md := hash.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(md))
}

// 生成sign
func GenSign(agentID , password string, params map[string]string) (result string) {
	accountString := genAccountString(agentID)
	logger.Info("加密", "accountString", accountString)

	secrest := genSecrest(agentID, password)
	logger.Info("加密", "secrest", secrest)

	var keys []string
	for k, v := range params {
		keys = append(keys, k)
		logger.Debug("v", "v", v)
	}
	sort.Strings(keys)
	logger.Info("加密", "sortKeys", keys)
	for _, k := range keys {
		result += k + params[k]
	}

	result = accountString + result + secrest
	logger.Info("加密", "result_before", result)

	result = genSign(result)
	logger.Info("加密", "result_final", result)
	return
}
 */

 /**
 sign 签名规则:

 sig=MD5(请求参数键值对（按参数名的升序排序），加（请注意“加”字无需输入）私钥)；

例如：

请求服务为“testservice”；

请求参数分别为“a=23，b=12，d=48，f=8，c=67”；

私钥为“bbbbb”

则数字签名为：sig=md5(a=23&b=12&c=67&d=48&f=8bbbbb)

注意：

生成签名的内容，（上文提到的拼装的参数，也就是md5()中的内容），必须为utf-8编码格式。
在计算md5的参数如果出现＋号，请正常计算sig，但在请求的时候，需要用urlencode进行编码再请求。
请求参数排序需要注意，如果参数名的第一个字母顺序相同，就比较第二个字母。以此类推，直至得到排序结果。
  */
// 生成数字签名
func GenSign(privateKey string, params map[string]string)(sig string){
	// keys 用来接收请求参数
	var keys []string
	for k, _ := range params {
		keys = append(keys, k)
	}
	// 按参数名的升序排序
	sort.Strings(keys)
	logger.Info("加密", "sortKeys", keys)
	// 请求参数拼接成字符串
	// params.Encode()

	var result string
	for _, k := range keys {
		result += k + "=" + params[k] + "&"
	}
	result = strings.TrimRight(result, "&")
	// 拼接请求参数与私钥
	result = result + privateKey
	logger.Info("加密", "加密之前:", result)
	// md5加密
	sig = EncryptByMD5(result)
	return
}

// md5加密
func EncryptByMD5(data string) string {
	m := md5.New()
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum(nil))
}

// unix timestampStr to datetimeStr
func UnixTimestampStrToTime(data string) string {
	d, _ := strconv.Atoi(data)
	return time.Unix(int64(d), 0).Format("2006-01-02 15:04:05")
}

// 字符串转json对象
func ParseJson(str string, result interface{}) error {
	//return json.Unmarshal([]byte(str), &result)
	return json.Unmarshal([]byte(str), result)
}

func ParseJsonFromBytes(b []byte, result interface{}) error {
	//return json.Unmarshal(b, &result)
	return json.Unmarshal(b, result)
}

// json对象转字符串
func StringifyJson(obj interface{}) string {
	//b, err := json.Marshal(obj)
	b, err := json.Marshal(obj)
	if err != nil {
		logger.Info("err","转换json字符串出错", err.Error())

		return ""
	}
	return string(b)
}

func StringifyJsonToBytes(obj interface{}) []byte {
	//b, err := json.Marshal(obj)
	b, err := json.Marshal(obj)
	if err != nil {
		logger.Info("err","转换json字符串出错", err.Error())
		return nil
	}
	return b
}