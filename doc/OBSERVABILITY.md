# OBSERVABILITY — 可观测性

## 指标（Prometheus）
- `room_online_total`、`member_online_total`
- `ws_msg_in_total`、`ws_msg_out_total`
- `ws_broadcast_latency_ms_bucket`
- `auth_fail_total`、`errors_total`

## 日志
- 结构化 JSON：时间、级别、模块、user_id、room_id、seq；
- 慢日志：广播 > 阈值记录 payload 大小与耗时。

## Grafana
- 提供 Dashboard JSON（放在 deploy/grafana/）；
- 面板：在线量、吞吐、延迟、错误率、GC/CPU/内存。
