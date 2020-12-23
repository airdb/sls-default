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


## FAQ

问题1：输入 serverless 时没有默认弹出中文引导。

解决方案： 在 .env 文件中增加配置 SERVERLESS_PLATFORM_VENDOR=tencent 即可。

问题2：在境外网络环境，输入 sls deploy 后部署十分缓慢。

解决方案：在 .env 文件中增加配置 GLOBAL_ACCELERATOR_NA=true 则开启境外加速 。

问题3：输入 sls deploy 后部署报网络错误。

解决方案：在 .env 文件中增加以下代理配置。