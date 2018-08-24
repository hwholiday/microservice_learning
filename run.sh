#!/usr/bin/env bash
nohup etcd >/dev/null 2>&1 &
echo "启动etcd成功"
nohup nsqlookupd >/dev/null 2>&1 &
echo "启动nsqlookupd成功"
nohup nsqd --lookupd-tcp-address=127.0.0.1:4160 -tcp-address=0.0.0.0:4152 -http-address=0.0.0.0:4153 >/dev/null 2>&1 &
echo "启动nsqd成功"
nohup nsqd --lookupd-tcp-address=127.0.0.1:4160 >/dev/null 2>&1 &
echo "启动nsqd成功"
nohup nsqadmin --lookupd-http-address=127.0.0.1:4161 >/dev/null 2>&1 &
echo "启动nsqadmin成功"
nohup micro --registry=etcdv3   --broker=nsq   api  --handler=http >/dev/null 2>&1 &
echo "启动micro成功"