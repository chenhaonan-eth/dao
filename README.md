# Dao

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

根据公开信息实现的股票指标系统，方便查看和编写指标，追踪经济走势

数据-接口参考：<https://www.akshare.xyz/index.html>

所有数据均来自网络公开数据，本库只为学习交流。

## 使用前提

因为使用了chromedp库，所以Ubuntu使用需要事先安装chrome

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

### Swagger

```bash
http://0.0.0.0:50053/openapiv2/server.swagger.json

http://0.0.0.0:50053/swagger-ui/
```
