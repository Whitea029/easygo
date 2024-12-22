# easygo 需求文档
## 需求
### easygo
#### 需求分析
**概括：**
实现一键生成go的web项目的脚手架的cli工具，web框架可以用户选择Gin，echo等等，orm可以选择gorm，xorm，ent等等，cli工具打算使用cobra提高用户体验。代码的生成打算使用go的text/template进行代码生成。目前的开发流程计划为，先本地写好一个脚手架，确保项目是可以运行的，然后将各个模块用template进行修改，然后写代码生成的部分，最后写cli工具。

**基本功能：**
- Web 框架选择：支持选择不同的 Web 框架，比如 Gin、Echo、Fiber 等。
- ORM 选择：支持选择不同的 ORM 工具，比如 GORM、Ent、XORM 等。
- 数据库配置：支持 MySQL、PostgreSQL 等数据库配置。
- 常见功能模块生成：比如用户认证、权限控制、日志、错误处理等常见功能模块的自动生成。

**可拓展性：**
- 模板系统：支持根据用户选择的框架和配置生成自定义的代码模板。用户可以添加新的模板或修改现有模板。
- 灵活配置：允许用户通过配置文件或命令行参数来定制项目结构和模块，比如数据库连接、日志配置等。 

**生成项目目录结构：**
> /project-name
  ├── /cmd              # 主程序入口
  ├── /config           # 配置文件
  ├── /internal         # 内部模块（用于实现业务逻辑）
  │   ├── /controller   # 控制器
  │   ├── /model        # 数据模型
  │   ├── /service      # 服务层
  │   ├── /repository   # 数据访问层（与 ORM 相关）
  ├── /pkg              # 公共库，其他项目可引用
  ├── /scripts          # 脚本文件
  ├── /web               # 前端页面相关（如果是后端生成前端项目）
  ├── go.mod
  └── README.md

#### 技术实现
- 模板引擎选择`text/template`
### easygo-cli
#### 第三方库
cobra
#### 命令设计
- `init`：初始化一个新的项目。
  - `-n (--name)` 项目名称
  - `-m (--module)` 项目的Go mudule
  - `-o (--orm)` orm框架 gorm ent xorm等
  - `-w (--web)` web框架 Gin Echo等
  - `-db (--database)` 数据库 MySql PosegreSql等
- `generate`：根据用户的配置生成代码模块。
- `add`：向现有项目添加新的功能模块（如添加身份验证模块等）。
- `list`：列出当前支持的模块和框架。
- `version`：显示脚手架工具的版本。
#### 使用示例
`easygo init --module` 


~~日志的完善~~ 
~~conf和logger的使用~~
~~异常和错误处理~~
~~api层完善~~
~~model的构建~~
~~整个项目结构完善~~

conf层代码生成
pkg层代码生成
api层代码生成
dal层代码生成
cmd层代码生成
Dockerfile
docker-compose
项目结构补全
go mod tidy

cobra cli引导
大量条件判断完善
配置文件读取

idl的读取

测试