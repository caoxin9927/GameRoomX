# RUNBOOK — 运维手册

## 常用操作
- 查看在线/房间：Grafana 面板；
- 强制踢人/解散房间：管理接口；
- 回滚：执行 deploy/script/rollback.sh。

## 故障排查剧本
1) 延迟飙升：看 ws_broadcast_latency 直方图 → 火焰图定位 → 检查热点房间。
2) 连接大量掉线：检查心跳超时配置与 Redis 网络。
3) CPU 打满：检查日志级别与批量广播开关。

## 紧急联系人与升级窗口（自填）。
