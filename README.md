<h1 align="center">DAO - 经济数据库</h1>
<p align="center"><b>根据公开信息实现的本地指标系统，方便查看和编写指标，追踪经济走势</b></p>

[English](./README.EN.md) | 简体中文

# 目录

- [TODO List](#TODOList)
- [开始](#开始)
  - [前提条件](#前提条件)
  - [安装](#安装)
  - [启动](#启动)
  - [使用](#使用)
  - [目录](#目录)
- [感谢](#感谢)
- [SwaggerUi](#SwaggerUi)
- [License](#license)

## TODOList

| Name                                                              | ⭕ - ✔️|
| -----------                                                       | ----------- |
| 沪深300市盈率                                                      | ✔️       |
| 中美国债2、5、10、30年收益率                                        | ✔️        |
| 社会融资总量 Total social financing                                | ✔️       |
| 社会融资存量  Total social Stock                                   | ✔️        |
| 货币供应 M0 M1 M2（Money Supply）                                  | ✔️       |
| 社会消费品零售总额 Retail Sales of Consumer Goods                   | ✔️        |
| 期货 伦铜:CAD                                                      | ✔️       |
| PPI                                                                | ✔️        |
| GDP                                                                | ✔️       |
| CPI                                                                | ✔️        |
| 国家统计局制造业采购经理人指数 NBS manufacturing PMI                  | ✔️        |
| 国家统计局非制造业采购经理人指数 NBS non-manufacturing PMI             | ✔️       |
| 财新制造业采购经理人指数 Caixin manufacturing PMI                     | ⭕        |
| 财新服务采购经理人指数 Caixin service PM                              | ⭕        |
| 工业生产 Industrial Production                                       | ⭕        |
| 电力生产与消费 Electricity Production and Consumption                 | ⭕        |
| 商品贸易 Merchandise Trade                                           | ⭕       |
| 政府收支及收支平衡 Government Revenue, Expenditure and Balance        | ⭕        |
| 国际收支平衡 Balance of Payments                                      | ⭕        |
| 存贷款 Loan and Deposit                                               | ⭕        |
| 房屋销售 Home Sales                                                   | ⭕        |
| 外汇储备 Foreign Exchange Reseres                                     | ⭕        |
| 铁路货运 Rail Freight Traffic                                         | ⭕        |
| 服务贸易 Service Trade                                                | ⭕        |
| 工业利润 Industrial profits                                           | ⭕        |
| 国内生产总值平减指数 GDP Deflator                                       | ⭕        |
| 库存房 Home inventory                                                  | ⭕       |
| 房屋开工、在建和竣工 Housing Starts, Under Construction and Completion  | ⭕        |
| 土地收购、交易和发展 Land Acquisition, Transaction and Development      | ⭕        |
| 地价指数 Land Price Indices                                           | ⭕        |
| 固定资产投资 Fixed Asset Investment                                   | ⭕       |
| 汽车销售 Auto Sales                                                   | ⭕        |
| 物业价格指数 Property Price Index                                     | ⭕        |
| 房地产开发投资 Investment in Real Estate Development                  | ⭕       |
| 家庭收入调查 Household Income Survey                                  | ⭕        |
| 家庭支出调查 Household Expenditure Survey                             | ⭕        |
| 国家外债 External Debt                                                | ⭕        |
| 就业总人数 Total Employment                                           | ⭕       |
| 城镇登记失业率 Urban Registered Unemployment Rate                     | ⭕        |
| 城镇单位职工平均/总工资 Average/total wage of Employees in Urban Units | ⭕        |
| 农民工平均收入 Average income of migrant workers                      | ⭕        |

## 开始

要启动本项目，你需要以下步骤

### 前提条件

chromedp库，需要先安装chrome

```bash
### 下载源加入到系统的源列表
sudo wget <http://www.linuxidc.com/files/repo/google-chrome.list> -P /etc/apt/sources.list.d/

# 导入谷歌软件公钥
wget -q -O - <https://dl.google.com/linux/linux_signing_key.pub>  | sudo apt-key add -

### 更新列表
sudo apt-get update

## 安装chrome
sudo apt-get install google-chrome-stable
```

### 安装

   ```sh
   git clone https://github.com/chenhaonan-eth/dao.git

   cd dao

   make build
   ```

### 启动

   ```sh
   make start
   ```

### 使用

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

## 目录说明

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

## SwaggerUi

```bash
http://0.0.0.0:50053/openapiv2/server.swagger.json

http://0.0.0.0:50053/swagger-ui/
```

## 感谢

[Akshare](<https://www.akshare.xyz/index.html>)

数据均来自网络公开数据。

## License

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## 表示您的支持

如果这个项目带给您一点帮忙，请点个 ⭐️
