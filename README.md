# GoToDo - 现代化待办清单应用

![GoToDo Logo](https://picsum.photos/400/200)

GoToDo是一个使用Golang开发的待办清单应用，帮助用户高效管理个人任务和日程安排。本项目采用MVC架构，结合Gin框架和GORM，实现了用户认证、清单管理和任务跟踪等核心功能。

## 🌟 功能亮点
- **用户管理**：支持注册、登录、密码加密存储
- **清单管理**：创建、查看、编辑和删除任务清单
- **任务操作**：添加、更新、标记完成状态和优先级管理
- **数据持久化**：使用MySQL存储用户数据，保证数据安全
- **可扩展性**：模块化设计，便于功能扩展和维护

## 🚀 快速开始

### 环境准备
- Go 1.20+
- MySQL 8.0+
- Postman (用于API测试)

### 安装步骤
1. **克隆项目**
   ```bash
   git clone https://github.com/SBHHX/go-to-list.git
   cd go-to-list
   ```

2. **安装依赖**
   ```bash
   go mod tidy
   ```

3. **配置数据库**
   - 创建MySQL数据库：`todo_list_db`
   - 修改`db/db.go`中的数据库连接信息：
     ```go
     dsn := "user:password@tcp(127.0.0.1:3306)/todo_list_db?charset=utf8mb4&parseTime=True&loc=Local"
     ```

4. **运行项目**
   ```bash
   go run main.go
   ```

5. **测试API**
   使用Postman导入以下API集合进行测试：
   ```
   POST /api/v1/user/register - 用户注册
   POST /api/v1/user/login - 用户登录
   POST /api/v1/list/create - 创建清单
   POST /api/v1/task/add - 添加任务
   ```

## 📦 项目结构project/
├── db/                 - 数据库连接和迁移
├── models/             - 数据模型
├── controllers/        - 控制器层
├── routers/            - 路由配置
├── services/           - 业务逻辑
└── main.go             - 入口文件
## 🛠️ 技术栈
- **后端框架**：Gin
- **ORM**：GORM
- **数据库**：MySQL
- **密码加密**：bcrypt
- **API测试**：Postman

## 🤝 贡献指南
1. Fork本仓库
2. 创建你的特性分支 (`git checkout -b feature/fooBar`)
3. 提交你的更改 (`git commit -am 'Add some fooBar'`)
4. 将你的更改推送到分支 (`git push origin feature/fooBar`)
5. 创建新的Pull Request

## 📄 许可证
木有许可证

## 📧 联系我们
- 项目负责人：SBHHX
- 邮箱：1573057383@qq.com
- GitHub：https://github.com/SBHHX/go-to-list
    
