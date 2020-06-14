# gtodo
> gtodo 是 go 语言实现的一个 todo 应用的 restful api 服务，使用 Gin+Vue。

gtodo 建立的目的是为了学习 go web 开发，使用了前后端分离的方式，使用 JWT 进行身份鉴权，
前端使用 Vue+Antd，后端使用 Gin，数据库ORM 使用 Gorm。

## 展示
![screentshot](todogif.gif)

## how to run
```
> git clone https://github.com/shiniao/mini_todo.git

# mysql create database:
> create database todo;

# Then change the file config.yaml for your needs.

# run server
> cd server
> go run main.go

# run client
cd client
npm install
npm run serve
```

## todos api

You can visit `.../swagger/index.html` for more detail.

First, use todos api you need to get a token, Then take tokens with each visit:

```shell script
POST /token get a token, need take `key:matata`
```

That's all api:
```
GET    /v1/todos 
GET    /v1/todos/:id
POST   /v1/todos
PUT    /v1/todos/:id
DELETE /v1/todos/:id
```