## Why

SCIP 索引器在输出 `SymbolInformation` 时，其 `Kind` 字段始终为 0（即 `UnspecifiedKind`），导致索引结果中无法区分函数、变量、类型等不同符号类型，降低了代码导航和语义分析的用户体验。

## What Changes

- **添加符号类型映射函数**：在代码库中实现从 Go `types.Object` 到 SCIP `SymbolInformation_Kind` 的映射逻辑
- **修改 SymbolInformation 创建点**：在以下位置调用映射函数并设置 `Kind` 字段
  - `internal/visitors/visitor_file.go` - 局部符号（local symbols）
  - `internal/document/document.go` - 文档级符号（package/struct/interface/function 等）
  - `internal/lookup/lookup.go` - 包名符号（package names）
- **增强快照测试**：更新现有快照以验证 Kind 字段的正确性

## Capabilities

### New Capabilities
- `symbol-kind-mapping`: 定义 Go 类型系统与 SCIP 符号类型之间的映射规则

### Modified Capabilities
（无）

## Impact

**Affected Code**:
- `internal/document/document.go` - `SetNewSymbol` 和 `SetNewSymbolForPos` 方法
- `internal/visitors/visitor_file.go` - `ToScipDocument` 方法中的局部符号创建
- `internal/lookup/lookup.go` - `SetPkgName` 方法
- `internal/testdata/snapshots/` - 所有快照测试输出

**Dependencies**:
- 使用 `github.com/sourcegraph/scip/bindings/go/scip` 中的 `SymbolInformation_Kind` 枚举值

**Systems**:
- SCIP 索引输出格式兼容性提升
- Sourcegraph 代码导航和语义搜索功能增强
