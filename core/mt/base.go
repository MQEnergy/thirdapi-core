package mt

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"time"
)

const (
	BaseUrl = "https://waimaiopen.meituan.com/api/v1/" // 请求根地址
	Timeout = 3000                                     // 超时时间 3000ms
)

type Base struct {
	AppId       string                 // app id
	AppSecret   string                 // app secret
	Url         string                 // 请求url
	RequestData map[string]interface{} // 请求参数map
	RequestBody string                 // 请求参数：排序的字符串连接 例如：app_poi_code=&medicine_data=
	SystemData  map[string]interface{} // 系统参数map
	SystemBody  string                 // 系统参数：排序的字符串连接 放在get参数中 例如：app_id=&timestamp=&sign=
	req         *http.Response
	err         error
}

// New 初始化
func New(appId, appSecret string) *Base {
	return &Base{
		AppId:     appId,
		AppSecret: appSecret,
	}
}

// WithClient 请求
func (b *Base) WithClient(method string) (string, error) {
	client := &http.Client{
		Timeout: time.Millisecond * Timeout, // Set 3000ms timeout.
	}
	sortUrl := BaseUrl + b.Url + "?" + b.SystemBody
	if method == "POST" {
		b.req, b.err = client.PostForm(sortUrl, formatRequestData(b.RequestData))
	} else {
		sortUrl += "&" + b.RequestBody
		b.req, b.err = client.Get(sortUrl)
	}
	if b.err != nil {
		return "", b.err
	}
	defer b.req.Body.Close()
	result, err := ioutil.ReadAll(b.req.Body)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// WithRequestParams
// get请求参数设置 所有参数 包括系统参数也通过get方式去请求 例如：app_id=&timestamp=&sig=&{请求参数}
// post请求参数设置
func (b *Base) WithRequestParams(url string, requestData map[string]interface{}) *Base {
	b.Url = url
	b.RequestData = requestData
	systemData := map[string]interface{}{
		"app_id":    b.AppId,
		"timestamp": time.Now().Unix(),
	}
	// 将request和system数据变成url参数
	b.RequestBody = buildUrlParams(requestData)
	sign := b.generateSign(systemData, requestData)
	systemData["sig"] = sign
	b.SystemData = systemData

	b.SystemBody = buildUrlParams(systemData)
	return b
}

// VerifySign 验证回调签名
func (b *Base) VerifySign(params map[string]interface{}) bool {
	value, ok := params["sig"]
	if !ok {
		return false
	}
	originSign := cast.ToString(value)
	delete(params, "sig")
	systemData := map[string]interface{}{
		"app_id":    params["app_id"],
		"timestamp": params["timestamp"],
	}
	sign := b.generateSign(systemData, params)
	return originSign == sign
}

// generateSign 生成sign
func (b *Base) generateSign(systemData, requestData map[string]interface{}) string {
	for key, val := range systemData {
		requestData[key] = val
	}
	_, keys := b.kSort(requestData)
	var tmp []string
	for _, val := range keys {
		var params string
		if reflect.TypeOf(requestData[val]).String() == "map[string]interface {}" {
			marshal, _ := json.Marshal(requestData[val])
			params = string(marshal)
		} else {
			params = cast.ToString(requestData[val])
		}
		tmp = append(tmp, val+"="+params)
	}
	sortUrl := BaseUrl + b.Url + "?" + strings.Join(tmp, "&") + b.AppSecret
	md5 := md5.New()
	io.WriteString(md5, sortUrl)
	md5Str := fmt.Sprintf("%x", md5.Sum(nil))
	return md5Str
}

// kSort 按照键排序 map
func (b *Base) kSort(arr map[string]interface{}) (map[string]interface{}, []string) {
	var keys []string
	for k := range arr {
		keys = append(keys, k)
	}
	var newArr = make(map[string]interface{})
	sort.Strings(keys)
	for _, val := range keys {
		newArr[val] = arr[val]
	}
	return newArr, keys
}

// formatRequestData 格式化请求参数
func formatRequestData(requestData map[string]interface{}) map[string][]string {
	var (
		urlValues = make(map[string][]string)
		temp      string
	)
	for key, val := range requestData {
		if reflect.TypeOf(val).String() == "map[string]interface {}" {
			marshal, _ := json.Marshal(val)
			temp = string(marshal)
		} else {
			temp = cast.ToString(val)
		}
		urlValues[key] = []string{temp}
	}
	return urlValues
}

// buildUrlParams 将map数据生成url参数字符串
func buildUrlParams(data map[string]interface{}) string {
	var params []string
	for key, val := range data {
		params = append(params, key+"="+cast.ToString(val))
	}
	return strings.Join(params, "&")
}
