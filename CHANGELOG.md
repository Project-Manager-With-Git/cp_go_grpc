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

