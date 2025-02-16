package common

import (
	"github.com/MQEnergy/thirdapi-core/core/ele"
	"github.com/MQEnergy/thirdapi-core/routers/ele/common"
)

type Common struct {
	Base *ele.Base
}

// New 实例化
func New(appId, appSecret string) *Common {
	return &Common{
		Base: ele.New(appId, appSecret),
	}
}

//PICTURE_UPLOAD_URL       = "picture.upload"       // 上传图片
//COMMON_TOKEN_GETUSER_URL = "common.token.getuser" // 通过access_token 获取授权用户及店铺信息
//SHORT_LINK_GENERATE_URL  = "shortLink.generate"   // 饿了么app相关短链生成

// PictureUpload 上传图片
func (a *Common) PictureUpload(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(common.PICTURE_UPLOAD_URL, params).WithClient()
}

// CommonTokenGetUser 通过access_token 获取授权用户及店铺信息
func (a *Common) CommonTokenGetUser(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(common.COMMON_TOKEN_GETUSER_URL, params).WithClient()
}

// ShortLinkGenerate 饿了么app相关短链生成
func (a *Common) ShortLinkGenerate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(common.SHORT_LINK_GENERATE_URL, params).WithClient()
}
