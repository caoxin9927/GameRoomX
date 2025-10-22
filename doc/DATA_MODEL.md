# DATA_MODEL — 数据模型与 Redis 设计

## Key 命名
- `room:{room_id}` → 房间元信息（hash）
- `room:{room_id}:members` → 成员集合（set）
- `user:{user_id}` → 用户信息（hash）
- `rooms:active` → 活跃房间列表（zset/sorted by ts）

## 结构示例
- room hash:
```
{ id, name, capacity, owner, created_at, updated_at, status }
```

## 过期与清理
- 长时间无成员的房间定期清理；
- 断线重连的会话 TTL；

## 一致性
- 通过事务/Lua 脚本保证“容量检查 + 加入”原子性。
