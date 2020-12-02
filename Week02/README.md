
# 学习笔记

## 知识点
### 1. error.New() 返回值为什么是一个指针？

### 2. 传统error的使用几种方式？
1. Sentinel Error
2. 自定义 Error types + 断言
3. Opaque errors (不透明处理)

### 3. 更优雅的error处理方式 - wrap errors
1. 错误要被日志记录
2. 应用程序处理错误，保证100%完整性
3. 之后不要再报当前错误。`you should only handle errors once. Handling an error means inspecting the error value, and making a single decision.`

## 编码规范
1. `Indented flow is for errors` 无错误的正常流程代码，将成为一条直线，而不是缩进的代码
2. `Eliminate error handling by eliminating errors`
3. `you should only handle errors once. Handling an error means inspecting the error value, and making a single decision.`

## 作业
### 作业要求
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

### 囧囧
go还没学习相关web框架，也不知道这次写的程序符不符合作业规范。