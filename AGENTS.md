# scip-go 开发指南

本项目是用于 Go 语言的 SCIP (Source Code Intelligence Protocol) 索引器。本指南旨在为 AI 代理和人类开发者提供统一的开发规范和工程实践。

## 1. 核心构建与测试命令

### 1.1 构建与安装
- **构建二进制**: `go build ./cmd/scip-go`
- **安装到本地**: `go install ./cmd/scip-go`
- **版本检查**: `scip-go --version` (使用 Go 1.24.3)
- **编译清理**: `go clean`

### 1.2 测试机制
项目深度依赖快照测试 (Snapshot Testing) 来验证索引结果。
- **运行所有测试**: `go test ./...`
- **运行特定包测试**: `go test ./internal/index`
- **运行特定快照用例**: `go test ./internal/index -filter <test_case_name>`
- **更新快照**: `go test ./internal/index -update-snapshots`
  - *注：当修改了 AST 遍历逻辑或符号提取算法时，必须更新并提交快照差异。*

### 1.3 调试与辅助工具
- **本地索引测试**: 在目标 Go 项目根目录运行 `scip-go`，生成 `index.scip` 文件。
- **检查缺失文件**: 使用 `scip-go --list-missing` 查看哪些文件未被正确索引。
- **查看 SCIP 索引**: 使用 `scip inspect` (需安装 SCIP CLI) 查看生成的索引内容。

## 2. 代码风格规范

### 2.1 导入管理 (Import Grouping)
严格遵循三段式分组，组间以空行分隔：
1. 标准库 (Standard Library)
2. 第三方库 (如 `github.com/sourcegraph/scip/...`)
3. 内部包 (以 `github.com/sourcegraph/scip-go/internal/...` 开头)

```go
import (
    "fmt"
    "go/ast"

    "github.com/charmbracelet/log"
    "github.com/sourcegraph/scip/bindings/go/scip"

    "github.com/sourcegraph/scip-go/internal/config"
    "github.com/sourcegraph/scip-go/internal/loader"
)
```

### 2.2 命名约定
- **包名**: 必须为小写单单词（如 `index`, `loader`, `visitors`），严禁使用下划线或混合大小写。
- **接口**: 遵循 Go 惯例，通常以 `-er` 结尾（如 `Writer`），或根据功能命名。
- **测试文件**: 推荐使用 `package_test` (黑盒测试) 以确保只测试公开 API。
- **私有成员**: 内部包中不打算导出的变量、函数和结构体应以小写字母开头。

### 2.3 错误处理
- **显式返回**: 库代码中必须返回 `error`，严禁使用 `panic`（极端不可恢复情况除外）。
- **错误包装**: 使用 `fmt.Errorf` 配合 `%v` 提供上下文信息。
  - *示例*: `return fmt.Errorf("failed to load package: %v", err)`
- **空检查**: 对 `loader.LoadPackages` 等返回的指针或 map 必须进行 nil 检查。
- **自定义错误**: 频繁出现的错误类型应在相关包中定义为导出的变量。

### 2.4 日志规范
- **结构化日志**: 统一使用 `github.com/charmbracelet/log`。
- **级别**: 关键路径使用 `Debug` 或 `Info`，可恢复问题使用 `Warn`，致命错误使用 `Error`。

## 3. 架构设计与约束

### 3.1 核心包职责
- `internal/index`: 协调整个索引过程。
- `internal/loader`: 封装 `go/packages`，负责加载类型信息和 AST。
- `internal/visitors`: 实现 AST 访问者，负责具体的符号和位置提取。
- `internal/document`: 管理 SCIP 文档的构建和状态。
- `internal/symbols`: 处理符号标识符的生成。

### 3.2 访问者模式 (Visitor Pattern)
核心逻辑位于 `internal/visitors`。索引过程分为两个主要阶段：
1. **全局扫描**: 遍历所有包的语法树，收集全局符号定义（`VisitPackageSyntax`）。
2. **详细索引**: 深度遍历每个文件，生成 SCIP 的 Occurrences 和符号详细信息。

### 3.3 并发与性能
- **并行处理**: 处理多个包时使用 `sync.WaitGroup` 和原子操作。
- **内存管理**: 避免在访问者状态中持有大量不必要的 AST 节点引用，以降低长周期内存压力。

### 3.4 数据流
- 输入: `go/packages` 加载的类型信息和 AST。
- 处理: 映射 Go 符号到 SCIP 符号标识符。
- 输出: 通过 `writer` 函数（通常是 Protobuf 序列化）发射 SCIP 消息。

## 4. 测试与验证

### 4.1 快照测试详解
- **输入路径**: `internal/testdata/snapshots/input/` 下的子目录。
- **预期输出**: `internal/testdata/snapshots/output/` 下对应的 `.scip` 格式化文本。
- **验证流程**: 修改逻辑 -> 运行测试 -> 检查差异 -> 确认无误后提交快照更新。

### 4.2 单元测试
- 对于独立的逻辑组件（如符号格式化），应编写常规的单元测试。
- 测试覆盖率应重点关注符号提取和位置计算。

## 5. 外部规则集成

### 5.1 Cursor/Copilot 规则
如果项目根目录存在 `.cursorrules` 或 `.cursor/rules/` 下的规则文件，请务必将其视为最高优先级的编码指令。目前的分析未发现这些文件，但未来添加时应严格遵守。

## 6. 开发建议与提交规范
1. **优先遵循现有模式**: 在修改代码前，参考同包内的其他文件。
2. **原子化提交**: 每次提交应专注于一个功能或修复，并伴随相应的快照更新。
3. **性能验证**: 对于核心遍历逻辑的修改，建议在大型项目（如 Kubernetes 或 Go 源码）上测试性能损耗。
4. **注释要求**: 对非显而易见的逻辑，特别是复杂的 AST 匹配逻辑，必须添加详细注释。
5. **Lint 检查**: 提交前确保通过 `golangci-lint run` (如果配置了)。
