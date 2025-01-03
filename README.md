
# EasyGo
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

English | [中文](./README_CN.md)

A CLI tool for scaffolding Go web projects  
    
## Features

- **Web Framework Support**: Choose from popular web frameworks such as Gin, Echo, Fiber, and more to suit your project requirements.
  
- **ORM Tool Support**: Easily integrate with different ORM libraries, including GORM, Ent, XORM, allowing you to work with your database in an object-oriented manner.

- **Database Configuration**: Configure and connect to various relational databases, including MySQL, PostgreSQL, and more, with minimal setup.

- **Common Functional Modules**: Automatically generate essential modules for your application, such as:
  - **User Authentication**: Set up user login, registration, and session management with ease.
  - **Permission Control**: Implement role-based or fine-grained access control to secure your application.
  - **Logging**: Automatically generate logging functionality to capture application activity and debug information.
  - **Error Handling**: Built-in error handling module for consistent and user-friendly error reporting across your application.

## Roadmap

- [x] Web framework support: Echo, Gin
- [ ] ORM framework support: Xorm, Ent
- [ ] i18n internationalization
- [ ] API testing
- [ ] Containerization support: Generate Dockerfile and Kubernetes YAML files
- [ ] CI/CD support: Generate CI/CD configuration files (GitHub Actions, GitLab CI, Jenkins)
- [ ] API documentation: Swagger, Markdown, HTML support
- [ ] Code quality checks

**🤔However, I currently plan to give priority to reconstructing the static generation part of the template code.**

## Getting Started

EasyGo is a command-line tool for quickly scaffolding Go web applications with support for multiple frameworks, ORM tools, and databases.

### Prerequisites

The things you need before installing the software.

- Go (v1.18+): Install Go from Go Downloads.
- Git: Install Git from Git Downloads.
- Access to GitHub (Optional, for installing via go install): Ensure your network allows access to GitHub.
- A GitHub Account (Optional, for contributing or creating issues): Needed for reporting issues or contributing.

### Installation

To get started, you can install EasyGo via go install:

```bash
# Go 1.18 and later version
GOPROXY=https://goproxy.cn/,direct go install github.com/Whitea029/easygo@latest
```
Alternatively, clone the repository and build it locally:
```bash
git clone https://github.com/Whitea029/easygo.git
cd easygo
go build -o easygo
```
### Usage/Examples

#### Initialize a New Project

Run the init command to create a new project:

```bash
easygo init --name easygo_start --module example.com/yourname/easygo_start
```
Start the Go project using the go run command:
```bash
go run ./cmd/main.go
```

## Contributing

Contributions are welcome! Here's how you can help:

1. **Report Issues**: Found a bug? Open an issue with details and steps to reproduce.
2. **Suggest Features**: Have an idea? Create an issue labeled `enhancement`.
3. **Submit Changes**:
   - Fork the repo and create a feature branch: `git checkout -b feature/your-feature-name`
   - Commit your changes and open a pull request.

Thank you for your support!
