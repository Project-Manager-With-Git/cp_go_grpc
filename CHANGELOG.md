# v0.0.4

## 更新组件

1. grpc版本更新到1.44.0

## 新增功能

1. 增加对grpc-admin的支持

## 实现优化

1. 将服务侧的工具服务注册以及xds服务单独拆出来维护

# v0.0.3

## 细节优化

1. 为要扩展的位置增加注释
2. 明确必填项
3. 更新`github.com/Golang-Tools/loggerhelper`版本
4. 删除无用字段
5. sdk和simplesdk的配置项`app_name`和`app_version`改名为`requester_app_name`和`requester_app_version`以明确含义
6. sdk和simplesdk的配置项`address`改名为`query_addresses`以明确含义

## 更新组件

1. sdk和simplesdk的`NewCtx()`方法增加参数opt可以进行扩展添加元数据
2. sdk和simplesdk的`NewCtx()`方法现在也可以为流请求提供上下文

## 修复bug

修复服务端一处引用bug

# v0.0.2

## 更新组件

+ `sdk`现在`Conn`将有`NewCtx`接口而`SDK`的`NewCtx`将被取消
+ `sdk`增加对xds的支持
+ `serv`使用`Unimplemented<Service>Server struct`
+ `serv`增加对xds的支持

## 新增组件

| 组件名            | 说明                                  |
| ----------------- | ------------------------------------- |
| `simplesdk`       | grpc的去掉拦截器的简化版客户端sdk模板 |
| `simpleserv`      | grpc的去掉拦截器的简化版服务端模板    |
| `sdkinterceptor`  | grpc的客户端拦截器模板                |
| `servinterceptor` | grpc的服务端拦截器模板                |

# v0.0.1

创建了项目

## 组件包括

| 组件名           | 说明                       |
| ---------------- | -------------------------- |
| `sdk`            | grpc的客户端sdk模板        |
| `serv`           | grpc的服务端模板           |
| `main`           | 入口文件                   |
| `service_schema` | 项目定义service的proto文件 |
