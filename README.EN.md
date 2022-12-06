
<div align="center">
  <h1 align="center">DAO - Local economic database</h1>
<p align="center"><b>Local indicator system based on publicly available information for easy viewing and compiling of indicators to track economic trends</b></p>
</div>

English | [简体中文](./README.md)

# Table of Contents
- [Getting Started](#getting-started)
  - [Setup](#setup)
  - [Install](#install)
  - [Run](#run)
  - [Usage](#usage)
  - [Directory](#directory)
- [Swagger Ui](#swagger-ui)
- [Thanks](#thanks)
- [License](#license)

## 💻 Getting Started<a id="getting-started"></a>

To get a local copy up and running, follow these steps.

### Prerequisites

In order to run this project you need:

install chrome:

```sh
 sudo wget <http://www.linuxidc.com/files/repo/google-chrome.list> -P /etc/apt/sources.list.d/

 wget -q -O - <https://dl.google.com/linux/linux_signing_key.pub>  | sudo apt-key add -

 sudo apt-get update

 sudo apt-get install google-chrome-stable
```

```sh
   go get -u github.com/bufbuild/buf/cmd/buf
   go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
   go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
   go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
   go get -u google.golang.org/protobuf/cmd/protoc-gen-go
```

### Setup

Clone this repository to your desired folder:

```sh
  cd my-folder
  git clone https://github.com/chenhaonan-eth/dao.git
```

### Install

Install this project with:

```sh
  cd dao
  make build
```

### Run

To run the project, execute the following command:

```sh
    make start
```

### Usage

   ```sh
    # HTTP
    curl -X GET http://localhost:50053/v1/macroscopic/FuturesForeignHist/CAD
   ```

   ```go
    // gRPC
    package main

    import (
        "log"

        pb "github.com/chenhaonan-eth/dao/proto/server"
        "golang.org/x/net/context"
        "google.golang.org/grpc"
        "google.golang.org/grpc/credentials/insecure"
    )

    func main() {
        conn, err := grpc.Dial(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
        if err != nil {
            log.Fatal(err)
        }
        defer conn.Close()
        c := pb.NewGreeterClient(conn)
        context := context.Background()
        body := &pb.FturesFoewignRequest{
            Symbol: "CAD",
        }

        r, err := c.GetFuturesForeignHist(context, body)
        if err != nil {
            log.Println(err)
        }
        log.Println(r.Results)
    }
```

## Directory

   ```yaml
    ├─ biz            // 业务处理
    ├─ client         // gRPC 客户端演示
    ├─ cmd
    │  ├─ app         //cobra
    │  ├─ gormgen     //生成DB代码
    ├─ config         //zaplog、viper
    ├─ dal
    │  ├─ initialize  //初始化数据库
    │  ├─ model       //存放所有数据模型
    │  ├─ query       //gorm-gen生成的CURD代码
    ├─ economic       //经济指标数据获取
    │  ├─ ftures      //期货
    │  ├─ macroscopic //宏观经济
    ├─ pkg            //外部包
    ├─ proto          //gRPC、grpc-gateway生成的代码
    ├─ server         //gRPC、grpc-gateway服务
    ├─ spider         //定时爬虫更新数据
   ```

## Swagger Ui<a id="swagger-ui"></a>

```bash
http://0.0.0.0:50053/openapiv2/server.swagger.json

http://0.0.0.0:50053/swagger-ui/
```

## Thanks

[Akshare](<https://www.akshare.xyz/index.html>)

数据均来自网络公开数据。

## License

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

- [MIT](https://opensource.org/licenses/MIT)

## Show your support

Give a ⭐️ if this project helped you!