# thirdapi-core
对接美团、饿了么、京东健康、百度健康、京东药急送的服务商或品牌商等三方平台的接口

# 在项目中安装使用
```shell
go get -u github.com/MQEnergy/thirdapi-core
```

# 测试
## 1、安装依赖
```shell
go mod tidy
```

## 运行example
配置appId，appSecret
```shell
go run examples/mt/order/order.go
```
查看效果

## 2、单元测试
```shell
go test
```