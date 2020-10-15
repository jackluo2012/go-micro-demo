### go-micro 微服务练习

描述：gomicro微服务的demo，实现了web服务和user服务以及数据库的增删查改

运行要求：

1）修改/share/config/config.go中的MysqlDSN，将用户名和密码改成自己的

2）运行dbhelper下的db.go，生成相应的数据库和表

### 启动user 的服务
```
cd dbhelper
go run db.go
cd user-srv
//启动go-micro 自带的注册中心
go run main.go —registry=mdns

```
### 启动web服务
```
cd web
go run main.go 
```
### 测试方法，可以参加一下 导入postmain 中