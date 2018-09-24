# 短网址生成服务

提供短网址生成和短网址跳转服务

[Demo](http://u.liuin.cn/)

## 设计

采用id对应一个网址的形式，加入一个网址时生成一个与之对应的id后存入数据库。一个网址可能对应多个id，为了减少空间浪费，采用Redis缓存最近经常被转换的长网址。

### id生成

参考分布式id生成思想，id由时间序列+随机数+机器标识组成

### 缓存

使用两个缓存，一个是`short url`转`long url`的时候的缓存。另一个是`long url`转`short url`的缓存，减少一个长网址可能对应多个短网址所造成的空间浪费

## 接口

提供`long url`转`short url`的api接口：
```
url: http://u.liuin.cn 
method: POST

param:
url: string required  # 需要转换的长网址

response:
{
	"code": integer,     # 状态码，0为成功，其他为失败
	"message": string,   # 解释
	"id_str": string,    # 对应的短URL
}
```

## 部署

### 准备

首先需要搭建好需要的MySQL数据库环境和Redis环境

然后再`conf`文件夹下添加`conf.ini`配置文件，格式如下：
```
[mysql]
user=[user]
passwd=[passwd]
host=[host]
port=[port]
database=surl
[redis]
host=[host]
port=[port]
```

### 运行

使用Docker部署，监听本地`18080`端口
``` bash
docker-compose up -d
```