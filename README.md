# shortlink-go

## GoDotEnv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建 .env 文件设置环境变量便于使用 (建议开发环境使用)

```shell
MYSQL_DSN="db_user:db_password@tcp(localhost)/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接地址
GIN_MODE="debug"
HOST="https://i7f.cn/" #配置处理的域名
```

## 运行

```shell
go run main.go
```

项目运行后将监听 3000 端口
