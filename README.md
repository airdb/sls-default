# SCF
云函数（Serverless Cloud Function，SCF）是腾讯云为企业和开发者们提供的无服务器执行环境，帮助您在无需购买和管理服务器的情况下运行代码。

## Install serverless
```
sudo su -c 'curl -sL https://deb.nodesource.com/setup_15.x | bash -'
sudo apt update
sudo apt install -y nodejs

npm install -g serverless
```

## Deploy

Ensure `.env` exists and correct.
```
serverless deploy
sls deploy
```

