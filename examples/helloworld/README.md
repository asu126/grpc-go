### protobuf 的安装与使用
- 1.1. 载源码安装包
```
https://github.com/protocolbuffers/protobuf
```

- 1.2 解压
```
tar -zxvf protobuf-2.4.1.tar.gz
```

- 1.3 编译/安装
```
cd protobuf-2.4.1
（可以参考README思路来做。）
./configure
make
make check  (check结果可能会有错误，但不用管她，因为暂时那些功能用不到)
make install
（完了之后会在 /usr/local/bin 目录下生成一个可执行文件 protoc
```

- 1.4 检查安装是否成功
protoc --version
如果成功，则会输出版本号信息。如果有问题，则会输出错误内容。

- 1.5 错误及解决方法
protoc: error while loading shared libraries: libprotoc.so.8: cannot open shared
错误原因：
protobuf的默认安装路径是/usr/local/lib，而/usr/local/lib 不在Ubuntu体系默认的 LD_LIBRARY_PATH 里，所以就找不到该lib
解决方法：
1). 创建文件 /etc/ld.so.conf.d/libprotobuf.conf，在该文件中输入如下内容：
/usr/local/lib
2). 执行命令
sudo ldconfig
这时，再运行protoc --version 就可以正常看到版本号了

- 1.6 安装protoc-gen-go
```
go get -u github.com/golang/protobuf/protoc-gen-go
export PATH=$PATH:$GOPATH/bin
```

[参考地址](https://blog.csdn.net/yahstudio/article/details/48995077)


### 生成客户端和服务器端代码
```
protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
```


### gRPC 三种方式
方式      |  server 端　 | client 端
--       |  --------   | ----
简单 RPC  |   -         | -
服务端 |   Send()    | Recv
客户端 |   Recv(),SendAndClose()    | Send(),CloseAndRecv()
双向  |   Recv(), Send()    | Recv(), Send()

[Google Protocol Buffer 的使用和原理](https://www.ibm.com/developerworks/cn/linux/l-cn-gpb/index.html)
[官方文档](https://godoc.org/google.golang.org/grpc)
