# SCF
云函数（Serverless Cloud Function，SCF）是腾讯云为企业和开发者们提供的无服务器执行环境，帮助您在无需购买和管理服务器的情况下运行代码。


## 思路

1. 使用 API Gateway 将 HTTP 请求转发给 SCF
2. SCF 中读取 `cloudevents.APIGatewayProxyRequest`
3. 解析映射 `QueryString(GET), Body(POST)` 进行处理。

## Reference

[安装 Serverless Framework](https://cloud.tencent.com/document/product/1154/42990)


## Deploy scf
```
make
```
