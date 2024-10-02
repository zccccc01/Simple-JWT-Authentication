# 简易的登录功能实现
## 使用前提
```
go version 1.23
python version 3.12
mysql version 8.0
```


## 安装后端go的包
```
go get github.com/gofiber/fiber/v2
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/cosmtrek/air
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
```
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
