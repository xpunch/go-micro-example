# Hello World

### 运行我们的Hello World程序
```
go run .
```

我们使用Micro工具程序查看我们刚运行的服务程序，可以使用命令行或者Web网页查看
### 一、运行Micro服务
```
micro
```

### 使用命令行打印正在运行的服务，可以看到我们刚才启动的helloworld服务
```
micro services
```

### 二、运行Micro Web
```
micro web
```

打开[localhost:8082/registry](http://localhost:8082/registry)就可以看到所有正在运行的服务程序，打开[helloworld](http://localhost:8082/registry?service=helloworld)可以查看我们刚才运行的程序详细信息
### Metadata
```
broker=http protocol=grpc registry=mdns server=grpc transport=grpc
```
可以看到不同模块当前使用的组件
