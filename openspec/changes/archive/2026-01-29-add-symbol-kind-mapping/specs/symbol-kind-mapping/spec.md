## ADDED Requirements

### Requirement: 符号类型映射

系统必须为所有输出的 `SymbolInformation` 正确设置 `Kind` 字段，使符号类型能够被消费者（如 Sourcegraph）准确识别。

#### Scenario: 函数符号映射
- **WHEN** Go 代码中定义一个函数（非方法）
- **THEN** 符号的 `Kind` 字段必须设置为 `Function`
- **AND** 符号信息必须包含正确的类型签名

#### Scenario: 方法符号映射
- **WHEN** Go 代码中定义一个方法（带有接收者）
- **THEN** 符号的 `Kind` 字段必须设置为 `Method`
- **AND** 符号信息必须包含方法签名和接收者类型

#### Scenario: 常量符号映射
- **WHEN** Go 代码中定义一个常量（const）
- **THEN** 符号的 `Kind` 字段必须设置为 `Constant`
- **AND** 符号信息必须包含常量的类型和值

#### Scenario: 结构体符号映射
- **WHEN** Go 代码中定义一个结构体（struct）
- **THEN** 符号的 `Kind` 字段必须设置为 `Class`
- **AND** 符号信息必须包含结构体的字段定义

#### Scenario: 接口符号映射
- **WHEN** Go 代码中定义一个接口（interface）
- **THEN** 符号的 `Kind` 字段必须设置为 `Interface`
- **AND** 符号信息必须包含接口的方法列表

#### Scenario: 结构体字段符号映射
- **WHEN** Go 代码中定义一个结构体字段
- **THEN** 符号的 `Kind` 字段必须设置为 `Field`
- **AND** 符号信息必须包含字段的类型

#### Scenario: 变量符号映射
- **WHEN** Go 代码中定义一个变量（var）
- **THEN** 符号的 `Kind` 字段必须设置为 `Variable`
- **AND** 符号信息必须包含变量的类型

#### Scenario: 包名符号映射
- **WHEN** Go 文件声明一个包名（package foo）
- **THEN** 符号的 `Kind` 字段必须设置为 `Namespace`
- **AND** 符号信息必须包含包路径

#### Scenario: 类型定义符号映射
- **WHEN** Go 代码中定义一个类型别名或其他类型（非 struct/interface）
- **THEN** 符号的 `Kind` 字段必须设置为 `Type`
- **AND** 符号信息必须包含类型的描述

#### Scenario: 无法识别的符号处理
- **WHEN** Go 代码中包含一个无法归类到上述类型的符号
- **THEN** 符号的 `Kind` 字段必须设置为 `UnspecifiedKind`
- **AND** 系统必须继续处理，不应 panic 或崩溃

#### Scenario: 局部符号类型映射
- **WHEN** 函数或方法内部定义局部变量
- **THEN** 符号的 `Kind` 字段必须根据其 Go 类型正确设置（如 `Variable`）
- **AND** 符号必须被正确添加到文档的 `Symbols` 列表中
