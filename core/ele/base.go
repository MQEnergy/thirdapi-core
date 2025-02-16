package ele

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"time"
)

const (
	BaseUrl = "https://api-be.ele.me/" // 请求根地址
	Timeout = 8000                     // 超时时间 3000ms
	Version = 3                        // api版本
	Encrypt = "des.v1"                 // 加密方式
)

type Base struct {
	Source      string
	Secret      string
	AccessToken string
	RequestData map[string]interface{} // 请求参数map

}

func New(source, secret string) *Base {
	return &Base{
		Source: source,
		Secret: secret,
	}
}

// WithClient 请求
func (b *Base) WithClient() (string, error) {
	client := &http.Client{
		Timeout: time.Millisecond * Timeout, // Set 3000ms timeout.
	}
	req, err := client.PostForm(BaseUrl, formatRequestData(b.RequestData))
	if err != nil {
		return "", err
	}
	defer req.Body.Close()
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	return string(result), nil
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

func (b *Base) WithRequestParams(cmd string, requestData map[string]interface{}) *Base {
	body, _ := json.Marshal(requestData)
	systemData := map[string]interface{}{
		"cmd":       cmd,
		"encrypt":   Encrypt,
		"secret":    b.Secret,
		"source":    b.Source,
		"ticket":    b.generateTicket(),
		"timestamp": time.Now().Unix(),
		"version":   Version,
		"body":      string(body),
	}
	systemData["sign"] = b.generateSign(systemData)
	b.RequestData = systemData
	return b
}

// generateSign 生成sign
func (b *Base) generateSign(systemData map[string]interface{}) string {
	_, keys := b.kSort(systemData)
	var tmp = make([]string, 0)
	for _, val := range keys {
		tmp = append(tmp, fmt.Sprintf("%s=%v", val, systemData[val]))
	}
	strSign := strings.Join(tmp, "&")
	md5 := md5.New()
	io.WriteString(md5, strSign)
	md5Str := fmt.Sprintf("%x", md5.Sum(nil))
	md5Str = strings.ToUpper(md5Str)
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

func (b *Base) generateTicket() string {
	var uuid [16]byte
	_, err := rand.Read(uuid[:])
	if err != nil {
		return ""
	}
	str := fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
	return strings.ToUpper(str)
}
