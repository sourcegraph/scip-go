## 1. 实现映射函数

- [x] 1.1 在 `internal/symbols/symbols.go` 中添加 `GetSymbolKind` 函数
- [x] 1.2 实现包名符号（`*types.PkgName`）映射到 `Namespace`
- [x] 1.3 实现常量符号（`*types.Const`）映射到 `Constant`
- [x] 1.4 实现变量符号（`*types.Var`）映射逻辑
  - [x] 1.4.1 结构体字段映射到 `Field`（使用 `obj.IsField()` 判断）
  - [x] 1.4.2 其他变量映射到 `Variable`
- [x] 1.5 实现函数/方法符号（`*types.Func`）映射逻辑
  - [x] 1.5.1 方法映射到 `Method`（检查签名是否有接收者）
  - [x] 1.5.2 函数映射到 `Function`
- [x] 1.6 实现类型定义符号（`*types.TypeName`）映射逻辑
  - [x] 1.6.1 结构体映射到 `Class`（检查 underlying type）
  - [x] 1.6.2 接口映射到 `Interface`（检查 underlying type）
  - [x] 1.6.3 其他类型映射到 `Type`
- [x] 1.7 为未知类型添加默认处理（返回 `UnspecifiedKind`）
- [x] 1.8 在 `internal/symbols` 包中添加单元测试
  - [x] 1.8.1 测试函数/方法的区分
  - [x] 1.8.2 测试结构体/接口/类型的映射
  - [x] 1.8.3 测试字段/参数/变量的区分
  - [x] 1.8.4 测试未知类型的兜底处理

## 2. 集成映射函数

- [x] 2.1 修改 `internal/document/document.go::SetNewSymbol` 方法
  - [x] 2.1.1 为 `SymbolInformation` 添加 `Kind` 字段赋值
  - [x] 2.1.2 调用 `symbols.GetSymbolKind(obj)` 获取类型
- [x] 2.2 修改 `internal/document/document.go::SetNewSymbolForPos` 方法（如果使用）
  - [x] 2.2.1 为 `SymbolInformation` 添加 `Kind` 字段赋值
- [x] 2.3 修改 `internal/visitors/visitor_file.go::ToScipDocument` 方法
  - [x] 2.3.1 在创建局部符号时设置 `Kind` 字段
  - [x] 2.3.2 使用 `symbols.GetSymbolKind(local.Obj)` 获取类型
- [x] 2.4 修改 `internal/lookup/lookup.go::SetPkgName` 方法
  - [x] 2.4.1 为包名符号设置 `Kind` 为 `Namespace`

## 3. 更新和验证测试

- [x] 3.1 运行快照测试更新命令：`go test ./internal/index -update-snapshots`
- [x] 3.2 审查快照 diff
  - [x] 3.2.1 确认所有符号的 `Kind` 字段已正确设置
  - [x] 3.2.2 验证只有 Kind 字段改变，其他字段保持不变
- [x] 3.3 验证关键符号类型
  - [x] 3.3.1 函数符号正确设置为 `Function`
  - [x] 3.3.2 方法符号正确设置为 `Method`
  - [x] 3.3.3 结构体符号正确设置为 `Class`
  - [x] 3.3.4 接口符号正确设置为 `Interface`
  - [x] 3.3.5 常量符号正确设置为 `Constant`
  - [x] 3.3.6 字段符号正确设置为 `Field`
  - [x] 3.3.7 变量符号正确设置为 `Variable`
  - [x] 3.3.8 包名符号正确设置为 `Namespace`
- [x] 3.4 运行完整测试套件：`go test ./...`
- [x] 3.5 修复任何失败的测试

## 4. 代码审查和文档

- [x] 4.1 对新添加的 `GetSymbolKind` 函数进行代码审查
- [x] 4.2 验证所有修改点的代码风格一致性
- [x] 4.3 确保没有引入新的 lint 警告
- [x] 4.4 检查文档注释的完整性和准确性
- [x] 4.5 提交代码前进行最终验证

## 5. 可选任务（根据测试结果决定）

- [x] 5.1 为泛型类型添加特殊处理（如果快照中包含泛型）
- [x] 5.2 添加日志输出以记录未分类的符号类型（用于调试）
- [x] 5.3 性能优化：检查 `GetSymbolKind` 函数的性能影响
