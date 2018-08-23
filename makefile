
　protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. *.proto

　etcd

  nsqlookupd

　nsqd --lookupd-tcp-address=127.0.0.1:4160 -tcp-address=0.0.0.0:4152 -http-address=0.0.0.0:4153

　nsqadmin --lookupd-http-address=127.0.0.1:4161

  micro --registry=etcdv3   --broker=nsq   api  --handler=http
