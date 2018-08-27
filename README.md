# microservice_learning

- 是使用[Micro](https://github.com/micro/micro),[etcd](https://github.com/coreos/etcd),[nsq](https://github.com/nsqio/nsq),[grpc](https://github.com/grpc/grpc-go),[gin](https://github.com/gin-gonic/gin)等相关技术的一个微服务实例

```
 只有一个gateway服务器，api节点,db节点,log节点都可以自由添加
 #TODO 添加令牌服务器,链路追踪等功能,prometheus
 
 ```
![演示图](https://github.com/hwholiday/microservice_learning/blob/master/file/20180824104336.png) 

#### 启动命令
```
etcd

nsqlookupd

nsqd --lookupd-tcp-address=127.0.0.1:4160

nsqadmin --lookupd-http-address=127.0.0.1:4161

micro --registry=etcdv3   --broker=nsq   api  --handler=http

执行conf中的文件将配置信息读取到etcd中

启动api_agent,log_agent,db_agent

curl http://127.0.0.1:8080/api/v1/test 查看整个程序的运行情况
```



## 关于 Micro
- Micro 是一个微服务工具集。它被用来实现它的特性和接口，同时提供强大的可插拔的架构来保证基础组件可以被替换掉。
- Micro 专注于解决构建微服务系统的基础需求。它采用了深思熟虑地富有预见性的方式来实现它的设计。
- 如果你想深入研究 Micro 工具集请[点击这里](https://github.com/micro/micro)


## 联系

    QQ: 3355168235


