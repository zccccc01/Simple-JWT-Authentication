# 简易的登录功能实现
## 使用前提
```
go version 1.23
python version 3.12
mysql version 8.0
```


## 安装后端go的包
```
go mod tidy
```
## 注意
记得看go-auth/database/connection.go里面的注释  
`username改为自己的MySQL用户名,passwd改为自己的密码,yourdb改为你使用的库`


## 后端启动
```bash
cd ./go-auth
air -c .air.toml
```
## 前端启动
```bash
cd ./static
python -m http.server 3000
```
## 浏览器打开
`http://localhost:3000`

即可看到界面,实现简易的登录,登出,注册功能  
**注:我这里没有使用react去构建前端**

## learn from
[youtube教学视频](https://youtu.be/d4Y2DkKbxM0?si=BNUstHwr6lPhwjSJ)
freeCodeCamp  
React and Golang JWT Authentication - Tutorial
