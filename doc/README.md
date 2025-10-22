# GameRoomX 文档目录

> 最近更新：2025-10-22

这是一套可直接使用的开发文档骨架（Markdown）。建议将整个文件夹放到你的仓库 `docs/` 目录下，或导入到 Notion/Obsidian。

**如何使用：**
1. 从 `DEVELOPMENT_GUIDE.md` 开始阅读与维护规范；
2. 优先补全 `ARCHITECTURE.md`、`API.md`、`WS_PROTOCOL.md`、`TEST_PLAN.md`；
3. 每完成功能/修复/调优，更新 `CHANGELOG.md` 与相应小节；
4. 每次关键设计决策写一份 `ADR`（Architecture Decision Record）。

**核心文档：**
- `DEVELOPMENT_GUIDE.md`：开发流程、分支策略、代码规范、DoD、里程碑；
- `ARCHITECTURE.md`：总体架构、模块边界、依赖图、关键时序；
- `API.md`：REST 接口规范；
- `WS_PROTOCOL.md`：WebSocket 协议、消息格式、状态机；
- `DATA_MODEL.md`：数据模型（Redis key 设计、结构定义）；
- `CONCURRENCY.md`：并发/线程/协程模型、锁策略；
- `OBSERVABILITY.md`：日志、指标、追踪，Grafana 面板；
- `TEST_PLAN.md`：单测/集成/压测计划与用例；
- `PERF_PLAN.md`：性能目标、压测方法、优化记录；
- `DEPLOYMENT.md`：环境、配置、部署与回滚；
- `RUNBOOK.md`：运维手册、故障排查剧本；
- `ADR/`：架构决策记录；
- `ROADMAP.md`：路线图与里程碑；
- `CHANGELOG.md`：变更历史；
- `TASK_BOARD.md`：任务看板（可复制到 Notion/Trello）。
