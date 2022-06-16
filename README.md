# cp_go_grpc/v2

go语言grpc项目的组件.必须go 1.18+才可以使用,如果低版本使用请使用v0版本

## 组件包括

| 组件名             | 说明                                                                   |
| ------------------ | ---------------------------------------------------------------------- |
| `sdk`              | grpc的客户端sdk模板                                                    |
| `serv`             | grpc的服务端模板                                                       |
| `main`             | 入口文件                                                               |
| `service_schema`   | 项目定义service的proto文件                                             |
| `sdkinterceptor`   | grpc的客户端拦截器模板                                                 |
| `servinterceptor`  | grpc的服务端拦截器模板                                                 |
| `googleapi_schema` | 项目定义service的proto文件,用于构造grpc-gateway                        |
| `openapi_schema`   | 项目定义google/api的proto文件夹,用于构造grpc-gateway的swagger.json文件 |
| `simple_serv`      | grpc不含gateway的的服务端模板                                          |
