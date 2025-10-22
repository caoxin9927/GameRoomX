# API — REST 接口规范（草案）

## 约定
- BaseURL: `http://HOST:8080/api/v1`；
- Header: `Authorization: Bearer <JWT>`（游客可选）。
- 响应：统一包裹 `{ code, message, data }`。

## 登录
- POST `/auth/guest`
  - req: `{ device_id: string }`
  - res: `{ token: string, user_id: string, expire: int }`

## 创建房间
- POST `/rooms`
  - req: `{ name: string, capacity: int }`
  - res: `{ room_id: string }`

## 加入房间
- POST `/rooms/{room_id}/join`
  - res: `{ room_id, members: [...], ws_url }`

## 离开房间
- POST `/rooms/{room_id}/leave`

## 房间列表
- GET `/rooms?status=active&page=1&page_size=20`
  - res: `{ list: [...], total: int }`
