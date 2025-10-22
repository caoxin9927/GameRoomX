# DEVELOPMENT_GUIDE — 开发指南

## 目标
提供统一的工程实践与协作方式，确保 GameRoomX 在“可运行、可验证、可维护”的前提下快速迭代。

## 分支策略
- `main`: 可发布分支（CI 通过；文档与版本号同步）。
- `dev`: 日常集成分支。
- `feature/*`: 功能开发分支。
- `hotfix/*`: 紧急修复。

## 提交规范（建议）
- 格式：`<type>(scope): message`
- type: feat | fix | refactor | perf | docs | test | chore | ci
- 例：`feat(room): add heartbeat and reconnect`

## 代码规范
- Go：`go fmt`, `golangci-lint`；
- C++：`clang-format`, `clang-tidy`；
- 命名与目录遵循“按领域划分”，API/WS/模型/存储/配置/指标分离。

## Definition of Done（DoD）
- 完成功能 + 单测/集成测试通过；
- 更新文档（API/协议/架构变化）；
- 更新 Grafana 面板截图（如涉及性能/指标变化）；
- 更新 CHANGELOG；
- CI 通过。

## 里程碑节奏（示例）
- M0（第1周）：工程骨架 + /healthz + Docker/CI；
- M1（第2周）：REST 登录/建房/入房 + WS 心跳/重连；
- M2（第3周）：压测指标 + Grafana 面板 + 火焰图；
- M3（第4周）：客户端 Demo + 线上演示；
- M4（第5-6周）：增强模块（网关/匹配/帧同步）；
- M5（第8周）：收尾与博客/面试材料。

## 评审清单（PR Checklist）
- [ ] 是否有测试覆盖？
- [ ] 是否更新相应文档？
- [ ] 是否影响性能/指标？有无截图与数据？
- [ ] 是否存在并发/资源泄漏风险？
- [ ] 是否符合编码规范与安全要求？
