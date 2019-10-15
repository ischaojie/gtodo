# A todos restful-api service base on go language.
> gin+gorm+vue

just for fun,  for learn.

![截图](todogif.gif)

## api

访问`[host:port]/swagger/index.html`查看api详情

```shell script
POST /token 申请token, 需携带key:matata
```
### todos api
```
# 访问todos api 需携带token
GET    /v1/todos 
GET    /v1/todos/:id
POST   /v1/todos
PUT    /v1/todos/:id
DELETE /v1/todos/:id
```
### health api: api状态检测
```shell script
GET /sd/health  api status
GET /sd/cpu     cpu status
GET /sd/ram     ram usage
GET /sd/disk    disk usage
```

### how to run
```
> git clone https://github.com/shiniao/mini_todo.git

# mysql创建数据库
> create database todo;

# 根据需要修改config.yaml文件内容

# server
> cd server
> go run main.go

# client
cd client
npm install
npm run serve
```

### mini系列
[mini_search](https://github.com/shiniao/mini_search)——hakuna电影搜索引擎(elasticsearch+flask+vue)

[mini_sms_classify](https://github.com/shiniao/mini_sms_classify)——小型垃圾邮件分类系统（naive_bayes+flask+vue）

[mini_mnist](https://github.com/shiniao/mini_mnist)——mini手写数字识别(CNN+flask+vue)

### 联系我
email: zhuzhezhe95@gmail.com