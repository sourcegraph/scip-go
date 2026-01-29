## Context

SCIP 协议定义了 `SymbolInformation_Kind` 枚举来分类不同类型的符号（如 `Function`、`Variable`、`Class` 等）。当前 scip-go 实现虽然创建了 `SymbolInformation` 对象，但从未填充其 `Kind` 字段，导致所有符号都默认为 `UnspecifiedKind`（值为 0）。

**当前状态**：
- 代码中存在三个创建 `SymbolInformation` 的主要位置：
  1. `internal/document/document.go::SetNewSymbol` - 全局符号定义
  2. `internal/visitors/visitor_file.go::ToScipDocument` - 局部符号（函数内变量等）
  3. `internal/lookup/lookup.go::SetPkgName` - 包名符号
- 所有位置都使用相同的初始化方式：`&scip.SymbolInformation{Symbol: ..., Kind: 0}`

**约束条件**：
- 必须使用 SCIP SDK 提供的枚举值
- 需要准确映射 Go 类型系统到 SCIP 类型系统
- 必须保持现有快照测试的结构，仅更新 Kind 字段
- 不应引入性能敏感的操作（如对每个符号的复杂 AST 遍历）

## Goals / Non-Goals

**Goals:**
- 为所有输出的符号正确设置 `Kind` 字段
- 实现一个可维护的映射函数，明确 Go `types.Object` 到 SCIP `SymbolInformation_Kind` 的转换规则
- 通过快照测试验证所有符号类型的正确性

**Non-Goals:**
- 修改 SCIP 协议或添加新的枚举值
- 改变符号命名策略或层次结构
- 优化索引性能（仅关注正确性）

## Decisions

### 1. 映射函数位置和签名

**决策**：在 `internal/symbols/symbols.go` 包中创建 `GetSymbolKind` 函数

**理由**：
- `symbols` 包已存在，职责是处理 SCIP 符号相关的工具函数
- 避免在多个文件中重复映射逻辑（DRY 原则）
- 该包已有 `FromDescriptors` 和 `RangeFromName` 等 SCIP 辅助函数，放置映射函数符合现有模式

**函数签名**：
```go
// GetSymbolKind 根据 Go 类型对象返回对应的 SCIP 符号类型
func GetSymbolKind(obj types.Object) scip.SymbolInformation_Kind
```

### 2. Go 类型到 SCIP Kind 的映射规则

**决策**：采用以下映射表

| Go `types.Object` 类型 | SCIP `SymbolInformation_Kind` | 备注 |
|----------------------|------------------------------|------|
| `*types.PkgName` | `Namespace` | 包名符号 |
| `*types.Const` | `Constant` | 常量 |
| `*types.Var` - 字段 | `Field` | 结构体字段（通过 `obj.IsField()` 判断） |
| `*types.Var` - 参数 | `Parameter` | 函数/方法参数（需要在调用方判断） |
| `*types.Var` - 其他 | `Variable` | 局部变量、全局变量 |
| `*types.Func` - 方法 | `Method` | 通过 `obj.Type().(*types.Signature).Recv()` 是否存在判断 |
| `*types.Func` - 函数 | `Function` | 非方法的函数 |
| `*types.TypeName` - Struct | `Class` | `obj.Type().Underlying().(*types.Struct)` |
| `*types.TypeName` - Interface | `Interface` | `obj.Type().Underlying().(*types.Interface)` |
| `*types.TypeName` - 其他 | `Type` | 其他类型定义（如 type alias、基本类型等） |

**理由**：
- SCIP 的 `Class` 对应 Go 的 struct（虽然 Go 没有类的概念，但 struct 是最接近的构造）
- `Method` 和 `Function` 的区分对代码导航很重要（方法与接收者关联）
- 字段、参数、变量的区分有助于理解代码结构

**考虑的替代方案**：
- 方案 A：在 `SetNewSymbol` 等函数内直接根据 `types.Object` 类型判断
  - *缺点*：逻辑分散，难以维护
- 方案 B：使用 visitor 模式为每种类型创建专门的处理器
  - *缺点*：过度设计，增加复杂度

### 3. 特殊情况处理

**决策**：对于无法明确分类的对象，返回 `UnspecifiedKind` 而非 panic

**理由**：
- Go 类型系统可能存在 scip-go 未覆盖的边缘情况
- 保持索引器稳定性，不应因未知类型崩溃
- 便于后续调试（可以在日志中记录未分类的类型）

**处理策略**：
```go
default:
    return scip.SymbolInformation_UnspecifiedKind
```

### 4. 快照测试更新策略

**决策**：运行 `go test ./internal/index -update-snapshots` 批量更新所有快照

**理由**：
- 快照数量较多，手动更新效率低
- 修改逻辑后，所有符号的 Kind 字段都会改变，需要全局更新
- 快照测试能自动验证新逻辑的正确性

**注意事项**：
- 审查 diff 确保只有 Kind 字段改变
- 如果出现意外差异，可能需要新增测试用例

## Risks / Trade-offs

### 风险 1：映射不完整导致某些符号被错误分类

**缓解措施**：
- 在实现后运行完整测试套件
- 检查快照 diff，关注 Kind 字段的值是否符合预期
- 可以添加日志输出未分类的符号类型，用于后续完善

### 风险 2：方法 vs 函数的判断逻辑复杂

**缓解措施**：
- 通过检查 `types.Signature` 的 `Recv()` 字段来判断是否为方法
- 该逻辑在 `internal/visitors/visitor_func.go::receiverTypeName` 中已有实践，可参考
- 添加单元测试覆盖各种方法签名情况（值接收者、指针接收者等）

### 权衡：增加代码复杂度 vs 提升索引质量

**决策**：优先考虑索引正确性，接受适量的代码复杂度

**说明**：
- 虽然增加了 `GetSymbolKind` 函数和多个类型判断，但逻辑集中且可测试
- 对 SCIP 用户的价值远超实现成本

## Migration Plan

1. **实现映射函数**
   - 在 `internal/symbols/symbols.go` 中添加 `GetSymbolKind` 函数
   - 添加单元测试覆盖各种 Go 类型

2. **集成映射函数**
   - 修改 `internal/document/document.go::SetNewSymbol`，添加 `Kind: symbols.GetSymbolKind(obj)`
   - 修改 `internal/visitors/visitor_file.go::ToScipDocument`，为局部符号设置 Kind
   - 修改 `internal/lookup/lookup.go::SetPkgName`，设置包名符号的 Kind

3. **更新测试**
   - 运行 `go test ./internal/index -update-snapshots`
   - 审查快照 diff
   - 手动验证关键符号类型的 Kind 值

4. **验证和修复**
   - 运行完整测试套件：`go test ./...`
   - 修复任何失败的测试
   - 提交代码前进行代码审查

## Open Questions

1. **Q**: SCIP 的 `Class` 类型是否准确映射 Go 的 struct？
   - **A**: 虽然 Go 没有 OOP 的 class 概念，但 struct 在语义上最接近 class。 SCIP 协议中 `Class` 被广泛用于表示类型定义，因此此映射是合理的。

2. **Q**: 如何处理 type alias（类型别名）？
   - **A**: type alias 也是 `*types.TypeName`，其 underlying type 决定了实际类型（struct、interface 等）。因此会根据 underlying type 分类，这可能符合预期。

3. **Q**: 需要为泛型添加特殊处理吗？
   - **A**: Go 1.18+ 引入了泛型，但在 scip-go 当前的快照中未包含泛型案例。如果泛型符号出现在索引中，当前逻辑会将其分类为 `Type`，这是可接受的默认行为。后续可以根据需求细化。
