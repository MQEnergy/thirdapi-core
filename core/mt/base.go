/*

 */
package mt

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
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
		fmt.Println(formatRequestData(b.RequestData))
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
	//https://waimaiopen.meituan.com/api/v1/order/getOrderDetail?app_id=2742&order_id=4612593754970170004&timestamp=1657238693a016042051a1a78004a37b57dd4cdee5
	//https://waimaiopen.meituan.com/api/v1/order/getOrderDetail?app_id=2742&order_id=4612593754970170004&timestamp=1657238544a016042051a1a78004a37b57dd4cdee5
	// 将request和system数据变成url参数
	b.RequestBody = buildUrlParams(requestData)
	sign := b.GenerateSign(systemData, requestData)
	systemData["sig"] = sign
	b.SystemData = systemData

	b.SystemBody = buildUrlParams(systemData)
	return b
}

// VerifySign 验证回调签名
func (b *Base) VerifySign(data map[string]interface{}) bool {
	if data["sig"] != nil {
		return false
	}
	originSign := data["sig"]
	delete(data, "sig")
	systemData := map[string]interface{}{
		"app_id":    data["app_id"],
		"timestamp": data["timestamp"],
	}
	sign := b.GenerateSign(systemData, data)
	return originSign == sign
}

// GenerateSign 生成sign
func (b *Base) GenerateSign(systemData, requestData map[string]interface{}) string {
	for key, val := range systemData {
		requestData[key] = val
	}
	keys := b.ArraySort(requestData)
	var tmp []string
	for _, val := range keys {
		var params string
		if reflect.TypeOf(requestData[val]).String() == "map[string]interface {}" {
			marshal, _ := json.Marshal(requestData[val])
			params = string(marshal)
		} else {
			params = gconv.String(requestData[val])
		}
		tmp = append(tmp, val+"="+params)
	}
	sortUrl := BaseUrl + b.Url + "?" + strings.Join(tmp, "&") + b.AppSecret
	md5 := md5.New()
	io.WriteString(md5, sortUrl)
	md5Str := fmt.Sprintf("%x", md5.Sum(nil))
	return md5Str
}

// ArraySort 数组排序
func (b *Base) ArraySort(arr map[string]interface{}) []string {
	arr, keys := b.KSort(arr)
	for _, val := range keys {
		if reflect.TypeOf(arr[val]).String() == "map[string]interface {}" {
			b.ArraySort(arr[val].(map[string]interface{}))
		}
	}
	return keys
}

// KSort 按照键排序 map
func (b *Base) KSort(arr map[string]interface{}) (map[string]interface{}, []string) {
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
			temp = gconv.String(val)
		}
		urlValues[key] = []string{temp}
	}
	return urlValues
}

// buildUrlParams 将map数据生成url参数字符串
func buildUrlParams(data map[string]interface{}) string {
	var params []string
	var temp string
	for key, val := range data {
		if reflect.TypeOf(val).String() == "map[string]interface {}" {
			marshal, _ := json.Marshal(val)
			temp = string(marshal)
		} else {
			temp = gconv.String(val)
		}
		params = append(params, key+"="+temp)
	}
	return strings.Join(params, "&")
}
