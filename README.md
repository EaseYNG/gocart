 # 项目概述

## 项目名称

gocart - 基于Gin+Redis+MySQL的高并发电商秒杀系统

## 技术栈

- 后端：Gin、Redis、MySQL

## 业务功能

1. 商品秒杀活动配置
2. 用户秒杀下单
3. 限制每人限购1件
4. 高并发下防库存超卖
5. 接口限流防刷
6. 异步下单削峰（协程 + 队列）

# 项目结构

```
gocart/
├── main.go         # 入口、初始化路由、启动消费协程
├── model           # 商品、订单、用户结构体
├── config          # MySQL、Redis配置
├── router          # 路由分组
├── controller      # 秒杀接口入口
├── service         # 秒杀业务逻辑（核心高并发逻辑全在这）
├── middleware      # 限流、跨域、鉴权中间件
├── util            # Redis工具、加密、分布式锁工具
```

