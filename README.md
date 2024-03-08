# Go语言 Gin+Vue 前后端分离实战 - OceanLearn
## Go语言 Gin+Vue 前后端分离实战 - OceanLearn】https://www.bilibili.com/video/BV1CE411H7bQ?vd_source=4edab2c311fa4f49601588d3604bc273
本教程是 Gin + Vue 前后端分离项目实战， 课程采用循序渐进的方式，有简单到复杂逐步实现MVC架构
项目适合人群：有一定编程经验的程序猿，在开始学习本视频之前你需要具备一下技能
1.基本的go语言语法
2.javasc
3.git基础
你能学到什么
1. gin 框架的使用
2. jwt 与 中间件 在go 项目中的运用
3. 渐进式的项目优化 gin 项目 MVC
4. cors处理
5. vue 的项目搭建
6. vue 的路由中间件
7. vuex 的使用

## 环境
go 版本: go version go1.21.4 windows/amd64
安装其他版本命令 : 
```go
go get golang.org/dl/go1.13.8
go $GOPATH/bin/go1.13.8 download
// 指定 go 版本执行文件
go $GOPATH/bin/go1.13.8 run filename
```
node 版本 v14.21.3

可以使用 nvm 下载和切换指定版本的 node

yarn 版本 1.22.21

vue-V  @vue/cli 5.0.8

数据库 mysql

后端配置文件在 `config/application.yml` 文件中， 参考与如下, 根据自己本地修改
```yml
server:
  port: 1016 // 后端端口号
datasources:
  driverName: mysql // 数据库驱动类型
  host: 127.0.0.1
  port: 3306
  database: ginessential // 数据库名称
  username: root // 本地数据库账号
  password: root // 本地数据库密码
  charset: utf8mb4
```
## 运行

后端项目运行

```go
# 进入项目路径
cd BackEnd
go run routes.go main.go
// 或者
go build
./Blog
```
前端项目运行

```
# 进入前端项目目录
cd  frontend
# 安装项目依赖
yarn install
# 运行项目
yarn serve
```
