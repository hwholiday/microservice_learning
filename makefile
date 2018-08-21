
　protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. *.proto

　//go run main.go –registrer=etcdv3 –registrer-address=http://127.0.0.1:2379

　//go run main.go –registrer=etcdv3 –registrer-address=http://127.0.0.1:2379 -broker=nsq -broker-address=http://127.0.0.1:4153

　etcd

  nsqlookupd

　nsqd --lookupd-tcp-address=127.0.0.1:4160 -tcp-address=0.0.0.0:4152 -http-address=0.0.0.0:4153

　nsqadmin --lookupd-http-address=127.0.0.1:4161