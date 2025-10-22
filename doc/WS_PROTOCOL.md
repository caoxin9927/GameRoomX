# WS_PROTOCOL — WebSocket 协议与状态机

## 连接
- 客户端携带 `token` 与 `room_id` 建立 WS：`ws://HOST/ws?token=xxx&room_id=yyy`。
- 握手成功后，服务端下发房间快照。

## 心跳
- 客户端每 5s 发送：`{ type: "ping", ts }`
- 服务端回应：`{ type: "pong", ts }`
- 超过 15s 未收到心跳，标记掉线。

## 消息格式（统一 envelope）
```json
{
  "seq": 12345,
  "type": "move",     // 事件类型：join/leave/move/chat/state/ack/error
  "room_id": "r1",
  "sender": "u1",
  "payload": { "x": 1.2, "y": 3.4 },
  "ts": 1730000000
}
```

## 广播策略
- 小房间：全量广播；
- 帧同步模式：每 tick 聚合输入，广播状态增量。

## 错误与重试
- 错误：`{ type: "error", code, message }`
- 服务器支持幂等 seq；重复包丢弃。
