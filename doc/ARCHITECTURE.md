# ARCHITECTURE — 总体架构

## 系统目标
- 房间管理、实时广播、心跳与断线重连；
- 指标可观测、易压测、易部署。

## 组件
- Gateway（可选，C++/epoll）：多协议接入、限流、熔断；
- Room Server（Go/C++）：REST、WS、房间逻辑；
- Redis：元数据存储、Pub/Sub；
- Observability：Prometheus + Grafana；
- Client Demo：Unity/团结引擎。

## 依赖与边界（示意）
```
[Client] <--WS/HTTP--> [Gateway] <--RPC/WS--> [Room] <---> [Redis]
                                         \--> [Prometheus/Grafana]
```

## 关键时序（加入房间与广播）
1. 客户端 Auth（游客/JWT）
2. 创建/加入房间（REST）
3. 建立 WS，服务端下发当前房间状态
4. 客户端发送输入/状态 → 服务端广播给房间内成员
5. 心跳与断线重连：重连后补发状态快照

## 非功能性
- SLO：P95 < 60ms, P99 < 120ms（2k 连接，10min）；
- 可观测：关键指标 + 结构化日志；
- 安全：限频/黑名单/JWT 有效期短；
- 伸缩：后续支持分布式匹配与房间迁移。
