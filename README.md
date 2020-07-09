# SCF
云函数（Serverless Cloud Function，SCF）是腾讯云为企业和开发者们提供的无服务器执行环境，帮助您在无需购买和管理服务器的情况下运行代码。

## Reference

[SCF Golang 开发指南](https://cloud.tencent.com/document/product/583/18032)

## 思路

1. 使用 API Gateway 将 HTTP 请求转发给 SCF
2. SCF 中读取 `cloudevents.APIGatewayProxyRequest`
3. 解析映射 `QueryString(GET), Body(POST)` 进行处理。


## 上传及调用
好像不支持自动化，需要在页面操作。上传，发布。

```buildoutcfg
curl https://service-xxxx.apigw.tencentcs.com/release/<function_name>
```


## Deploy by github action
```
scf configure set --appid $APPID --region $REGION --secret-id $SECRET_ID --secret-key $SECRET_KEY
scf  init  -r go1
scf deploy -t template.yaml  -f
```
