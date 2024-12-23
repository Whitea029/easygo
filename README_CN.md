# EasyGo
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
中文 | [English](./README.md)

一个用于快速搭建 Go Web 项目的命令行工具  

## 功能特性

- **Web 框架支持**：支持选择常见的 Web 框架，如 Gin、Echo、Fiber 等，满足不同项目需求。
  
- **ORM 工具支持**：集成多种 ORM 库，如 GORM、Ent、XORM，便于以面向对象的方式操作数据库。

- **数据库配置**：轻松连接多种关系型数据库，包括 MySQL、PostgreSQL 等，配置简单高效。

- **常用功能模块**：自动生成应用程序的核心模块：
  - **用户认证**：快速搭建用户登录、注册与会话管理功能。
  - **权限控制**：实现基于角色或细粒度的权限管理，保护应用安全。
  - **日志记录**：自动生成日志功能，便于捕获应用活动和调试信息。
  - **错误处理**：内置错误处理模块，提供一致且友好的错误信息反馈。

## 路线图

- [x] Web 框架支持：Echo、Hertz
- [ ] ORM 框架支持：Xorm、Ent
- [ ] i18n 国际化支持
- [ ] API 测试功能
- [ ] 容器化支持：生成 Dockerfile 和 Kubernetes YAML 文件
- [ ] CI/CD 支持：生成 CI/CD 配置文件（GitHub Actions、GitLab CI、Jenkins 等）
- [ ] API 文档支持：Swagger、Markdown、HTML 格式
- [ ] 代码质量检查

**🤔目前优先计划重构模板代码的静态生成部分。**

## 快速开始

EasyGo 是一个命令行工具，支持多框架、ORM 工具和数据库，帮助快速搭建 Go Web 应用程序。

### 前置要求

使用前需要确保安装以下软件：

- Go (v1.17+): 从 [Go Downloads](https://golang.org/dl/) 安装。
- Git: 从 [Git Downloads](https://git-scm.com/) 安装。
- Make (可选，用于从源码构建): 根据系统需求安装 Make（如 macOS 使用 `brew install make`，Linux 使用 `sudo apt-get install make`）。
- GitHub 访问权限 (可选，用于通过 `go install` 安装): 确保网络允许访问 GitHub。
- GitHub 账号 (可选，用于贡献或创建问题): 便于报告问题或贡献代码。

### 安装

通过 `go install` 安装 EasyGo：

```bash
# Go 1.15 及更早版本
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get github.com/Whitea029/easygo@latest

# Go 1.16 及更高版本
GOPROXY=https://goproxy.cn/,direct go install github.com/Whitea029/easygo@latest
```
或者，克隆仓库并本地构建：
```bash
git clone https://github.com/Whitea029/easygo.git
cd easygo
go build -o easygo
```
### 使用示例
#### 初始化新项目
使用`init`命令创建新项目：
```bash
easygo init --name easygo_start --module example.com/yourname/easygo_start
```
通过`go run`命令启动项目：
```bash
go run ./cmd/main.go
```
## 贡献指南
欢迎贡献代码！以下是贡献的简单步骤：

1. **报告问题**：发现 Bug？请在 Issues 页面描述问题和复现步骤。
2. **建议新功能**：有好的想法？请创建一个带有 `enhancement` 标签的 Issue。
3. **提交代码**：
4. **`Fork` 仓库并创建分支**：`git checkout -b feature/your-feature-name`
提交更改并创建 `Pull Request`。

感谢你的支持
