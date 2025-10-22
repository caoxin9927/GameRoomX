# CONCURRENCY — 并发模型

## Room Server
- Go：每连接 goroutine + 带缓冲 channel；
- C++：Reactor(epoll) + 线程池；
- 每个房间维护独立的发送队列/状态机；

## 写事件集中化（C++ 增强）
- 同一 fd 的写事件仅由 owner 注册/注销 EPOLLOUT；
- 直到 EAGAIN 撤销 EPOLLOUT，避免事件风暴；

## 竞态与幂等
- 使用递增 seq 与 ACK；
- 重要路径加 CAS/锁，读多写少用 RWLock。

## 定时任务
- 最小堆/时间轮（心跳、超时、清理）。
