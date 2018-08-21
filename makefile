
 protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. *.proto

 go run main.go –registrer=etcdv3 –registrer-address=http://127.0.0.1:2379
